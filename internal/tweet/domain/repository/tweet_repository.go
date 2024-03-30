package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/madruga665/twitter-api/internal/tweet/domain/entity"
)

type tweetRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *tweetRepository {
	return &tweetRepository{
		db: db,
	}
}

func (repository *tweetRepository) GetAll() (*sql.Rows, error) {
	query := "SELECT id, description FROM tweets"
	result, err := repository.db.Query(query)

	return result, err
}

func (repository *tweetRepository) GetById(tweetID string) (*sql.Rows, error) {
	query := "SELECT id, description FROM tweets WHERE id = $1"
	result, error := repository.db.Query(query, tweetID)

	return result, error
}

func Create(db *sql.DB, tweet entity.Tweet) error {
	query := `INSERT INTO tweets (id, description) VALUES ($1, $2);`
	_, err := db.Exec(query, tweet.ID, tweet.Description)

	return err
}

func Delete(db *sql.DB, tweetID string) error {
	query := "DELETE FROM tweets WHERE id = $1"
	_, err := db.Exec(query, tweetID)

	return err
}

func Update(db *sql.DB, tweetID string, descriptionBody string) error {
	query := "UPDATE tweets SET description = $1 WHERE id = $2"
	_, error := db.Exec(query, descriptionBody, tweetID)

	return error
}
