package settings

func Save() error {
	return obj.SaveToFile(filePath)
}

func Teardown() error {
	return Save()
}
