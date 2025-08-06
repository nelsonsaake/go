package app

import (
	"fmt"
	"log"
)

func bootRecover() {

	err := recover()

	if err != nil {
		log.Println("boot failed:", err)
	}
}

func Boot() error {

	log.Println("Booting app ...")
	defer bootRecover()

	log.Println("running setups ..")

	for name, setup := range setups() {

		log.Println("setting up", name, "...")

		err := setup.Setup()
		if err != nil {
			return fmt.Errorf("error setting up %s: %v", name, err)
		}

		log.Println("setup", name, "complete")
	}

	log.Println("setups complete")

	return nil
}
