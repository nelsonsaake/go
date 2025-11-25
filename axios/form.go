package axios

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// FormBuilder helps build multipart form requests for the Client.
type FormBuilder struct {
	client *Client
	values map[string]string
	files  []formFile
}

type formFile struct {
	field string
	path  string
}

// Form creates a new FormBuilder bound to this client.
func (c *Client) Form() *FormBuilder {
	return &FormBuilder{
		client: c,
		values: make(map[string]string),
		files:  []formFile{},
	}
}

// With adds a key/value field to the form.
func (f *FormBuilder) With(name, value string) *FormBuilder {
	f.values[name] = value
	return f
}

// WithFile adds a file to be uploaded for the given form field.
func (f *FormBuilder) WithFile(field, path string) *FormBuilder {
	f.files = append(f.files, formFile{field: field, path: path})
	return f
}

// Post builds the multipart/form-data request and delegates to the underlying client's Do method.
func (f *FormBuilder) Post(path string) (*Response, error) {
	// build body
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	// add form values
	for k, v := range f.values {
		if err := writer.WriteField(k, v); err != nil {
			return nil, err
		}
	}

	// add files
	for _, fi := range f.files {
		file, err := os.Open(fi.path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(fi.field, filepath.Base(fi.path))
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(part, file); err != nil {
			return nil, err
		}
	}

	// close writer to set terminating boundary
	if err := writer.Close(); err != nil {
		return nil, err
	}

	url, err := f.client.Url(path)
	if err != nil {
		return nil, err
	}

	// create http request
	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return nil, err
	}

	// set content-type to multipart with boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// use underlying axios client's Do
	return f.client.DoForm(req)
}
