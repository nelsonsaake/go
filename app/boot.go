package app

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func bootRecover() {

	err := recover()

	if err != nil {
		logrus.Println("boot failed:", err)
	}
}

func Boot() error {

	logrus.Println("Booting app ...")
	defer bootRecover()

	logrus.Println("running setups ..")

	for name, setup := range setups() {

		logrus.Println("setting up", name, "...")

		err := setup.Setup()
		if err != nil {
			return fmt.Errorf("error setting up %s: %v", name, err)
		}

		logrus.Println("setup", name, "complete")
	}

	logrus.Println("setups complete")

	return nil
}
