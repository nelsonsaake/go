package do

// Die: panic if err is not nil
func Die(err error) {
	if err != nil {
		panic(err)
	}
}
