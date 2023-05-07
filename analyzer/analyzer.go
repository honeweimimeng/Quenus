package analyzer

type Analyzer struct {
	tokenizer      Tokenizer
	charsetFilters []CharsetFilter
	tokenFilters   []TokenFilter
}

func (a *Analyzer) Analyze(text string) []string {
	for _, charFilter := range a.charsetFilters {
		text = charFilter.Filter(text)
	}
	tokens := a.tokenizer.Tokenize(text)
	for _, tokenFilter := range a.tokenFilters {
		tokens = tokenFilter.Filter(tokens)
	}
	return tokens
}

func (a *Analyzer) AnalyzeBin(blob any) []any {
	for _, charFilter := range a.charsetFilters {
		blob = charFilter.FilterBin(blob)
	}
	tokens := a.tokenizer.TokenizeBin(blob)
	for _, tokenFilter := range a.tokenFilters {
		tokens = tokenFilter.FilterBin(tokens)
	}
	return tokens
}

type Tokenizer interface {
	Tokenize(text string) []string
	TokenizeBin(text any) []any
}

type CharsetFilter interface {
	Filter(text string) string
	FilterBin(text any) any
}

type TokenFilter interface {
	Filter(token []string) []string
	FilterBin(token []any) []any
}
