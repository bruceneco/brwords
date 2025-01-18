package brwords

// Meaning is the
type Meaning struct {
	Tag, Content string
}

// Phrase is an example of word usage.
type Phrase struct {
	// Content is the phrase itself.
	Content string
	// By is the source of the Phrase.
	By string
}

// Word is the core result of an extraction.
type Word struct {
	// Content is the word itself
	Content string
	// Meanings are a list of possible meanings
	Meanings []Meaning
	// Synonyms are a list of possible synonyms
	Synonyms []string
	// Etymologies are a list of possible etymologies
	Etymologies []string
	// Phrases are a list of example phrases
	Phrases []Phrase
}
