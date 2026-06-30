package request

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=4,max=90"`
	Password string `json:"password" binding:"required,min=6,max=100,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz"`
	Email    string `json:"email" binding:"required,email"`
	Age      int8   `json:"age" binding:"required,min=1,max=140"`
}
