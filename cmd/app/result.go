package app

import "sync"

type Results struct {
	Mu  *sync.Mutex
	Res []Response
}

func NewResults() *Results {
	return &Results{}
}
