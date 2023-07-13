package server

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		Email string `json:"email"`
	}
	GetUserResponse struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	queryParams := r.URL.Query()

	// Access specific query parameters by their key
	email := queryParams.Get("id")

	req = GetUserRequest{
		Email: email,
	}
	return req, nil
}
