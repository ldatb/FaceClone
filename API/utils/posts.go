package utils

import (
	"os"

	"faceclone-api/data"
	"faceclone-api/data/models"

	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

func GetPost(id int64) (bool, *models.Post, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, DBengine, err
	}

	// Get post
	postRequest := new(models.Post)
	has, err := DBengine.Table("post").Where("id = ?", id).Desc("id").Get(postRequest)
	if err != nil {
		return false, postRequest, DBengine, err
	}

	// Post not found
	if !has {
		return false, postRequest, DBengine, nil
	}

	return true, postRequest, DBengine, nil
}

func GetPostMedia(id int64) (bool, *models.PostMedia, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, DBengine, err
	}

	// Get post
	postMediaRequest := new(models.PostMedia)
	has, err := DBengine.Table("post_media").Where("id = ?", id).Desc("id").Get(postMediaRequest)
	if err != nil {
		return false, postMediaRequest, DBengine, err
	}

	// Post not found
	if !has {
		return false, postMediaRequest, DBengine, nil
	}

	return true, postMediaRequest, DBengine, nil
}

func GetPostComments(id int64) (bool, []models.PostComments, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, DBengine, err
	}

	// Get comments
	var comments []models.PostComments
	err = DBengine.Table("post_comments").Where("post_id = ?", id).Desc("id").Find(&comments)
	if err != nil {
		return false, comments, DBengine, err
	}

	// No comments
	if comments == nil {
		return false, comments, DBengine, nil
	}

	return true, comments, DBengine, nil
}

func GetComment(id int64) (bool, *models.PostComments, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, DBengine, err
	}

	// Get user
	commentRequest := new(models.PostComments)
	has, err := DBengine.Table("post_comments").Where("id = ?", id).Desc("id").Get(commentRequest)
	if err != nil {
		return false, commentRequest, DBengine, err
	}

	// Post not found
	if !has {
		return false, commentRequest, DBengine, nil
	}

	return true, commentRequest, DBengine, nil
}

func CreatePostMediaUrl(filename string) (string, error) {
	// Load url from .env
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	// Get image url
	URL := os.Getenv("APP_URL") + "/posts/post/media/" + filename

	return URL, nil
}