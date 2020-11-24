package model

import (
	"time"
)

// UserModel database interface
type UserModel interface {
	ListUsers() ([]User, error)
}

// UserActivity represents a single user initiated activity.
type User struct {
    dbDates
    Name string `db:"name"`
    Login string `db:"login"`
    Password string `db:"login"`
    Role int `db: "role"`
}

// ListUsers returns list of all users.
// TODO: pagination
func (m *model) ListUsers() ([]User, error) {
	var users []User
	err := db.Select(&activity, `
		-- START ListUsers
		SELECT
			u.name as name,
			u.login as login,
            u.password as password
            r.role as role
        FROM
			user as u 
        LEFT JOIN user_role as u_r ON u.id=u_r.user_id
		ORDER BY user.id DESC
		-- END ListUsers
	`)
	if err != nil {
		return nil, err
	}
	return users, nil
}
