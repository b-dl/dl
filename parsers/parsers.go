package parsers

import "sync"

var lock sync.RWMutex
var parseMap = make(map[string]Parser)

func Register(key string, p Parser) {
	lock.Lock()
	parseMap[key] = p
	lock.Unlock()
}

func (p *ParserData) Init(options *ParserOptions) error {
	return nil
}
