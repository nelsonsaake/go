package crypt

func Verify(promp string, hash string) (bool, error) {

	hashed, err := Hash(promp)
	if err != nil {
		return false, err
	}

	isVerified := hashed == hash

	return isVerified, nil
}
