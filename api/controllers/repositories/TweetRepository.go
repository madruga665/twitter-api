package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/madruga665/twitter-api/api/entities"
)

func Create(db *sql.DB, tweet entities.Tweet) error {
	query := `INSERT INTO tweets (id, description) VALUES ($1, $2);`
	_, err := db.Exec(query, tweet.ID, tweet.Description)

	return err
}

func GetAll(db *sql.DB) ([]entities.Tweet, error) {
	query := "SELECT id, description FROM tweets"
	result, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var tweets []entities.Tweet
	for result.Next() {
		var tweet entities.Tweet
		err := result.Scan(&tweet.ID, &tweet.Description)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}

func GetById(db *sql.DB, tweetID string) (entities.Tweet, error) {
	query := "SELECT id, description FROM tweets WHERE id = $1"
	result, err := db.Query(query, tweetID)
	if err != nil {
		return entities.Tweet{}, err
	}
	defer result.Close()

	var tweet entities.Tweet
	for result.Next() {
		err := result.Scan(&tweet.ID, &tweet.Description)
		if err != nil {
			return entities.Tweet{}, err
		}
	}

	return tweet, nil
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
