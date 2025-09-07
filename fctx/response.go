package fctx

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Data(v any) func(*Response) {

	return func(r *Response) {
		r.Data = v
		r.Success = true
	}
}

func Message(v string) func(*Response) {

	return func(r *Response) {
		r.Message = v
		r.Success = true
	}
}
