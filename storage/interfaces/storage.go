package interfaces

// Saver: takes in a base64 string, store it and return the path
type Saver interface {
	Save(b64content string) (path string, err error)
}

type SaverAs interface {
	SaveAs(path, b64content string) (err error)
}

// Deleter: takes in location of a file, deletes it returns an error if something goes wrong
type Deleter interface {
	Delete(path string) (err error)
}

// Storage: saves files and can delete them
type Storage interface {
	Saver
	Deleter
}
