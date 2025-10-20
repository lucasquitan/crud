package service

import (
	"fmt"

	"github.com/lucasquitan/crud/src/configuration/logger"
	"github.com/lucasquitan/crud/src/configuration/rest_err"
	"github.com/lucasquitan/crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))
	
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}