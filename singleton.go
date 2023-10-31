package main

import (
	"sync"
)

type Singleton struct {
    data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{data: "Singleton instance initialized"}
    })
    return instance
}

func (s *Singleton) GetData() string {
    return s.data
}
