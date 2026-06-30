package response

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int8   `json:"age" binding:"required"`
}
