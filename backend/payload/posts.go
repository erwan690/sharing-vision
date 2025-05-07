package payload

type PostRequest struct {
	Title    string `json:"title" binding:"required,min=20"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3"`
	Status   string `json:"status" binding:"required,oneof=Publish Draft Trash"`
}

type UpdatePostRequest struct {
	Title    string `json:"title,omitempty" binding:"min=20"`
	Content  string `json:"content,omitempty" binding:"min=200"`
	Category string `json:"category,omitempty" binding:"min=3"`
	Status   string `json:"status,omitempty" binding:"oneof=Publish Draft Trash"`
}

type PostResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	Status      string `json:"status"`
}

type FilterPostRequest struct {
	Title    string `form:"title,omitempty"`
	Content  string `form:"content,omitempty"`
	Category string `form:"category,omitempty"`
	Status   string `form:"status,omitempty"`
	Limit    int    `form:"limit,omitempty"`
	Offset   int    `form:"offset,omitempty"`
}

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
