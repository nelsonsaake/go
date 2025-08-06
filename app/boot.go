package app

import (
	"fmt"
)

func bootRecover() {
	res := recover()

	fmt.Println("boot failed")
	fmt.Println(res)
}

func Boot() error {

	fmt.Println("Booting app ...")
	defer bootRecover()

	fmt.Println("running setups ..")

	for name, setup := range setups() {

		fmt.Println("setting up", name, "...")

		err := setup.Setup()
		if err != nil {
			return fmt.Errorf("error setting up %s: %v", name, err)
		}

		fmt.Println("setup", name, "complete")
	}

	fmt.Println("setups complete")

	return nil
}
