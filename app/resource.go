package app

type Teardown interface {
	Teardown() error
}

type Setup interface {
	Setup() error
}

type Resource interface{}
