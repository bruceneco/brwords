package brwords

// Meaning is the definition of a word
type Meaning struct {
	// Tag is the type of meaning
	Tag string `json:"tipo,omitempty"`
	// Content is the meaning itself
	Content string `json:"conteudo"`
}

// Phrase is an example of word usage.
type Phrase struct {
	// Content is the phrase itself.
	Content string `json:"conteudo"`
	// By is the source of the Phrase.
	By string `json:"autor"`
}

// Word is the core result of an extraction.
type Word struct {
	// Content is the word itself
	Content string `json:"palavra"`
	// Meanings are a list of possible meanings
	Meanings []Meaning `json:"significados"`
	// Synonyms are a list of possible synonyms
	Synonyms []string `json:"sinonimos"`
	// Etymologies are a list of possible etymologies
	Etymologies []string `json:"etimologias"`
	// Phrases are a list of example phrases
	Phrases []Phrase `json:"frases"`
}
