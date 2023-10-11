package graph

import (
	"database/sql"
	"github.com/kkoji/gqlgen-todos/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *sql.DB
}

func (r *Resolver) createUser(input model.NewUser) (*model.User, error) {
	stmt, err := r.DB.Prepare("INSERT INTO users (name) VALUES ($1) RETURNING id, name")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	user :=&model.User{}
	err = stmt.QueryRow(input.Name).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Resolver) getUsers() ([]*model.User, error) {
	rows, err := r.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}

	var users []*model.User
	u := &model.User{}
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}