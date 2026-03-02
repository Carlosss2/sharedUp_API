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

func (mysql *MySQL) Save(p entities.Post) (int64, error) {

	query := "INSERT INTO posts (title,text,like_count,dislike_count,iduser) VALUES(?,?,?,?,?)"

	result, err := mysql.DB.Exec(query,p.Title,p.Text,p.LikeCount,p.DisLikeCount,p.IdUser)
	if err != nil {
		return 0, fmt.Errorf("[MySQL] Error al guardar la publicacion: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
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

func (mysql *MySQL) GetAll() ([]entities.PostResponse, error) {

	query := `
	SELECT 
		p.idposts,
		p.title,
		p.text,
		p.like_count,
		p.dislike_count,
		p.created_at,
		u.iduser,
		u.name,
		u.career
	FROM posts p
	INNER JOIN users u ON p.iduser = u.iduser
	ORDER BY p.idposts DESC
	`

	rows, err := mysql.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entities.PostResponse

	for rows.Next() {
		var post entities.PostResponse
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Text,
			&post.LikeCount,
			&post.DisLikeCount,
			&post.Date,
			&post.IdUser,
			&post.UserName,
			&post.UserCareer,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (mysql *MySQL) GetByUser(idUser int) ([]entities.PostResponse, error) {

	query := "SELECT idposts, title, text, like_count, dislike_count, created_at FROM posts WHERE iduser = ? ORDER BY idposts DESC"

	rows, err := mysql.DB.Query(query, idUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entities.PostResponse

	for rows.Next() {
		var post entities.PostResponse
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Text,
			&post.LikeCount,
			&post.DisLikeCount,
			&post.Date,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (mysql *MySQL) IncrementLike(postID int) error {

	query := "UPDATE posts SET like_count = like_count + 1 WHERE idposts = ?"

	result, err := mysql.DB.Exec(query, postID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("publicacion no encontrada")
	}

	return nil
}

func (mysql *MySQL) IncrementDislike(postID int) error {

	query := "UPDATE posts SET dislike_count = dislike_count + 1 WHERE idposts = ?"

	result, err := mysql.DB.Exec(query, postID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("publicacion no encontrada")
	}

	return nil
}

func (mysql *MySQL) GetByID(id int64) (entities.PostResponse, error) {

	query := `
	SELECT 
		p.idposts,
		p.title,
		p.text,
		p.like_count,
		p.dislike_count,
		p.created_at,
		u.iduser,
		u.name,
		u.career
	FROM posts p
	INNER JOIN users u ON p.iduser = u.iduser
	WHERE p.idposts = ?
	`

	var post entities.PostResponse

	err := mysql.DB.QueryRow(query, id).Scan(
		&post.Id,
		&post.Title,
		&post.Text,
		&post.LikeCount,
		&post.DisLikeCount,
		&post.Date,
		&post.IdUser,
		&post.UserName,
		&post.UserCareer,
	)

	return post, err
}