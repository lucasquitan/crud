package service

import (
	"github.com/lucasquitan/crud/src/configuration/logger"
	"github.com/lucasquitan/crud/src/configuration/rest_err"
	"github.com/lucasquitan/crud/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) FindUser (string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUser", zap.String("journey", "FindUser"))
	return nil, nil
}