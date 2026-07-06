package users_transport_http

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUserRe(rw http.ResponseWriter, req *http.Request) {
	var request CreateUserRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {

		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
