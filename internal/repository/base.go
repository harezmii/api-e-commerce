package repository

type IBaseRepository[T any] interface {
	Create(t T) (T, error)
	Show() (T, error)
	Index() ([]T, error)
	Update(t T) error
	Delete(t T) (T, error)
}

type BaseRepository[T any] struct {
}

func (r BaseRepository[T]) Create(t T) (T, error) {
	return t, nil
}
