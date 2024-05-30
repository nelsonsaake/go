package do

import (
	_ "bytes"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func recorderPayload(r *httptest.ResponseRecorder) []byte {
	return r.Body.Bytes()
}

func responsePayload(r *http.Response) []byte {
	raw, _ := io.ReadAll(r.Body)
	return raw
}

func requestPayload(r *http.Request) []byte {
	raw, _ := io.ReadAll(r.Body)
	return raw
}

// Payload: copies out the from httptest.ResponseRecorder and http.Response,
// if something else is supplied it return an empty array of bytes
func Payload(resp any) []byte {

	switch resp := resp.(type) {
	case *httptest.ResponseRecorder:
		return recorderPayload(resp)
	case *http.Response:
		return responsePayload(resp)
	case *gin.Context:
		return requestPayload(resp.Request)
	default:
		return []byte{}
	}
}
