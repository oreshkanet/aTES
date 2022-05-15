package models

import "errors"

var ErrorInvalidAccessToken = errors.New("invalid auth authorizer")
var ErrorUserDoesNotExist = errors.New("user does not exist")
var ErrorUserAlreadyExists = errors.New("user already exists")
