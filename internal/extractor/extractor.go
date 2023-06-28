package extractor

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wouove/learning-words/internal/model"
)

type WordsExtractorPort interface {
	Extract() ([]model.WordPair, error)
}

type SqlLiteExtractorAdapter struct {
	DataBase *sql.DB
}

const wordLanguage = "English"
const translationLanguage = "Dutch"

func NewSqlLiteExtractorAdapter(DbFilepath string) (SqlLiteExtractorAdapter, error) {
	sqliteDatabase, err := sql.Open("sqlite3", DbFilepath)
	if err != nil {
		return SqlLiteExtractorAdapter{}, fmt.Errorf("setting up SqlLiteExtractorAdapter: %w", err)
	}
	//defer sqliteDatabase.Close()
	return SqlLiteExtractorAdapter{
		DataBase: sqliteDatabase,
	}, nil
}

func (a SqlLiteExtractorAdapter) Extract() ([]model.WordPair, error) {
	var result []model.WordPair
	query := "SELECT Text FROM WordList"
	row, err := a.DataBase.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	defer a.DataBase.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var word string
		row.Scan(&word)
		//log.Println("Word: ", word)
		result = append(result, model.WordPair{
			Word:                strings.ToLower(word),
			Translation:         "",
			WordLanguage:        wordLanguage,
			TranslationLanguage: translationLanguage,
		})
	}
	return result, nil
}
