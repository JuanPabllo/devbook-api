package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type user struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

func (repositories user) Create(user models.User) (uint64, error) {
	statement, erro := repositories.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastInsertID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertID), nil
}

func (repositories user) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, erro := repositories.db.Query(
		"select id, name, nick, email, createdAt from users where name like ? or nick like ?",
		nameOrNick,
		nameOrNick,
	)
	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repositories user) SearchById(userID uint64) (models.User, error) {
	lines, erro := repositories.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		userID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repositories user) Update(userID uint64, user models.User) error {
	statement, erro := repositories.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		fmt.Println("erro 3")
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, userID); erro != nil {
		return erro
	}

	return nil
}

func (repositories user) Delete(userID uint64) error {
	statement, erro := repositories.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID); erro != nil {
		return erro
	}

	return nil
}

func (repositorie user) SearchByEmail(email string) (models.User, error) {
	line, erro := repositorie.db.Query(
		"select id, password from users where email = ?",
		email,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repositorie user) Follow(userID, followerID uint64) error {
	statement, erro := repositorie.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

func (repositorie user) UnFollow(userID, followerID uint64) error {
	statement, erro := repositorie.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

func (repositorie user) SearchFollowers(userID uint64) ([]models.User, error) {
	lines, erro := repositorie.db.Query(`
	  select u.id, u.name, u.nick, u.email, u.createdAt
		from users u inner join followers f on u.id = f.follower_id where f.user_id = ?
	`, userID)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repositorie user) SearchFollowing(userID uint64) ([]models.User, error) {
	lines, erro := repositorie.db.Query(`
	select u.id, u.name, u.nick, u.email, u.createdAt
	from users u inner join followers f on u.id = f.user_id where f.follower_id = ? 
	`, userID)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repositorie user) SearchPassword(userID uint64) (string, error) {
	line, erro := repositorie.db.Query(
		"select password from users where id = ?",
		userID,
	)
	if erro != nil {
		return "", erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.Password); erro != nil {
			return "", erro
		}
	}

	return user.Password, nil
}

func (repositorie user) UpdatePassword(userID uint64, password string) error {
	statement, erro := repositorie.db.Prepare("update users set password = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(password, userID); erro != nil {
		return erro
	}

	return nil
}
