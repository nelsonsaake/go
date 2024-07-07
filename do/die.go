package do

// Die: panic if err is not nil
func Die(err any) {

	switch err := err.(type) {
	case bool:
		if !err {
			panic("something went wrong")
		}
	case string:
		panic(err)
	}

	if err != nil {
		panic(err)
	}
}
