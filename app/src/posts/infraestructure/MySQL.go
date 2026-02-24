package infraestructure

import (
	"database/sql"
	"sharedup/app/src/posts/domain/entities"
	"fmt"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL{
	return &MySQL{DB:db}
}

func (mysql *MySQL) Save(p entities.Post) error{
	query := "INSERT INTO posts (title,text,like_count,dislike_count,iduser) VALUES(?,?,?,?,?)"
	_, err := mysql.DB.Exec(query,p.Title,p.Text,p.LikeCount,p.DisLikeCount,p.IdUser)
	 if err != nil {
            return fmt.Errorf("[MySQL] Error al guardar la publicacion: %w", err)
        }
        return nil
}