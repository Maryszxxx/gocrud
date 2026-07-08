package service

import (
	"github.com/Maryszxxx/gocrud.git/src/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/config/rest_err"
	"github.com/Maryszxxx/gocrud.git/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {

	logger.Info("Init LoginUser model.",
		zap.String("journey", "LoginUser"))

	userDomain.EncryptPassword()

	user, err := ud.userRepository.FindUserByEmailAndPassword(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "loginUser"))
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		logger.Error("Error trying to generate token",
			err,
			zap.String("journey", "loginUser"))
		return nil, "", err
	}

	logger.Info(
		"LoginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "LoginUser"))
	return user, token, nil
}
