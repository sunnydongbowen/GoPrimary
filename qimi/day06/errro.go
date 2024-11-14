package main

import "errors"
var (
	ErrConnectDB = errors.New("connect db failed")
)
type DBError struct {
	msg string
}

func (e *DBError) Error() string {
	return e.msg
}

func ConnectDB(username, password string) (err error) {
    if password !="123"{
		return  ErrConnectDB
	}

	if password=="123"{
		return &DBError{
			msg: "password is weak",
		}
	}
	return nil
}