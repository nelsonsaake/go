package meta

func Teardown() error {
	return obj.SaveToFile(filePath)
}
