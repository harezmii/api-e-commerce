package repository

//
//import "gorm.io/gorm"
//
//type IBaseRepository[T any] interface {
//	Store(t T) (T, error)
//	Index() ([]T, error)
//	Show(id int) error
//	Update(t T) error
//}
//
//type BaseRepository[T any] struct {
//	db *gorm.DB
//}
//
//func (r BaseRepository[T]) Store(t T) (T, error) {
//	err := r.db.Create(&t).Error
//	return t, err
//}
//
//func (r BaseRepository[T]) Index() ([]T, error) {
//	var t []T
//	err := r.db.Find(&t).Error
//	return t, err
//}
//
//func (r BaseRepository[T]) Show(id int) (T, error) {
//	var t T
//	err := r.db.Where("id = ?", id).First(&t).Error
//	return t, err
//}
//
//func (r BaseRepository[T]) Update(t T) error {
//	err := r.db.Save(&t).Error
//	return err
//}
//
//func (r BaseRepository[T]) Delete(t T) error {
//	err := r.db.Delete(&t).Error
//	return err
//}
