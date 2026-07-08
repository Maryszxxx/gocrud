package request

type UserLogin struct {
	Password string `json:"password" binding:"required,min=6,max=100,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz"`

	Email string `json:"email" binding:"required,email"`
}
