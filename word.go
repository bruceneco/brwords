package brwords

type Meaning struct {
	Tag, Content string
}
type Phrase struct {
	Content, By string
}
type Word struct {
	Content     string
	Meanings    []Meaning
	Synonyms    []string
	Etymologies []string
	Phrases     []Phrase
}
