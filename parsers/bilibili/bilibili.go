package bilibili

import (
	"github.com/b-dl/dl/parsers"
)

func init() {
	parsers.Register("bilibili", New())
}

func New() parsers.Parser {
	return &parser{}
}

func (p *parser) Parse() (*parsers.Data, error) {
	return nil, nil
}
