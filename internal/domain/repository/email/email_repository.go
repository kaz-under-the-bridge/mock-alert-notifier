package email

type RepositoryInterface interface {
	Send() error
}
