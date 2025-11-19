package ufs

// GetRootDir: the root dir is a the dir that contains the go.mod file
func GetRootDir() (string, error) {
	return Root()
}
