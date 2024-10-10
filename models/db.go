// Package models models/db.go
package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB DB는 데이터베이스 연결을 나타냅니다.
var DB *sql.DB

// InitDB InitDB는 데이터베이스를 초기화하고 필요한 테이블을 생성합니다.
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// posts 테이블 생성
	createTableSQL := `CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL
    );`

	statement, err := DB.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = statement.Exec(); err != nil {
		log.Fatal(err)
	}
}
