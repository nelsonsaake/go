package serr

func From(e error) *ServerError {

	err, ok := e.(*ServerError)
	if ok {
		return err
	}

	if e.Error() == "ErrNotFound" {
		return NotFound()
	}

	return New().WithError(e.Error())
}
