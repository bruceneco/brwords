package brwords

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/kennygrant/sanitize"
	"strings"
	"sync"
)

var _ Scrapper = (*Scrap)(nil)

type (
	// Scrapper is definition of what the Scrap implements.
	Scrapper interface {
		Word(rawWord string) (Word, error)
		MostSearched(page int) ([]string, error)
	}
	// Scrap is the base instance to make word
	Scrap struct {
		collector *colly.Collector
	}
)

// NewScrap creates a Scrap instance.
func NewScrap() *Scrap {
	return &Scrap{
		collector: colly.NewCollector(colly.AllowURLRevisit()),
	}
}

func (*Scrap) Slug(word string) string {
	noAccentsWord := sanitize.Accents(word)
	noEdgingSpacesWord := strings.TrimSpace(noAccentsWord)
	noUppercasesWord := strings.ToLower(noEdgingSpacesWord)
	noSpacesWord := strings.ReplaceAll(noUppercasesWord, " ", "-")
	return noSpacesWord
}

// Word populates an instance of Word based on extracted content.
func (w *Scrap) Word(rawWord string) (Word, error) {
	sanitizedWord := w.Slug(rawWord)

	c := w.collector.Clone()

	var (
		mu   sync.Mutex
		word Word
	)

	setContent(c, &mu, &word)
	setMeanings(c, &mu, &word)
	setEtymologies(c, &mu, &word)
	setPhrases(c, &mu, &word)
	setSynonyms(c, &mu, &word)

	err := c.Visit(fmt.Sprintf("https://www.dicio.com.br/%s", sanitizedWord))
	if err != nil {
		return Word{}, fmt.Errorf("%w: %s", ErrCantVisit, err)
	}

	return word, nil
}

func setContent(c *colly.Collector, mu *sync.Mutex, word *Word) {
	c.OnHTML("div.title-header > h1", func(e *colly.HTMLElement) {
		mu.Lock()
		defer mu.Unlock()

		word.Content = strings.TrimSpace(strings.Trim(e.Text, "\n"))
	})
}

func setMeanings(c *colly.Collector, mu *sync.Mutex, word *Word) {
	c.OnHTML("p.significado > span:not(.cl):not(.etim)", func(e *colly.HTMLElement) {
		mu.Lock()
		defer mu.Unlock()

		tag := e.DOM.Find("span.tag").Text()

		definition := Meaning{
			Tag:     strings.TrimSuffix(strings.TrimPrefix(tag, "["), "]"),
			Content: strings.TrimSpace(strings.Replace(e.Text, tag, "", 1)),
		}

		word.Meanings = append(word.Meanings, definition)
	})
}

func setEtymologies(c *colly.Collector, mu *sync.Mutex, word *Word) {
	c.OnHTML("p.significado > span.etim", func(e *colly.HTMLElement) {
		mu.Lock()
		defer mu.Unlock()
		word.Etymologies = append(word.Etymologies, e.Text)
	})
}

func setPhrases(c *colly.Collector, mu *sync.Mutex, word *Word) {
	c.OnHTML(".frases > .frase", func(e *colly.HTMLElement) {
		mu.Lock()
		defer mu.Unlock()

		by := e.DOM.Find("em").Text()
		content := strings.ReplaceAll(e.Text, by, "")

		by = strings.TrimPrefix(by, "- ")
		content = strings.ReplaceAll(content, "\n", "")
		content = strings.Trim(content, " ")

		phrase := Phrase{
			By:      by,
			Content: content,
		}

		word.Phrases = append(word.Phrases, phrase)
	})
}
func setSynonyms(c *colly.Collector, mu *sync.Mutex, word *Word) {
	c.OnHTML("p.sinonimos > a", func(e *colly.HTMLElement) {
		mu.Lock()
		defer mu.Unlock()

		word.Synonyms = append(word.Synonyms, e.Text)
	})
}

// MostSearched get a list of most searched words.
// The list is paginated and the page index go from 1 to 50
// (any passed value below or above this range will be trimmed).
func (w *Scrap) MostSearched(page int) ([]string, error) {
	if page <= 0 {
		page = 1
	} else if page > 50 {
		page = 50
	}
	var words []string
	var mu sync.Mutex
	c := w.collector.Clone()

	c.OnHTML("ul.list > li > a", func(e *colly.HTMLElement) {
		mu.Lock()
		defer mu.Unlock()

		words = append(words, strings.TrimSpace(e.Text))
	})

	err := c.Visit(fmt.Sprintf("https://www.dicio.com.br/palavras-mais-buscadas/%d/", page))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCantVisit, err)
	}

	return words, nil
}
