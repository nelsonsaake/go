package do

import "github.com/nelsonsaake/go/str"

// Die: panic if err is not nil
func Die(err any) {

	switch err := err.(type) {
	case bool:
		if !err {
			panic("something went wrong")
		} else {
			return
		}
	case string:
		if !str.IsEmpty(err) {
			panic(err)
		} else {
			return
		}
	}

	if err != nil {
		panic(err)
	}
}
