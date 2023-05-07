package controller

import (
	"net/http"

	"auth_service/lib/data"
)

// Add a user to database
func RegisterHandler(resW http.ResponseWriter, req *http.Request) {
	// Extra error handling should be done on server to prevent malicious attacks
	if _, ok := req.Header["Email"]; !ok {
		resW.WriteHeader(http.StatusBadRequest)
		resW.Write([]byte("Emain Missing."))
		return
	}

	if _, ok := req.Header["Username"]; !ok {
		resW.WriteHeader(http.StatusBadRequest)
		resW.Write([]byte("Username Missing."))
		return
	}

	if _, ok := req.Header["PasswordHash"]; !ok {
		resW.WriteHeader(http.StatusBadRequest)
		resW.Write([]byte("PasswordHash Missing."))
		return
	} 

	if _, ok := req.Header["Fullname"]; !ok {
		resW.WriteHeader(http.StatusBadRequest)
		resW.Write([]byte("Fullname Missing."))
		return
	}

	// Validate, then add user
	isAdded := data.AddUser(req.Header["Email"][0], req.Header["Username"][0], req.Header["PasswordHash"][0], req.Header["Fullname"][0], 0)

	// User was already in database
	if !isAdded {
		resW.WriteHeader(http.StatusConflict)
		resW.Write([]byte("Email or Username aldready exists."))
		return
	}

	resW.WriteHeader(http.StatusOK)
	resW.Write([]byte("User Created."))
}