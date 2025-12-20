package settings

func Teardown() error {
	return obj.SaveToFile(filePath)
}
