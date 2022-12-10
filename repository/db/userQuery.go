package db

import (
	"fmt"
	"time"

	"github.com/guntoroyk/go-user-api/entity"
)

var (
	selectUsersQuery = `SELECT * FROM users`
	selectUserQuery  = `SELECT * FROM users WHERE %s ?`
	updateUserQuery  = `UPDATE users SET %s WHERE id = %d`
	insertUserQuery  = `INSERT INTO users (name, username, password, phone, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	deleteUserQuery  = `DELETE FROM users WHERE id = ?`
)

func buildSelectUserQuery(c *entity.UserFilter) (string, []interface{}) {
	var filters []string
	var filter string
	var args []interface{}

	if c.ID != 0 {
		filters = append(filters, "id=")
		args = append(args, c.ID)
	}
	if c.Username != "" {
		filters = append(filters, "username=")
		args = append(args, c.Username)
	}

	for i, f := range filters {
		if i == 0 {
			filter += f
		} else {
			filter += " AND " + f
		}
	}

	query := fmt.Sprintf(selectUserQuery, filter)
	return query, args
}

func buildUpdateUserQuery(c *entity.User) (string, []interface{}) {
	var query string
	var args []interface{}

	if c.Name != "" {
		query += "name=?,"
		args = append(args, c.Name)
	}
	if c.Username != "" {
		query += "username=?,"
		args = append(args, c.Username)
	}
	if c.Password != "" {
		query += "password=?,"
		args = append(args, c.Password)
	}
	if c.Phone != "" {
		query += "phone=?,"
		args = append(args, c.Phone)
	}
	if c.Role != "" {
		query += "role=?,"
		args = append(args, c.Role)
	}
	query += "updated_at=?,"
	args = append(args, time.Now().Format("2006-01-02 15:04:05"))

	query = fmt.Sprintf(updateUserQuery, query[:len(query)-1], c.ID)
	return query, args
}
