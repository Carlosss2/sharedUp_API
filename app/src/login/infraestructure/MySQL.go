package infraestructure

import (
	"errors"
	"database/sql"
	"sharedup/app/src/login/domain/entities"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}


func (mysql *MySQL) LogIn(email string) (entities.UserWithPassword, error) {
	query := `
		SELECT iduser, name, email, password,career
		FROM users
		WHERE email = ?
	`

	var user entities.UserWithPassword

	err := mysql.db.QueryRow(query, email).
		Scan(&user.Id, &user.Name, &user.Email, &user.PasswordHash,&user.Career)

	if err != nil {
		return entities.UserWithPassword{}, errors.New("user not found")
	}

	return user, nil
}