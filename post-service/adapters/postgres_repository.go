package adapters

import (
	"context"
	"fmt"

	"github.com/bektosh03/test-crud/post-service/domain/post"
	"github.com/jmoiron/sqlx"
)

const (
	postsTableName = "posts"
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

func (r *PostgresRepository) GetPost(ctx context.Context, postID int) (post.Post, error) {
	
}

func getPost(ctx context.Context, db *sqlx.DB, postID int) (postModel, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, postsTableName)
	var pm postModel

	err := db.GetContext(ctx, &pm, query, postID)

	return pm, err
}

func listPosts(ctx context.Context, db *sqlx.DB, page, limit int) ([]postModel, error) {
	query := fmt.Sprintf(
		`SELECT * FROM %s OFFSET $1 LIMIT $2`, postsTableName,
	)
	offset := (page - 1) * limit
	posts := make([]postModel, 0, limit)

	err := db.SelectContext(ctx, &posts, query, offset, limit)

	return posts, err
}

func updatePost(ctx context.Context, db *sqlx.DB, pm postModel) error {
	query := fmt.Sprintf(
		`UPDATE %s SET title = $1, body = $2 WHERE id = $3`, postsTableName,
	)

	_, err := db.ExecContext(ctx, query, pm.Title, pm.Body, pm.ID)

	return err
}

func deletePost(ctx context.Context, db *sqlx.DB, postID int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, postsTableName)

	_, err := db.ExecContext(ctx, query, postID)

	return err
}
