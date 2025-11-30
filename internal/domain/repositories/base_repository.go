package repositories

type baseRepository[T any] interface {
	Get(id int) (T, error)
	Del(id int) error
	GetAll() ([]T, error)
}
