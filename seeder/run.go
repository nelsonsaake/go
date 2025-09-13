package seeder

import (
	"fmt"
	"log"
)

func RunAll() error {
	for name, seeder := range Registry.GetAll() {

		log.Printf("running %s seeder ...\n", name)

		if seeder == nil {
			return fmt.Errorf("%s seeder not found", name)
		}

		err := seeder.Seed(10)
		if err != nil {
			return fmt.Errorf("error running %s seeder: %w", name, err)
		}
	}
	return nil
}
