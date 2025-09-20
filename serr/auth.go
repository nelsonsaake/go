package serr

func Unauthorized(f string, v ...any) *ServerError {
	return Message("You're not authorized to perform this action.").WithStatus(401).WithErrorf(f, v...)
}

func JWTExpired(f string, v ...any) *ServerError {
	return Message("Login expired, please login again.").WithStatus(401).WithErrorf(f, v...)
}

func InvalidJWT(f string, v ...any) *ServerError {
	return Message("Oops! You need to log in to access this.").WithStatus(401).WithErrorf(f, v...)
}

func AuthError(f string, v ...any) *ServerError {
	return Message("Something went wrong, please login and try again.").WithStatus(401).WithErrorf(f, v...)
}
