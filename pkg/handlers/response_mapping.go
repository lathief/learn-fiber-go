package handlers

import (
	"net/http"
)

func GetStatusMsg(status int) string {
	switch status {
	case http.StatusCreated:
		return "Created"
	case http.StatusAccepted:
		return "Accepted"
	case http.StatusBadRequest:
		return "Bad Request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusForbidden:
		return "Forbidden"
	case http.StatusNotFound:
		return "Not Found"
	case http.StatusInternalServerError:
		return "Internal Server Error"
	case http.StatusServiceUnavailable:
		return "Service Unavailable"
	default:
		return "Success"
	}
}
