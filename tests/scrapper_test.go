package scrapper_test

import (
	"fmt"
	"github.com/bruceneco/brwords"
	"github.com/segmentio/encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWordScrapping(t *testing.T) {
	service := brwords.NewScrap()
	type testCase struct {
		word     string
		expected brwords.Word
		err      error
	}

	testCases := []testCase{
		{
			word:     "Cizânia",
			expected: readWordJSON(t, service.Slug("Cizânia")),
		},
		{
			word:     "Estirpe",
			expected: readWordJSON(t, service.Slug("Estirpe")),
		},
		{
			word:     "Sandice",
			expected: readWordJSON(t, service.Slug("Sandice")),
		},
		{
			word:     "ablublublé",
			expected: brwords.Word{},
			err:      brwords.ErrCantVisit,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.word, func(t *testing.T) {
			t.Parallel()
			w, err := service.Word(testCase.word)
			assert.ErrorIs(t, err, testCase.err)
			assert.Equal(t, testCase.expected, w)
		})
	}
}
func readWordJSON(t *testing.T, slug string) brwords.Word {
	file, err := os.ReadFile(fmt.Sprintf("testdata/%s.json", slug))
	if err != nil {
		t.Fatal(err)
	}
	var word brwords.Word
	err = json.Unmarshal(file, &word)
	if err != nil {
		t.Fatal(err)
	}
	return word
}
