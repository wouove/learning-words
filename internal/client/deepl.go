package client

import (
	"context"
	"fmt"

	"github.com/DaikiYamakawa/deepl-go"
	"github.com/wouove/learning-words/internal/model"
)

type Client interface {
	Translate(word model.WordPair) (model.WordPair, error)
}

type DeepLClient struct {
	client *deepl.Client
}

var DeepLLanguage = map[string]string{
	"Dutch":   "NL",
	"English": "EN",
}

func NewDeepLClient() DeepLClient {
	sdkClient, err := deepl.New("https://api-free.deepl.com", nil)
	if err != nil {
		fmt.Printf("setting up deepL client: %w", err)
	}
	return DeepLClient{client: sdkClient}
}

func (c DeepLClient) Translate(word model.WordPair) (model.WordPair, error) {
	response, err := c.client.TranslateSentence(
		context.Background(),
		word.Word,
		DeepLLanguage[word.WordLanguage],
		DeepLLanguage[word.TranslationLanguage],
	)
	if err != nil {
		return model.WordPair{}, fmt.Errorf("getting deepl translation for word %s, from language %s to language %s: %w", word.Word, word.WordLanguage, word.TranslationLanguage, err)
	}
	word.Translation = response.Translations[0].Text
	return word, nil
}
