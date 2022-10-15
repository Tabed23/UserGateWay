package utils

import "errors"

//Common Errors for HTTP
var (
	BadRequest = errors.New("bad request, Invalid json request body") //Input json is wrong
	ConnectionDenied = errors.New("connection denied to RedShift Cluster") //DB Connection file 
)