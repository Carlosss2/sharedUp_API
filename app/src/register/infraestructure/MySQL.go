package infraestructure

import (
	"database/sql"
	"sharedup/app/src/register/domain/entities"
)

type MySQL struct {
	db *sql.DB
	
}

func NewMySQL(db *sql.DB)*MySQL{
	return &MySQL{db: db}
}

func (mysql *MySQL) Save(user entities.User) error {
	query := "INSERT INTO users (name, email, password, career) VALUES (?, ?, ?, ?)"
	_, err := mysql.db.Exec(query, user.Name, user.Email, user.Password,user.Career)
	if err != nil {
		return err
	}
	return nil

}

