package repo

import (
	"github.com/erwan690/blog/backend/db"
	"github.com/erwan690/blog/backend/model"
	"github.com/erwan690/blog/backend/payload"
	"gorm.io/gorm"
)

type PostRepo interface {
	GetPosts(filter *payload.FilterPostRequest) ([]*model.Post, error)
	CreatePost(post *model.Post) error
	GetPostByID(id string) (*model.Post, error)
	UpdatePost(id string, post *model.Post) error
	DeletePost(id string) error
}

type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(database *db.DatabaseConnection) PostRepo {
	return &postRepo{
		db: database.DB,
	}
}

func (p *postRepo) GetPosts(filter *payload.FilterPostRequest) ([]*model.Post, error) {
	posts := make([]*model.Post, 0)
	query := p.db.Model(&model.Post{})
	if filter.Title != "" {
		query = query.Where("title LIKE ?", "%"+filter.Title+"%")
	}
	if filter.Content != "" {
		query = query.Where("content LIKE ?", "%"+filter.Content+"%")
	}
	if filter.Category != "" {
		query = query.Where("category LIKE ?", "%"+filter.Category+"%")
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	err := query.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postRepo) CreatePost(post *model.Post) error {
	err := p.db.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *postRepo) GetPostByID(id string) (*model.Post, error) {
	post := &model.Post{}
	err := p.db.First(post, id).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *postRepo) UpdatePost(id string, post *model.Post) error {
	err := p.db.Where("id = ?", id).Updates(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *postRepo) DeletePost(id string) error {
	err := p.db.Delete(&model.Post{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
