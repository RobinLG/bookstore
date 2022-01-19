package factory

import "sync"

var (
	providerMu	sync.RWMutex
	providers	= make(map[string])
)
