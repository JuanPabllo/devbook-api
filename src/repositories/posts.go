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

func (repositorie Posts) SearchPosts(postID uint64) ([]models.Posts, error) {
	lines, erro := repositorie.db.Query(`
	  select distinct p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		inner join followers f on p.author_id = f.user_id
		where u.id = ? or f.follower_id = ?
		order by 1 desc
	`, postID, postID)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var posts []models.Posts

	for lines.Next() {
		var post models.Posts

		if erro = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repositorie Posts) Update(postID uint64, post models.Posts) error {
	statement, erro := repositorie.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(post.Title, post.Content, postID); erro != nil {
		return erro
	}

	return nil
}

func (repositorie Posts) Delete(postID uint64) error {
	statement, erro := repositorie.db.Prepare("delete from posts where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}

func (repositorie Posts) GetPostsByUser(userID uint64) ([]models.Posts, error) {
	lines, erro := repositorie.db.Query(`
		select p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		where p.author_id = ?
	`, userID)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var posts []models.Posts

	for lines.Next() {
		var post models.Posts

		if erro = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repositorie Posts) Like(userID uint64) error {
	statement, erro := repositorie.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID); erro != nil {
		return erro
	}

	return nil
}

func (repositorie Posts) Dislike(userID uint64) error {
	statement, erro := repositorie.db.Prepare(`
	update posts set likes = 
	CASE 
		WHEN likes > 0 THEN likes - 1
		ELSE 0 
	END
	where id = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID); erro != nil {
		return erro
	}

	return nil
}
