package adapters

import (
	"context"
	"fmt"

	"github.com/bektosh03/test-crud/data-service/domain/post"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const postsTableName = "posts"

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
				case posts, ok := <-postsBatch:
					if !ok {
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

func toPostModel(p post.Post) postModel {
	return postModel{
		ID:     p.ID,
		UserID: p.UserID,
		Title:  p.Title,
		Body:   p.Body,
	}
}
