package service

import (
	"github.com/lucasquitan/crud/src/configuration/logger"
	"github.com/lucasquitan/crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*userDomainService) DeleteUser(string) *rest_err.RestErr {
	logger.Info("Init DeleteUser model", zap.String("journey", "DeleteUser"))
	
	return nil
}