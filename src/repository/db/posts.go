package db

import (
	dbEntity "../../entity/db"
	connection "../../util/helper/mysqlconnection"
	"github.com/jinzhu/gorm"
)

type PostsRepository struct {
	DB gorm.DB
}

func PostsRepositoryHandler() *PostsRepository {
	return &PostsRepository{DB: *connection.GetConnection()}
}

type PostsRepositoryInterface interface {
	GetAllPosts(limit int, offset int) ([]dbEntity.Posts, error)
	GetPostsByUserID(userId int, limit int, offset int) ([]dbEntity.Posts, error)
}

func (repository *PostsRepository) GetAllPosts(limit int, offset int) ([]dbEntity.Posts, error) {
	posts := []dbEntity.Posts{}

	query := repository.DB.Table("posts")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&posts)

	return posts, query.Error
}

func (repository *PostsRepository) GetPostsByUserID(userId int, limit int, offset int) ([]dbEntity.Posts, error) {
	posts := []dbEntity.Posts{}

	query := repository.DB.Table("posts")
	query = query.Where("user_id=?", userId).Limit(limit).Offset(offset)
	query = query.Find(&posts)

	return posts, query.Error
}
