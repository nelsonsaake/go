package do

import (
	"net/http"
)

// AuthReq: attach token to request
func AuthReq(req *http.Request, token string) {

	// add access token to its header as part of Authorization
	bearer := "Bearer " + token
	req.Header.Add("Authorization", bearer)
}
