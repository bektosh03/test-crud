package adapters

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bektosh03/test-crud/data-service/domain/post"
	"github.com/bektosh03/test-crud/common/errs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	postsTableName          = "posts"
	downloadStatusTableName = "download_status"
)

type postModel struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Title  string `db:"title"`
	Body   string `db:"body"`
}

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) CreatePostsAsync(ctx context.Context, postsBatch <-chan []post.Post) <-chan error {
	errChan := make(chan error)
	go func() {
		defer close(errChan)
		err := r.transact(func(tx *sqlx.Tx) error {
			for {
				select {
				case posts, open := <-postsBatch:
					if !open {
						return nil
					}

					if err := createPosts(ctx, tx, posts); err != nil {
						logrus.Errorf("error while creating posts: %v", err)
						return err
					}

				case <-ctx.Done():
					logrus.Errorf("canceled create posts, error: %v", ctx.Err())
					return ctx.Err()
				}
			}
		})

		errChan <- err
	}()

	return errChan
}

func (r *PostgresRepository) SetDownloadStatus(ctx context.Context, success bool, downloadErr error) error {
	return setDownloadStatus(ctx, r.db, success, downloadErr)
}

func (r *PostgresRepository) GetDownloadStatus(ctx context.Context) (success bool, errMsg string, err error) {
	return getDownloadStatus(ctx, r.db)
}

func (r *PostgresRepository) transact(txFunc func(*sqlx.Tx) error) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil, don't change it
		} else {
			err = tx.Commit() // err is nil, if Commit returns error, update err
		}
	}()

	return txFunc(tx)
}

func createPosts(ctx context.Context, tx *sqlx.Tx, posts []post.Post) error {
	for _, p := range posts {
		if err := createPost(ctx, tx, toPostModel(p)); err != nil {
			return err
		}
	}

	return nil
}

func createPost(ctx context.Context, tx *sqlx.Tx, pm postModel) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (id, user_id, title, body)
		VALUES ($1, $2, $3, $4)`, postsTableName,
	)
	_, err := tx.ExecContext(ctx, query, pm.ID, pm.UserID, pm.Title, pm.Body)

	return err
}

func setDownloadStatus(ctx context.Context, db *sqlx.DB, success bool, downloadErr error) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (success, err) VALUES ($1, $2)`, downloadStatusTableName,
	)
	var errWrapper sql.NullString
	if downloadErr != nil {
		errWrapper = sql.NullString{
			String: downloadErr.Error(),
			Valid:  true,
		}
	}

	_, err := db.ExecContext(ctx, query, success, errWrapper)

	return err
}

func getDownloadStatus(ctx context.Context, db *sqlx.DB) (success bool, errMsg string, err error) {
	query := fmt.Sprintf(`SELECT * FROM %s ORDER BY created_at ASC LIMIT 1`, downloadStatusTableName)
	var errWrapper sql.NullString

	err = db.QueryRowxContext(ctx, query).Scan(&success, &errWrapper)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, "", errs.ErrNotFound
		}

		return false, "", err
	}

	return success, errWrapper.String, nil
}

func toPostModel(p post.Post) postModel {
	return postModel{
		ID:     p.ID,
		UserID: p.UserID,
		Title:  p.Title,
		Body:   p.Body,
	}
}
