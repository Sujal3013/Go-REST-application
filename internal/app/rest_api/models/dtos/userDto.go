package dtos

import(
	"example/rest-api/internal/app/rest_api/entities"
)

type UserResponse struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type GetAllUsersResponse struct{
	Users []*UserResponse `json:"users"`
}

type CreateUserRequest struct {
 FirstName   string `json:"first_name" binding:"required,min=3,max=50"`
 LastName    string `json:"last_name" binding:"required,min=3,max=50"`
 Email       string `json:"email" binding:"required,email,max=254"`
 PhoneNumber string `json:"phone_number" binding:"required"`
}

type UpdateUserRequest struct {
 FirstName   string `json:"first_name" binding:"required,min=3,max=50"`
 LastName    string `json:"last_name" binding:"required,min=3,max=50"`
 Email       string `json:"email" binding:"required,email,max=254"`
 PhoneNumber string `json:"phone_number" binding:"required"`
}

type CreateUserResponse struct {
 FirstName string `json:"first_name" binding:"required,min=3,max=50"`
 LastName  string `json:"last_name" binding:"required,min=3,max=50"`
 Email     string `json:"email" binding:"required,email,max=254"`
 Message   string `json:"message" binding:"required"`
}

func (r *GetAllUsersResponse) MapUsersResponse(users []*entities.User) {
 for _, users := range users {
  user := &UserResponse{
   FirstName:   users.FirstName,
   LastName:    users.LastName,
   Email:       users.Email,
   PhoneNumber: users.PhoneNumber,
  }
  r.Users = append(r.Users, user)
 }
}

func (r *UserResponse) MapUserResponse(user *entities.User) {
	r.FirstName = user.FirstName
	r.LastName = user.LastName
	r.Email = user.Email
	r.PhoneNumber = user.PhoneNumber
}


func (r *CreateUserRequest) ToUser() *entities.User{
	return &entities.User{
		FirstName: r.FirstName,
		LastName: r.LastName,
		Email: r.Email,
		PhoneNumber: r.PhoneNumber,
	}
}

func (r *UpdateUserRequest) ToUser() *entities.User{
	return &entities.User{
		FirstName: r.FirstName,
		LastName: r.LastName,
		Email: r.Email,
		PhoneNumber: r.PhoneNumber,
	}
}

func (r *CreateUserResponse) FromUser(user *entities.User) *CreateUserResponse{
	return &CreateUserResponse{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Message: "User created successfully",}
}