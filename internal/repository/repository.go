package repository

type DatabaseRepo interface {
	Authenticate(email, testPassword string) (int, string, error)
}
