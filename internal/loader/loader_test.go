package loader

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wouove/learning-words/internal/model"
)

func TestCSVLoaderAdapter_Load(t *testing.T) {
	path := "/Users/wouteroverbeek/personal_repositories/learning-words/output/"
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
	path := "/Users/wouteroverbeek/personal_repositories/learning-words/output/anki-format"
	//fileName := "test.apkg"
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
