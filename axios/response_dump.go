package axios

import (
	"fmt"

	"github.com/nelsonsaake/go/pty"
)

func (res *Response) dump(err error) error {

	if err != nil {
		return err
	}

	var (
		req = res.Request()
	)
	fmt.Println("REQUEST:")

	var (
		method = req.Method
		url    = req.URL.String()
	)
	fmt.Printf("[%s] %s \n\n", method, url)

	fmt.Println("RESPONSE:")

	var (
		code = res.StatusCode()
	)
	fmt.Print("status code: ", code, "\n\n")

	payload, err := res.Json()
	if err != nil {
		return err
	}

	if res.IsArray() || res.IsMap() {
		pty.Println(payload)
	} else {
		fmt.Println(payload)
	}

	return nil
}

func (res *Response) Dump(err error) (*Response, error) {

	err = res.dump(err)
	if err != nil {
		fmt.Println(err)
	}

	return res, err
}
