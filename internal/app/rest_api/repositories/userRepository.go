package repositories

import (
	"database/sql"
	"example/rest-api/internal/app/rest_api/database"
	"example/rest-api/internal/app/rest_api/entities"
)

type User struct{
	database.BaseSQLRepository[entities.User]
}

func NewUserRepository(db *sql.DB) *User {
	return &User{
		BaseSQLRepository: database.BaseSQLRepository[entities.User]{DB: db},
	}
}

func mapUser(row *sql.Row,u *entities.User) error{
	return row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PhoneNumber,
	)
}

func mapUsers(rows *sql.Rows,u *entities.User) error{
	return rows.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PhoneNumber,
	)
}

func (repo *User) FindByEmail(email string)(*entities.User,error){
	return repo.SelectSingle(
		mapUser,
		"SELECT id, first_name, last_name, email, phone_number FROM users WHERE email=$1",
		email,
	)
}

func (repo *User) FindById(id int)(*entities.User,error){
	return repo.SelectSingle(
		mapUser,
		"SELECT id, first_name, last_name, email, phone_number FROM users WHERE id=$1",
		id,
	)
}

func (repo *User) GetAllUsers()([]*entities.User,error){
	return repo.SelectMultiple(
		mapUsers,
		"SELECT id, first_name, last_name, email, phone_number FROM users",
	)
}

func (repo *User) Create(user *entities.User) error{
	_,err:=repo.ExecuteQuery(
		"INSERT INTO users (first_name, last_name, email, phone_number) VALUES ($1, $2, $3, $4)",
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
	)
	return err	
}

func (repo *User) Update(user *entities.User) error{
	_, err:=repo.ExecuteQuery(
		"UPDATE users SET first_name=$1, last_name=$2, email=$3, phone_number=$4 WHERE id=$5",
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.ID,
	)
	return err
}

func (repo *User) Delete(id int) error{
	_,err:=repo.ExecuteQuery("DELETE FROM users WHERE id =$1",id)
	return err
}
