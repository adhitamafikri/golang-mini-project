package service

import (
	httpEntity "../entity/http"
	repository "../repository/db"
	"github.com/jinzhu/copier"
)

type PostsService struct {
	postsRepository repository.PostsRepositoryInterface
}

func PostsServiceHandler() PostsService {
	return PostsService{
		postsRepository: repository.PostsRepositoryHandler(),
	}
}

type PostsServiceInterface interface {
	GetAllPosts() []httpEntity.PostsResponse
	GetPostsByUserID(userID int) []httpEntity.PostsResponse
}

func (service *PostsService) GetAllPosts() []httpEntity.PostsResponse {
	posts, _ := service.postsRepository.GetAllPosts(10, 0)
	result := []httpEntity.PostsResponse{}
	copier.Copy(&result, &posts)
	return result
}
