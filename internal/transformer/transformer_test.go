package transformer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wouove/learning-words/internal/model"
)

type MockExtractor struct{}

func (e MockExtractor) Extract() ([]model.WordPair, error) {
	return []model.WordPair{
		{
			Word:                "house",
			Translation:         "",
			WordLanguage:        "English",
			TranslationLanguage: "Dutch",
		},
	}, nil
}

type MockClient struct{}

func (c MockClient) Translate(word model.WordPair) (model.WordPair, error) {
	return model.WordPair{
		Word:                word.Word,
		Translation:         "huis",
		WordLanguage:        "English",
		TranslationLanguage: "Dutch",
	}, nil
}

type MockLoader struct{}

func (l MockLoader) Load(words []model.WordPair, skipHeader bool, multipleChoice bool) error {
	fmt.Println(words)
	return nil
}

func (l MockLoader) Open() ([]model.WordPair, error) {
	return []model.WordPair{}, nil
}

func TestTransformer_TranslateWords(t *testing.T) {
	extractor := MockExtractor{}
	client := MockClient{}
	testLoader := MockLoader{}

	transformer := NewTransformer(extractor, testLoader, client)
	transformer.Transform()

}

func TestTransformer_FilterTranslatedWords(t *testing.T) {
	test := []model.WordPair{
		{
			Word:                "house",
			Translation:         "huis",
			WordLanguage:        "English",
			TranslationLanguage: "Dutch",
		},
		{
			Word:                "cat",
			Translation:         "cat",
			WordLanguage:        "English",
			TranslationLanguage: "Dutch",
		},
	}
	extractor := MockExtractor{}
	client := MockClient{}
	testLoader := MockLoader{}

	transformer := NewTransformer(extractor, testLoader, client)
	result := transformer.FilterTranslatedWords(test)
	require.Equal(t, []model.WordPair{test[0]}, result)
}

func TestTransformer_FilterAlreadyTranslatedWords(t *testing.T) {
	testInput := []model.WordPair{
		{
			Word:                "house",
			Translation:         "huis",
			WordLanguage:        "English",
			TranslationLanguage: "Dutch",
		},
		{
			Word:                "cat",
			Translation:         "kat",
			WordLanguage:        "English",
			TranslationLanguage: "Dutch",
		},
	}
	testProcessedWords := []model.WordPair{
		{
			Word:                "house",
			Translation:         "huis",
			WordLanguage:        "English",
			TranslationLanguage: "Dutch",
		},
	}
	extractor := MockExtractor{}
	client := MockClient{}
	testLoader := MockLoader{}

	transformer := NewTransformer(extractor, testLoader, client)
	result := transformer.FilterAlreadyTranslatedWords(testInput, testProcessedWords)
	require.Equal(t, 1, len(result))
	require.Equal(t, []model.WordPair{testInput[1]}, result)
}
