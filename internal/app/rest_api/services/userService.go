package services

import(
	"database/sql"
	"errors"
	"example/rest-api/internal/app/rest_api/models"
	"example/rest-api/internal/app/rest_api/models/dtos"
	"example/rest-api/internal/app/rest_api/repositories"
	"net/http"
)

type User struct{
	userRepo *repositories.User
}

func NewUserService(user *repositories.User) *User{
	return &User{
		userRepo: user,
	}
}

func (us *User) GetAllUsers()(*dtos.GetAllUsersResponse,*models.ErrorResponse){
	response:=&dtos.GetAllUsersResponse{}
	queriedUsers,err:=us.userRepo.GetAllUsers()
	if err!=nil{
		return nil,&models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal server error",
		}	
	}
	response.MapUsersResponse(queriedUsers)
	return response,nil
}

func (us *User) GetUser(ID int)(*dtos.UserResponse,*models.ErrorResponse){
	reponse:=&dtos.UserResponse{}
	queriedUser,err:=us.userRepo.FindById(ID)
	if err!=nil{
		if errors.Is(err,sql.ErrNoRows){
			return nil,&models.ErrorResponse{
				Code: http.StatusNotFound,
				Message: "User not found",
			}
		}
		return nil,&models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal server error",
		}	
	}
	reponse.MapUserResponse(queriedUser)
	return reponse,nil
}


func (us *User) DeleteUser(userId int)*models.ErrorResponse{
	user,err:=us.userRepo.FindById(userId)
	if err!=nil{
		if errors.Is(err,sql.ErrNoRows){
			return &models.ErrorResponse{
				Code: http.StatusNotFound,
				Message: "User not found",
			}
		}
		return &models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal server error",
		}	
	}
	err=us.userRepo.Delete(user.ID)
	if err!=nil{
		return &models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal server error",
		}	
	}	
	return nil
}

func (us *User) CreateUser(createUserRequest *dtos.CreateUserRequest) (*dtos.CreateUserResponse, *models.ErrorResponse) {
 userResponse := &dtos.CreateUserResponse{}

 errEmail := us.checkIfEmailExists(createUserRequest.Email)
 if errEmail != nil {
  return nil, errEmail
 }

 user := createUserRequest.ToUser()

 err := us.userRepo.Create(user)
 if err != nil {
  return nil, &models.ErrorResponse{
   Code:    http.StatusInternalServerError,
   Message: "Failed to create user",
  }
 }

 return userResponse.FromUser(user), nil
}

func (us *User) UpdateUser(userID int,updateUserRequest *dtos.UpdateUserRequest) *models.ErrorResponse{
	existingUser,err:=us.userRepo.FindById(userID)
	if err!=nil{
		if errors.Is(err,sql.ErrNoRows){
			return &models.ErrorResponse{
				Code: http.StatusNotFound,
				Message: "User not found",
			}
		}
		return &models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal server error",
		}	
	}

	if updateUserRequest.Email!=existingUser.Email{
		errEmail:=us.checkIfEmailExists(updateUserRequest.Email)
		if errEmail!=nil{
			return errEmail
		}
	}

	existingUser=updateUserRequest.ToUser()
	existingUser.ID=userID

	err=us.userRepo.Update(existingUser)
	if err!=nil{
		return &models.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Failed to update user",
		}	
	}
	return nil
}

func (us *User) checkIfEmailExists(email string) *models.ErrorResponse {
 userWithEmail, err := us.userRepo.FindByEmail(email)
 if err != nil && !errors.Is(err, sql.ErrNoRows) {
  return &models.ErrorResponse{
   Code:    http.StatusInternalServerError,
   Message: "Internal Server Error",
  }
 }
 if userWithEmail != nil {
  return &models.ErrorResponse{
   Code:    http.StatusBadRequest,
   Message: "Email already in use",
  }
 }
 return nil
}