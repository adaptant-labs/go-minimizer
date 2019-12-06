package minimizers

import (
	"fmt"
	"sync"
)

type Minimizer func(level MinimizationLevel, data interface{}) interface{}

var TagMap = map[string]Minimizer{
	"name":	MinimizeName,
	"email": MinimizeEmail,
}

var tagMapLock sync.RWMutex

func AddCustomMinimizer(m Minimizer, tag string) error {
	tagMapLock.Lock()
	defer tagMapLock.Unlock()

	if TagMap[tag] == nil {
		TagMap[tag] = m
	} else {
		return fmt.Errorf("minimizer for %s already defined", tag)
	}

	return nil
}

func DeleteCustomMinimizer(tag string) {
	tagMapLock.Lock()
	defer tagMapLock.Unlock()

	TagMap[tag] = nil
}
