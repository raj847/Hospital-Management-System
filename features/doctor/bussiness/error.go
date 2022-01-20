package bussiness

import "errors"

var (
	ErrEmailorPass    = errors.New("username or password is incorrect")
	ErrDuplicateData  = errors.New("account already exist")
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found")
	ErrUnathorized    = errors.New("unauthorized")
	ErrPassNotSame    = errors.New("Password not same")

	ErrEmail = errors.New("incorrect username")
	ErrPass  = errors.New("incorrect password")
)
