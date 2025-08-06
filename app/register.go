package app

import "github.com/sirupsen/logrus"

// map[name]Resource
var resources = map[string]Resource{}

func Register(name string, resource Resource) {

	resources[name] = resource
}

func setups() map[string]Setup {

	logrus.Println("getting setups ...")

	var setups = map[string]Setup{}

	for name, resource := range resources {

		setup, ok := resource.(Setup)
		if ok {
			setups[name] = setup

		}
	}

	logrus.Println("found", len(setups), "...")

	return setups
}

func teardowns() map[string]Teardown {

	var teardowns = map[string]Teardown{}

	for name, resource := range resources {

		teardown, ok := resource.(Teardown)
		if ok {
			teardowns[name] = teardown
		}
	}

	return teardowns
}
