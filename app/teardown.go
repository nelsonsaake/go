package app

type Teardown interface {
	Teardown() error
}
