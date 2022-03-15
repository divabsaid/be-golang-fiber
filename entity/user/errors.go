package user

import "errors"

var (
	NOT_FOUND                = errors.New("Requested Data Not Found")
	LOGIN_FAILED             = errors.New("Login Failed, Username or Password not match")
	REGISTER_FAILED          = errors.New("Register Failed, Username or Email Already Exist")
	REQUEST_BODY_NOT_VALID   = errors.New("Request Body is not Valid")
	DELETE_FAILED            = errors.New("Delete Failed, Data not Found")
	UPDATE_FAILED            = errors.New("Update Failed, Data not Found")
	UPDATE_FAILED_DUPLICATE  = errors.New("Update Failed, Username or Email Already Exist")
	LOGIN_FAILED_INACTIVE    = errors.New("Login Failed, User Inactive")
	USERNAME_EMAIL_NOT_MATCH = errors.New("Username and Email not Match")
)
