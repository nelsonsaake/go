package sys

import (
	"fmt"
)

func UfwAllow(port string) error {
	out, err := Run("ufw", "allow", port)
	if err != nil {
		return fmt.Errorf("ufw allow %s failed: %v: %s", port, err, out)
	}
	return nil
}

func UfwAllowTCP(port string) error {
	out, err := Run("ufw", "allow", port+"/tcp")
	if err != nil {
		return fmt.Errorf("ufw allow %s/tcp failed: %v: %s", port, err, out)
	}
	return nil
}
