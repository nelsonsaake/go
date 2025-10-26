package repo

import "gorm.io/gorm"

// CreateTx creates a user within the given transaction
func CreateTx[T any](tx *gorm.DB, dto any) (*T, error) {

	tx = tx.Model(new(T))

	res, err := cast[T](dto)
	if err != nil {
		return nil, err
	}

	err = tx.Create(res).Error

	return res, err
}
