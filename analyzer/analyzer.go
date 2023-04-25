package analyzer

type Analyzer struct {
	tokenizer      Tokenizer
	charsetFilters []CharsetFilter
	tokenFilters   []TokenFilter
}

type Tokenizer interface {
	Tokenize(text string) []string
}

type CharsetFilter interface {
	Filter(text string) string
}

type TokenFilter interface {
	Filter(token []string) []string
}
