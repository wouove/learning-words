package loader

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/plandem/anki-apkg"
	"github.com/wouove/learning-words/internal/model"
)

type WordsLoaderPort interface {
	Open() ([]model.WordPair, error)
	Load([]model.WordPair, bool, bool) error
}

type CSVLoaderAdapter struct {
	path            string
	stringFormatter StringFormatter
}

func NewCSVLoaderAdapter(path string, stringFormatter StringFormatter) CSVLoaderAdapter {
	return CSVLoaderAdapter{path: path, stringFormatter: stringFormatter}
}

func (c CSVLoaderAdapter) Load(words []model.WordPair, skipHeader bool, multipleChoice bool) error {
	csvFile, err := os.Create(c.path)
	if err != nil {
		return fmt.Errorf("writing to csv, %w", err)
	}
	csvwriter := csv.NewWriter(csvFile)

	wordsString, err := c.stringFormatter.TransformToString(words, skipHeader)
	if err != nil {
		return fmt.Errorf("transforming to string %w", err)
	}
	for _, word := range wordsString {
		_ = csvwriter.Write(word)
	}
	csvwriter.Flush()
	csvFile.Close()

	return nil
}

func (c CSVLoaderAdapter) Open() ([]model.WordPair, error) {
	f, err := os.Open(c.path)
	if err != nil {
		return []model.WordPair{}, fmt.Errorf("reading output CSV: %w", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return []model.WordPair{}, fmt.Errorf("parsing output CSV: %w", err)
	}
	var result []model.WordPair
	for _, line := range records {
		result = append(result, model.WordPair{
			Word:                line[0],
			Translation:         line[1],
			WordLanguage:        "English",
			TranslationLanguage: "Dutch",
		})
	}

	return result, nil
}

type AnkiWordsLoaderAdapter struct {
	path string
}

func NewAnkiWordsLoaderAdapter(path string) AnkiWordsLoaderAdapter {
	return AnkiWordsLoaderAdapter{path: path}
}

func (l AnkiWordsLoaderAdapter) Load(words []model.WordPair, skipHeader bool) error {
	// Save file to .apkg or .colpkg
	// create file
	file := apkg.Create(l.path)
	// modify file to include the words
	if file == nil {
		return fmt.Errorf("creating ANKI .apkg file failed")
	}
	return nil
}
