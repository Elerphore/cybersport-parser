package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Insert(fn func(*sql.DB)) {
	db, sqliteconnectionErr := sql.Open("sqlite3", "news.db")

	if sqliteconnectionErr != nil {
		log.Fatal(sqliteconnectionErr)
	}

	defer db.Close()

	fn(db)
}

func Select(fn func(*sql.DB)) {
	db, sqliteconnectionErr := sql.Open("sqlite3", "news.db")

	if sqliteconnectionErr != nil {
		log.Fatal(sqliteconnectionErr)
	}

	defer db.Close()

	fn(db)
}

func InsertManyNews(newsList []News) {
	Insert(func(db *sql.DB) {
		for _, news := range newsList {
			result, insertErr := db.Exec("insert into news (link, news_source_id) values ($1, $2)", news.Link, news.NewsSourceId)

			if insertErr != nil {
				log.Fatal(insertErr)
			}

			fmt.Println(result.RowsAffected())
		}
	})
}

func InsertNews(news News) {
	Insert(func(db *sql.DB) {
		result, insertErr := db.Exec("insert into news (link, news_source_id) values ($1, $2)", news.Link, news.NewsSourceId)

		if insertErr != nil {
			log.Fatal(insertErr)
		}

		fmt.Println(result.RowsAffected())
	})
}

func GetNews() (newsList []News) {
	Select(func(db *sql.DB) {
		rows, rowsError := db.Query("SELECT * FROM news ORDER BY id DESC LIMIT 1")

		if rowsError != nil {
			log.Fatalln(rowsError)
		}

		defer rows.Close()

		for rows.Next() {
			n := News{}
			err := rows.Scan(&n.Id, &n.Link, &n.NewsSourceId)

			if err != nil {
				fmt.Println(err)
				continue
			}

			newsList = append(newsList, n)
		}

	})

	return newsList
}
