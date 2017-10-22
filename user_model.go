package main

import (
    "database/sql"
    //"errors"
)

type user struct {
    ID    	 int     `json:"id"`
    Name  	 string  `json:"name"`
    Gravatar string  `json:"gravatar"`
}

func (u *user) getUser(db *sql.DB) error {
  return db.QueryRow("SELECT name, gravatar FROM users WHERE id=$1", 
  	u.ID).Scan(&u.Name, &u.Gravatar)
}

func getUsers(db *sql.DB) ([]user, error) {

  rows, err := db.Query("SELECT * FROM users")

  if err != nil {
  	return nil, err
  }

  defer rows.Close()

  users := []user{}

  for rows.Next() {
  	var u user
  	if err := rows.Scan(&u.ID, &u.Name, &u.Gravatar); err != nil {
  		return nil, err
  	}
  	users = append(users, u)
  }

  return users, nil

}