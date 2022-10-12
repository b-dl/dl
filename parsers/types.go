package parsers

type Data struct {
}

type ParserData struct {
}

type ParserOptions struct {
}

type Parser interface {
	Init(options *ParserOptions) error
	Parse() (*Data, error)
}
