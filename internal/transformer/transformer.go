package transformer

import (
	"fmt"

	"github.com/wouove/learning-words/internal/client"
	"github.com/wouove/learning-words/internal/extractor"
	"github.com/wouove/learning-words/internal/loader"
	"github.com/wouove/learning-words/internal/model"
)

type Transformer struct {
	ExtractorPort extractor.WordsExtractorPort
	LoaderPort    loader.WordsLoaderPort
	client        client.Client
}

func NewTransformer(extractor extractor.WordsExtractorPort, loader loader.WordsLoaderPort, client client.Client) Transformer {
	return Transformer{
		ExtractorPort: extractor,
		LoaderPort:    loader,
		client:        client,
	}
}

func (t Transformer) Transform() error {
	// read input file
	wordsInput, err := t.ExtractorPort.Extract()
	if err != nil {
		return fmt.Errorf("getting input from file: %w", err)
	}
	fmt.Printf("Found %d words in the input file\n", len(wordsInput))

	// read already translated words
	processesWords, err := t.LoaderPort.Open()
	if err != nil {
		return fmt.Errorf("reading output file: %w\n", err)
	}
	fmt.Printf("There are %d words already translated in an earlier run\n", len(processesWords))
	// filter words from input that were already translated
	filteredWordsInput := t.FilterAlreadyTranslatedWords(wordsInput, processesWords)
	fmt.Printf("After filtering out already translated words, %d words remain that will be translated\n", len(filteredWordsInput))

	// translate words
	translatedWords, err := t.TranslateWords(filteredWordsInput)
	if err != nil {
		return fmt.Errorf("translating words: %w", err)
	}

	// filter words
	filteredWords := t.FilterTranslatedWords(translatedWords)
	fmt.Printf("The translation client returned %d words of which %d were actually translated\n", len(translatedWords), len(filteredWords))

	// store
	wordsToLoad := append(processesWords, filteredWords...)
	err = t.LoaderPort.Load(wordsToLoad, true, false)
	if err != nil {
		return fmt.Errorf("storing words: %w", err)
	}

	return nil
}

func (t Transformer) TranslateWords(words []model.WordPair) ([]model.WordPair, error) {
	var result []model.WordPair
	for _, word := range words {
		translatedWord, err := t.client.Translate(word)
		if err != nil {
			return []model.WordPair{}, fmt.Errorf("getting deepl translation for word %s, from language %s to language %s: %w", word.Word, word.WordLanguage, word.TranslationLanguage, err)
		}
		result = append(result, translatedWord)
	}
	return result, nil
}

func (t Transformer) FilterAlreadyTranslatedWords(input []model.WordPair, processedWords []model.WordPair) []model.WordPair {
	processedWordsMap := make(map[string]bool)
	for _, word := range processedWords {
		processedWordsMap[word.Word] = true
	}
	var result []model.WordPair
	for _, word := range input {
		if _, ok := processedWordsMap[word.Word]; ok {
			continue
		}
		result = append(result, word)
	}
	return result
}

func (t Transformer) FilterTranslatedWords(words []model.WordPair) []model.WordPair {
	// Filter words that were not translated
	var result []model.WordPair
	for _, word := range words {
		if word.Word == word.Translation {
			continue
		}
		result = append(result, word)
	}
	return result
}
