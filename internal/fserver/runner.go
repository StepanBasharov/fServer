package fserver

import "sync"

type ServerComponent interface {
	Run(group *sync.WaitGroup) error
}
