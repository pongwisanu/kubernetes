package repositories

import (
	"github.com/jmoiron/sqlx"
)

type userRepositoryDb struct {
	db *sqlx.DB
}

func NewUserRepositoryDb(db *sqlx.DB) UserRepository {
	return userRepositoryDb{db: db}
}

func (r userRepositoryDb) GetUsers() ([]User, error) {
	users := []User{}
	query := "SELECT * FROM user_tb"
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepositoryDb) GetUser(id int) (*User, error) {
	user := User{}
	query := "SELECT * FROM user_tb where id = ?"
	err := r.db.Select(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepositoryDb) AddUser(user NewUser) (*User, error) {

	query := "INSERT INTO user_tb (name, detail, role) VALUES ($1,$2,$3)"

	_, err := r.db.Exec(query, user.Name, user.Detail, user.Role)

	if err != nil {
		return nil, err
	}

	newUser := User{
		Name:   user.Name,
		Detail: user.Detail,
		Role:   user.Role,
	}

	return &newUser, nil
}

// func (r userRepositoryDb) EditUser(user *User) error {
// 	query := "UPDATE user_tb SET name=?, detail=?, role=?"
// 	_, err := r.db.Exec(query, user.Name, user.Detail, user.Role)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r userRepositoryDb) DeteteUser(id int) error {
	query := "DELETE FROM user_tb WHERE id=?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
