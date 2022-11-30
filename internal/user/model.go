package user

type User struct {
	ID       uint
	Email    string
	Password string
	Name     string
	Salt     string
	Hash     string
}
