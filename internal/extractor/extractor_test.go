package extractor

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wouove/learning-words/internal/model"
)

func TestSqlLiteExtractorAdapter_Extract(t *testing.T) {
	path := os.Getenv("INPUT_PATH")
	extractor, err := NewSqlLiteExtractorAdapter(path)
	require.NoError(t, err)

	result, err := extractor.Extract()
	expected := model.WordPair{}
	require.NoError(t, err)
	require.Equal(t, expected, result)

}
