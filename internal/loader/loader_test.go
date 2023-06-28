package loader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wouove/learning-words/internal/model"
)

func TestCSVLoaderAdapter_Load(t *testing.T) {
	path := os.Getenv("OUTPUT_PATH")
	fileName := "test.csv"
	loader := NewCSVLoaderAdapter(path + fileName)

	input := []model.WordPair{
		{
			Word:                "House",
			Translation:         "Huis",
			WordLanguage:        "English",
			TranslationLanguage: "Dutch",
		},
	}
	err := loader.Load(input, false, false)
	require.NoError(t, err)
}

func TestAnkiWordsLoaderAdapter_Load(t *testing.T) {
	path := os.Getenv("OUTPUT_PATH")
	loader := NewAnkiWordsLoaderAdapter(path)

	input := []model.WordPair{
		{
			Word:                "House",
			Translation:         "Huis",
			WordLanguage:        "English",
			TranslationLanguage: "Dutch",
		},
	}
	err := loader.Load(input, true)
	require.NoError(t, err)
}
