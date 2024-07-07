package do

// Die: panic if err is not nil
func Die(err any) {

	switch err := err.(type) {
	case error:
		if err != nil {
			panic(err)
		}
	case bool:
		if !err {
			panic("something went wrong")
		}
	case string:
		panic(err)
	default:
		panic("something went wrong")
	}
}
