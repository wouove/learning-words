package loader

import (
	"reflect"
	"testing"

	"github.com/wouove/learning-words/internal/model"
)

func TestMultipleChoiceStringFormatter_selectRandomWrongTranslations(t *testing.T) {
	type fields struct {
		header          []string
		numberOfAnswers int
	}
	type args struct {
		words            []model.WordPair
		indexCorrectWord int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			"base",
			fields{
				[]string{"q", "a"},
				3,
			},
			args{
				[]model.WordPair{
					{
						Word:                "House",
						Translation:         "Huis",
						WordLanguage:        "English",
						TranslationLanguage: "Dutch",
					},
					{
						Word:                "Apple",
						Translation:         "Appel",
						WordLanguage:        "English",
						TranslationLanguage: "Dutch",
					},
					{
						Word:                "Dog",
						Translation:         "Hond",
						WordLanguage:        "English",
						TranslationLanguage: "Dutch",
					},
				},
				0,
			},
			[]string{"Appel", "Hond"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MultipleChoiceStringFormatter{
				header:          tt.fields.header,
				numberOfAnswers: tt.fields.numberOfAnswers,
			}
			if got, _ := s.selectRandomWrongTranslations(tt.args.words, tt.args.indexCorrectWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectRandomWrongTranslations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleChoiceStringFormatter_TransformToString(t *testing.T) {
	type args struct {
		words      []model.WordPair
		skipHeader bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"happy flow",
			args{
				[]model.WordPair{
					{
						Word:                "House",
						Translation:         "Huis",
						WordLanguage:        "English",
						TranslationLanguage: "Dutch",
					},
					{
						Word:                "Apple",
						Translation:         "Appel",
						WordLanguage:        "English",
						TranslationLanguage: "Dutch",
					},
					{
						Word:                "Dog",
						Translation:         "Hond",
						WordLanguage:        "English",
						TranslationLanguage: "Dutch",
					},
				},
				false,
			},
			5,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMultipleChoiceStringFormatter(3)
			got, err := s.TransformToString(tt.args.words, tt.args.skipHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.args.words)+1 {
				t.Errorf("TransformToString() got = %v, want %v", len(got), len(tt.args.words)+1)
			}
			if len(got[0]) != tt.want {
				t.Errorf("TransformToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
