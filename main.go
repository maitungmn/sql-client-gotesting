package main

import (
	"errors"
	"fmt"

	"github.com/maitungmn/sql-client-gotesting/sqlclient"
)

const (
	queryGetUser = "SELECT id, email FROM users WHERE  id=?;"
)

var (
	dbClient sqlclient.SqlClient
)

type User struct {
	Id    int64
	Email string
}

func init() {
	sqlclient.StartMockServer()
	var err error
	dbClient, err = sqlclient.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "root", "127.0.0.1:3306", "users_db"))

	if err != nil {
		panic(err)
	}
}

func main() {
	user, err := GetUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Id)
	fmt.Println(user.Email)
}

func GetUser(id int64) (*User, error) {
	sqlclient.AddMock(sqlclient.Mock{
		Query:   "SELECT id, email FROM users WHERE  id=?;",
		Args:    []interface{}{1},
		Columns: []string{"id", "email"},
		Rows:    [][]interface{}{
			{1, "email1"},
			{2, "email2"},
		},
	})
	rows, err := dbClient.Query(queryGetUser, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	for rows.HasNext() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user not found")
}
