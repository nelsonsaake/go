package sys

import (
	"fmt"
)

func UfwAllow(port string) error {
	err := Run("ufw", "allow", port).Error
	if err != nil {
		return fmt.Errorf("ufw allow %s failed: %v", port, err)
	}
	return nil
}

func UfwAllowTCP(port string) error {
	err := Run("ufw", "allow", port+"/tcp").Error
	if err != nil {
		return fmt.Errorf("ufw allow %s/tcp failed: %v", port, err)
	}
	return nil
}
