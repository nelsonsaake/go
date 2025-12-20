package kvs

func Teardown() error {
	return obj.SaveToFile(filePath)
}
