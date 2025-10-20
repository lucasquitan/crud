package view

import (
	"github.com/lucasquitan/crud/src/model"
	"github.com/lucasquitan/crud/src/model/response"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID: "",
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}