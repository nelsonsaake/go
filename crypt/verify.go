package crypt

func Verify(promp string, hash string) bool {
	return Hash(promp) == hash
}
