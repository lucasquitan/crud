package service

import (
	"github.com/lucasquitan/crud/src/configuration/logger"
	"github.com/lucasquitan/crud/src/configuration/rest_err"
	"github.com/lucasquitan/crud/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser model", zap.String("journey", "UpdateUser"))
	
return nil
}