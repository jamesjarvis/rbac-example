package service

import "errors"

var Error_UNAUTHORISED = errors.New("user unauthorised to perform action")

var Error_NOTFOUND = errors.New("item not found")
