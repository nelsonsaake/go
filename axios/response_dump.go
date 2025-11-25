package axios

import (
	"fmt"

	"github.com/nelsonsaake/go/pty"
)

func (res *Response) Dump() error {

	if res == nil {
		return nil
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
