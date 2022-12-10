package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/guntoroyk/go-user-api/entity"
	"github.com/guntoroyk/go-user-api/repository"
)

type userRepo struct {
	db *sql.DB
}

// NewUserRepo will create new an UserRepo object representation of UserRepoItf interface
func NewUserRepo(db *sql.DB) repository.UserRepoItf {
	return &userRepo{
		db: db,
	}
}

// GetUsers will get all users
func (c *userRepo) GetUsers() ([]*entity.User, error) {
	rows, err := c.db.Query(selectUsersQuery)
	if err != nil {
		log.Println("func GetUsers error, ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.ID, &user.Name,
			&user.Username, &user.Password,
			&user.Phone, &user.Role,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			log.Println("func GetUsers error scanning, ", err.Error())
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// GetUser will get user by id
func (c *userRepo) GetUser(filter *entity.UserFilter) (*entity.User, error) {
	var user entity.User
	selectQuery, fields := buildSelectUserQuery(filter)
	err := c.db.QueryRow(selectQuery, fields...).Scan(
		&user.ID, &user.Name,
		&user.Username, &user.Password,
		&user.Phone, &user.Role,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		log.Println("func GetUser error, ", err.Error())
		return nil, err
	}

	return &user, nil
}

// CreateUser will create a user
func (c *userRepo) CreateUser(user *entity.User) (*entity.User, error) {
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	res, err := c.db.Exec(insertUserQuery,
		user.Name, user.Username,
		user.Password, user.Phone, user.Role,
		user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		log.Println("func CreateUser error, ", err.Error())
		return nil, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Println("func CreateUser error, ", err.Error())
		return nil, err
	}
	user.ID = int(lastId)
	return user, nil
}

// UpdateUser will update a user
func (c *userRepo) UpdateUser(user *entity.User) (*entity.User, error) {
	updateQuery, fields := buildUpdateUserQuery(user)
	resp, err := c.db.Exec(updateQuery, fields...)
	if err != nil {
		log.Println("func UpdateUser error exec, ", err.Error())
		return nil, err
	}
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		log.Println("func UpdateUser error getting rows affected, ", err.Error())
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return user, nil
}

// DeleteUser will delete a user
func (c *userRepo) DeleteUser(id int) error {
	resp, err := c.db.Exec(deleteUserQuery, id)
	if err != nil {
		log.Println("func DeleteUser error exec, ", err.Error())
		return err
	}
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		log.Println("func DeleteUser error getting rows affected, ", err.Error())
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
