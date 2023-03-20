package repository

import (
	"errors"
)

var (
	NotFoundErr     = errors.New("not_found_error")
	UnexpectedErr   = errors.New("unexpected_error")
	JsonDecodingErr = errors.New("json_decoding_error")
	JsonEncodingErr = errors.New("json_decoding_error")
)
