package repository

import (
	"asia-quest/config"
	"asia-quest/entity"
	"asia-quest/entity/request"
	"asia-quest/entity/response"
	"database/sql"
	"fmt"
)

type authRepository struct {
}

func NewAuthRepository() entity.AuthRepository {
	return &authRepository{}
}
func (r *authRepository) Login(request *request.LoginRequest) (*response.LoginResponse, error) {
	var data response.LoginResponse
	var username sql.NullString
	var password sql.NullString

	db, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		return nil, err
	}

	var query string = fmt.Sprintf(`SELECT username,password FROM 	authentication WHERE username = "%s"`, request.Username)
	db.QueryRow(query).Scan(
		&username,
		&password,
	)
	fmt.Println(query)
	data.Username = username.String
	data.Password = password.String
	defer db.Close()
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed Select SQL for books : %v", err)
	}

	return &data, nil
}
