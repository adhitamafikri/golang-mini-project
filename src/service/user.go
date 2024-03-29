package service

import (
	"sync"

	APIEntity "../entity/api"
	dbEntity "../entity/db"
	httpEntity "../entity/http"
	repositoryAPI "../repository/api"
	repository "../repository/db"
	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepository    repository.UserRepositoryInterface
	userRepositoryAPI repositoryAPI.FriendAPIRepositoryInterface
}

func UserServiceHandler() UserService {
	return UserService{
		userRepository:    repository.UserRepositoryHandler(),
		userRepositoryAPI: repositoryAPI.ThirdPartyAPIHandler(),
	}
}

type UserServiceInterface interface {
	GetUserByID(id int) httpEntity.UserDetailResponse
	GetAllUser(page int, count int) []httpEntity.UserResponse
	UpdateUserByID(id int) bool
}

func (service *UserService) GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse {
	waitGroup.Add(1)
	user := &dbEntity.User{}
	go service.userRepository.GetUserByID(id, user, waitGroup)

	waitGroup.Add(1)
	friend := &APIEntity.FriendResponse{}
	go service.userRepositoryAPI.GetFriendID(id, friend, waitGroup)

	waitGroup.Wait()

	result := &httpEntity.UserDetailResponse{}
	if user != nil {
		copier.Copy(result, user)
		if nil != user.UserStatus {
			result.Status = &user.UserStatus.Name
		}
		if friend != nil {
			result.Avatar = &friend.Data.Avatar
		}
	}
	return result
}

func (service *UserService) GetAllUser(page int, count int) []httpEntity.UserResponse {
	users, _ := service.userRepository.GetUsersList(page, count)
	result := []httpEntity.UserResponse{}
	copier.Copy(&result, &users)
	return result
}

func (service *UserService) UpdateUserByID(id int, payload httpEntity.UserRequest) bool {
	user := &dbEntity.User{}
	user.Name = payload.Name
	user.IDCardNumber = payload.IDCardNumber
	user.Address = payload.Address
	err := service.userRepository.UpdateUserByID(id, user)
	if nil != err {
		return false
	}
	return true
}
