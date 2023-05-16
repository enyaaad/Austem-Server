package auth

import "errors"

var ErrInvalidAccessToken = errors.New("invalid auth token")
var ErrUserDoesNotExist = errors.New("user does not exist")