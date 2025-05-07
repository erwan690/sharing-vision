package service

import (
	"github.com/erwan690/blog/backend/model"
	"github.com/erwan690/blog/backend/payload"
	"github.com/erwan690/blog/backend/repo"
)

type PostService interface {
	GetPosts(filter *payload.FilterPostRequest) ([]*payload.PostResponse, error)
	CreatePost(post *payload.PostRequest) error
	GetPostByID(id string) (*payload.PostResponse, error)
	UpdatePost(id string, post *payload.UpdatePostRequest) error
	DeletePost(id string) error
}

type postService struct {
	repo repo.PostRepo
}

func NewPostService(repo repo.PostRepo) PostService {
	return &postService{
		repo: repo,
	}
}

func (p *postService) GetPosts(filter *payload.FilterPostRequest) ([]*payload.PostResponse, error) {
	data, err := p.repo.GetPosts(filter)
	if err != nil {
		return nil, err
	}
	posts := make([]*payload.PostResponse, 0)
	for _, post := range data {
		posts = append(posts, &payload.PostResponse{
			ID:          post.ID,
			Title:       post.Title,
			Content:     post.Content,
			Category:    post.Category,
			CreatedDate: post.CreatedDate.Format("2006-01-02 15:04:05"),
			UpdatedDate: post.UpdatedDate.Format("2006-01-02 15:04:05"),
			Status:      post.Status,
		})
	}
	return posts, nil
}

func (p *postService) CreatePost(post *payload.PostRequest) error {
	newPost := &model.Post{
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Status:   post.Status,
	}
	err := p.repo.CreatePost(newPost)
	if err != nil {
		return err
	}
	return nil
}

func (p *postService) GetPostByID(id string) (*payload.PostResponse, error) {
	post, err := p.repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	postResponse := &payload.PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		Category:    post.Category,
		CreatedDate: post.CreatedDate.Format("2006-01-02 15:04:05"),
		UpdatedDate: post.UpdatedDate.Format("2006-01-02 15:04:05"),
		Status:      post.Status,
	}

	return postResponse, nil
}

func (p *postService) UpdatePost(id string, post *payload.UpdatePostRequest) error {
	_, err := p.repo.GetPostByID(id)
	if err != nil {
		return err
	}
	err = p.repo.UpdatePost(id, &model.Post{
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Status:   post.Status,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *postService) DeletePost(id string) error {
	_, err := p.repo.GetPostByID(id)
	if err != nil {
		return err
	}

	err = p.repo.UpdatePost(id, &model.Post{
		Status: "Trash",
	})
	if err != nil {
		return err
	}

	return nil
}
