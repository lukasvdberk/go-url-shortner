package database

import (
	"database/sql"
	"fmt"
	"os"
)

import _ "github.com/go-sql-driver/mysql"

func getSqlConnection() *sql.DB {
	// Gets connection info from environment variables
	dbHost := "tcp(go-url-shortner-database:3306)"
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	databaseName := os.Getenv("MYSQL_DATABASE")

	connectionString := user + ":" + password + "@" + dbHost + "/" + databaseName

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err)
	}

	// To test whether the connection worked
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return db
}

func SaveUrl(shortUrl *ShortUrl) *ShortUrl {
	// @return returns new shorturl with id filled or nil if failed.
	db := getSqlConnection()
	stmtIns, err := db.Prepare("INSERT INTO ShortUrl (realUrl) VALUES(?)") // ? = placeholder
	if err != nil {
		return nil
	}

	res, err := stmtIns.Exec(shortUrl.RealUrl)
	if err != nil {
		return nil
	}
	shortUrl.Id, err = res.LastInsertId()

	if err != nil {
		return nil
	}

	return shortUrl
}

func GetShortUrlById(id int64) *ShortUrl {
	db := getSqlConnection()
	stmtOut, err := db.Prepare("SELECT realUrl FROM ShortUrl WHERE id = ?") // ? = placeholder
	if err != nil {
		return nil
	}

	var realUrl string
	err = stmtOut.QueryRow(id).Scan(&realUrl)

	if err != nil {
		return nil
	}

	shortUrl := new(ShortUrl)
	shortUrl.Id = id
	shortUrl.RealUrl = realUrl
	return shortUrl
}
