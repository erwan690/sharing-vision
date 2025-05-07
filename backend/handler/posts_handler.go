package handler

import (
	"github.com/erwan690/blog/backend/payload"
	"github.com/erwan690/blog/backend/service"
	"github.com/gin-gonic/gin"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *postHandler {
	return &postHandler{
		service: service,
	}
}

// GetAllPosts godoc
// @Summary Get all posts
// @Description Get all posts
// @Tags posts
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param title query string false "Title"
// @Param content query string false "Content"
// @Param category query string false "Category"
// @Param status query string false "Status"
// @Success 200 {object} payload.SuccessResponse{Data=[]payload.PostResponse}
// @Failure 400 {object} payload.ErrorResponse
// @Failure 500 {object} payload.ErrorResponse
// @Router /article [get]
func (h *postHandler) GetAllPosts(c *gin.Context) {
	filter := &payload.FilterPostRequest{}
	if err := c.ShouldBindQuery(filter); err != nil {
		c.JSON(400, payload.ErrorResponse{
			Status:  false,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}
	// Validate filter
	if filter.Limit < 0 {
		filter.Limit = 10
	}

	if filter.Offset < 0 {
		filter.Offset = 0
	}

	posts, err := h.service.GetPosts(filter)
	if err != nil {
		c.JSON(500, payload.ErrorResponse{
			Status:  false,
			Message: "Failed to retrieve posts",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(200, payload.SuccessResponse{
		Status:  true,
		Message: "Posts retrieved successfully",
		Data:    posts,
	})
}

// GetPostByID godoc
// @Summary Get a post by ID
// @Description Get a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} payload.SuccessResponse{Data=payload.PostResponse}
// @Failure 400 {object} payload.ErrorResponse
// @Failure 500 {object} payload.ErrorResponse
// @Router /article/{id} [get]
func (h *postHandler) GetPostByID(c *gin.Context) {
	id := c.Param("id")

	post, err := h.service.GetPostByID(id)
	if err != nil {
		c.JSON(500, payload.ErrorResponse{
			Status:  false,
			Message: "Failed to retrieve post",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(200, payload.SuccessResponse{
		Status:  true,
		Message: "Post retrieved successfully",
		Data:    post,
	})
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body payload.PostRequest true "Post"
// @Success 201 {object} payload.SuccessResponse{Data=payload.PostResponse}
// @Failure 400 {object} payload.ErrorResponse
// @Failure 500 {object} payload.ErrorResponse
// @Router /article [post]
func (h *postHandler) CreatePost(c *gin.Context) {
	post := &payload.PostRequest{}
	if err := c.ShouldBindJSON(post); err != nil {
		c.JSON(400, payload.ErrorResponse{
			Status:  false,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	err := h.service.CreatePost(post)
	if err != nil {
		c.JSON(500, payload.ErrorResponse{
			Status:  false,
			Message: "Failed to create post",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(201, payload.SuccessResponse{
		Status:  true,
		Message: "Post created successfully",
		Data:    post,
	})
}

// UpdatePost godoc
// @Summary Update a post
// @Description Update a post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body payload.PostRequest true "Post"
// @Success 200 {object} payload.SuccessResponse{Data=payload.PostResponse}
// @Failure 400 {object} payload.ErrorResponse
// @Failure 500 {object} payload.ErrorResponse
// @Router /article/{id} [put]
func (h *postHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	body := &payload.UpdatePostRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(400, payload.ErrorResponse{
			Status:  false,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	err := h.service.UpdatePost(id, body)
	if err != nil {
		c.JSON(500, payload.ErrorResponse{
			Status:  false,
			Message: "Failed to update post",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(200, payload.SuccessResponse{
		Status:  true,
		Message: "Post updated successfully",
		Data:    body,
	})
}

// DeletePost godoc
// @Summary Delete a post
// @Description Delete a post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} payload.SuccessResponse
// @Failure 400 {object} payload.ErrorResponse
// @Failure 500 {object} payload.ErrorResponse
// @Router /article/{id} [delete]
func (h *postHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeletePost(id)
	if err != nil {
		c.JSON(500, payload.ErrorResponse{
			Status:  false,
			Message: "Failed to delete post",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(200, payload.SuccessResponse{
		Status:  true,
		Message: "Post deleted successfully",
		Data:    nil,
	})
}
