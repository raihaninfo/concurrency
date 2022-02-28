package main

type (
	Server interface {
		Start() error
		Stop()
	}
)
