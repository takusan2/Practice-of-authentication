package services

type UserService interface {
	SignUpWithEmailAndPassword(email, password string) (int64, error)
	LoginWithEmailAndPassword(email, password string) (int64, error)
}
