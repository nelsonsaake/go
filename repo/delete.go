package repo

import (
	"fmt"
	"strings"

	"github.com/nelsonsaake/go/strs"
)

// interpreteDeleteError interprets database delete errors and returns a more
// user-friendly error message if there are foreign key constraints.
func (r *Repo[T]) interpreteDeleteError(err error) error {

	resource := getTypeName(new(T))
	resource = strs.ToWords(resource)
	resource = strings.ToLower(resource)

	// Check for MySQL error 1451 using substring match
	if strings.Contains(err.Error(), "Error 1451") {
		return fmt.Errorf("cannot delete %q, other models depend on it", resource)
	}

	return nil
}

// Delete removes a record of type T with the given id from the database.
// Returns an error if deletion fails or is blocked by constraints.
func (r *Repo[T]) Delete(id string) error {

	var model = new(T)
	err := r.DB.Where("id = ?", id).Delete(model).Error
	if err != nil {
		return r.interpreteDeleteError(err)
	}

	return err
}

// Delete is a convenience function that removes a record of type T
// with the given id using the default repository instance.
func Delete[T any](id string) error {
	return Do[T]().Delete(id)
}
