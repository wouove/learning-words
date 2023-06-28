package extractor

import (
	"github.com/stretchr/testify/require"
	"github.com/wouove/learning-words/internal/model"
	"testing"
)

func TestSqlLiteExtractorAdapter_Extract(t *testing.T) {
	path := "/Users/wouteroverbeek/KoboReader.sqlite"
	extractor, err := NewSqlLiteExtractorAdapter(path)
	require.NoError(t, err)

	result, err := extractor.Extract()
	expected := model.WordPair{}
	require.NoError(t, err)
	require.Equal(t, expected, result)

}
