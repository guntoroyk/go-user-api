package entity

// User is the struct for a user
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required,min=3,max=255"`
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password,omitempty" validate:"required,min=3,max=255"`
	Phone     string `json:"phone" validate:"required,min=9,max=20"`
	Role      Role   `json:"role" validate:"required,min=3,max=255"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// UserFilter is the struct for filter get user
type UserFilter struct {
	ID       int
	Username string
}

// Role is the type for user role
type Role string

// Role const
const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

// ExcludePassword will exclude password from user
func (u *User) ExcludePassword() *User {
	u.Password = ""
	return u
}
