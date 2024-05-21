package repositories

type User struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Detail string `db:"detail"`
	Role   string `db:"role"`
}

type NewUser struct {
	Name   string `db:"name"`
	Detail string `db:"detail"`
	Role   string `db:"role"`
}

type UserRepository interface {
	GetUsers() ([]User, error)
	GetUser(int) (*User, error)
	AddUser(NewUser) (*User, error)
	// EditUser(User) (*User, error)
	DeteteUser(int) error
}
