package repository

import (
	"time"

	"github.com/nihadtz/simple_user/models"
	"github.com/nihadtz/simple_user/services"
)

type UserRepository struct{}

type IUserRepository interface {
	FindUserByID(id int64) (*models.User, error)
	CreateUser(user models.User) (*int64, error)
	UpdateUser(user models.User) (*models.User, error)
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) FindUserByID(id int64) (*models.User, error) {
	db := services.Access.GetDB()
	var user models.User

	query := `SELECT 
				id, name, updated, year_of_birth
			FROM
			 	Users
			WHERE
				id = ?`

	err := db.Get(&user, query, id)

	return &user, err
}

func (u *UserRepository) CreateUser(user models.User) (*int64, error) {
	db := services.Access.GetDB()

	query := `INSERT INTO Users
						(name, updated, year_of_birth) 
					VALUES
						(:name, :updated, :year_of_birth)`

	res, err := db.NamedExec(query, user)

	if err != nil {
		return nil, err
	}

	userID, err := res.LastInsertId()

	return &userID, err
}

func (u *UserRepository) UpdateUser(user models.User) (*models.User, error) {
	db := services.Access.GetDB()

	user.Updated = time.Now().Unix()

	query := `UPDATE Users SET
				name=?, year_of_birth=?, updated=?
			WHERE 
				id=?`

	_, err := db.Exec(query, user.Name, user.YearOfBirth, user.Updated, user.ID)

	if err != nil {
		return nil, err
	}

	return u.FindUserByID(user.ID)
}
