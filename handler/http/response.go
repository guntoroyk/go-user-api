package http

import "github.com/guntoroyk/go-user-api/entity"

// HttpResponse is a response struct for http handler
type HttpResponse struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

// TransformUsers will transform user data
func TransformUsers(users []*entity.User) []*entity.User {
	for _, user := range users {
		user.Password = ""
	}

	return users
}
