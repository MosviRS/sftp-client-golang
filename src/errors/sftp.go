package errors

import "errors"

var (
	ErrorOpenConnection  = errors.New("Error opening connection")
	ErrorParsePrivateKey = errors.New("Error parsing private key")
)
