package loader

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/wouove/learning-words/internal/model"
)

type StringFormatter interface {
	TransformToString(words []model.WordPair, skipHeader bool) ([][]string, error)
}

type CardStringFormatter struct {
	header []string
}

func NewCardStringFormatter() CardStringFormatter {
	return CardStringFormatter{
		header: []string{"Question", "Answer"},
	}
}

func (s CardStringFormatter) TransformToString(words []model.WordPair, skipHeader bool) ([][]string, error) {
	var result [][]string
	if !skipHeader {
		result = append(result, s.header)
	}
	for _, word := range words {
		result = append(result, []string{word.Word, word.Translation})
	}
	return result, nil
}

type MultipleChoiceStringFormatter struct {
	header          []string
	numberOfAnswers int
}

func NewMultipleChoiceStringFormatter(numberOfAnswers int) MultipleChoiceStringFormatter {
	header := []string{"Question", "Question code MC", "Answer pattern"}
	for i := 0; i < numberOfAnswers; i++ {
		header = append(header, fmt.Sprintf("Answer %d", i))
	}
	return MultipleChoiceStringFormatter{
		header:          header,
		numberOfAnswers: numberOfAnswers,
	}
}

func (s MultipleChoiceStringFormatter) TransformToString(words []model.WordPair, skipHeader bool) ([][]string, error) {
	var result [][]string
	if !skipHeader {
		result = append(result, s.header)
	}
	//numberOfWords := len(words)
	for i, word := range words {
		resultElement := make([]string, 3+s.numberOfAnswers)
		answerPattern := make([]int, s.numberOfAnswers)
		// get wrong translations
		wrongTranslations, err := s.selectRandomWrongTranslations(words, i)
		if err != nil {
			return [][]string{}, fmt.Errorf("transforming to string: %w", err)
		}
		// put the correct answer at a random spot in the list of answers, but remember spot
		indexCorrectTranslation := rand.Intn(s.numberOfAnswers)
		// set the answer pattern
		answerPattern[indexCorrectTranslation] = 1
		// assemble result element and append
		resultElement[0] = word.Word
		resultElement[1] = "1" // ANKI code for multiple choice
		resultElement[2] = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(answerPattern)), " "), "[]")
		wrongTranslationsToInsert := wrongTranslations
		for j := 0; j < s.numberOfAnswers; j++ {
			if j == indexCorrectTranslation {
				resultElement[j+3] = word.Translation
			} else {
				wrongTranslationToInsert, wrongTranslationsNotToInsert := wrongTranslationsToInsert[0], wrongTranslationsToInsert[1:]
				resultElement[j+3] = wrongTranslationToInsert
				wrongTranslationsToInsert = wrongTranslationsNotToInsert
			}
		}
		result = append(result, resultElement)
	}
	return result, nil
}

func (s MultipleChoiceStringFormatter) selectRandomWrongTranslations(words []model.WordPair, indexCorrectWord int) ([]string, error) {
	numberOfWords := len(words)
	var wrongTranslations []string
	numberOfWrongTranslations := s.numberOfAnswers - 1
	for j := 1; j <= numberOfWrongTranslations; j++ {
		index, err := s.getRandomIndex(numberOfWords, indexCorrectWord)
		if err != nil {
			return []string{}, fmt.Errorf("selecting random wrong translations: %w", err)
		}
		wrongTranslations = append(wrongTranslations, words[index].Translation)
	}
	return wrongTranslations, nil
}

func (s MultipleChoiceStringFormatter) getRandomIndex(numberOfWords int, indexCorrectWord int) (int, error) {
	if indexCorrectWord > numberOfWords {
		return 0, fmt.Errorf("index correct word %d is out bounds, there are only %d words", indexCorrectWord, numberOfWords)
	}
	index := indexCorrectWord
	for index == indexCorrectWord {
		index := rand.Intn(numberOfWords)
		if indexCorrectWord != index {
			return index, nil
		}
	}
	return 0, fmt.Errorf("getting random index")
}
