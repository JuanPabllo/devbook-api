package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repositorie Posts) Create(posts models.Posts) (uint64, error) {
	statement, erro := repositorie.db.Prepare(
		"insert into posts (title, content, author_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(posts.Title, posts.Content, posts.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}
func (repositorie Posts) GetPostByID(postID uint64) (models.Posts, error) {
	line, erro := repositorie.db.Query(`
		select p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		where p.id = ?
	`, postID)
	if erro != nil {
		return models.Posts{}, erro
	}
	defer line.Close()

	var post models.Posts
	if line.Next() {
		if erro = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); erro != nil {
			return models.Posts{}, erro
		}
	}

	return post, nil
}
