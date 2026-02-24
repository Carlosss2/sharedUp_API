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

func (mysql *MySQL) Update(p entities.Post) error {
	query := "UPDATE posts SET title = ?, text = ? WHERE idposts = ? AND iduser = ?"
	result, err := mysql.DB.Exec(query, p.Title, p.Text, p.Id, p.IdUser)
	if err != nil {
		return fmt.Errorf("[MySQL] Error al actualizar la publicacion: %w", err)
	}
	
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("publicacion no encontrada o no tienes permisos para editarla")
	}
	return nil
}

func (mysql *MySQL) Delete(id int, idUser int) error {
	query := "DELETE FROM posts WHERE idposts = ? AND iduser = ?"
	result, err := mysql.DB.Exec(query, id, idUser)
	if err != nil {
		return fmt.Errorf("[MySQL] Error al eliminar la publicacion: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("publicacion no encontrada o no tienes permisos para eliminarla")
	}
	return nil
}