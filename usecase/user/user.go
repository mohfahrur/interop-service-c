package user

import (
	"log"

	userD "github.com/mohfahrur/interop-service-c/domain/database"
	"github.com/mohfahrur/interop-service-c/entity"
)

type UserAgent interface {
	GetUser(userID string) (userData entity.User, err error)
}

type UserUsecase struct {
	userDomain userD.UserDomain
}

func NewUserUsecase(
	userD userD.UserDomain) *UserUsecase {

	return &UserUsecase{
		userDomain: userD}
}

func (uc *UserUsecase) GetUser(userID string) (userData *entity.User, err error) {

	userData, err = uc.userDomain.GetUserInfo(userID)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
