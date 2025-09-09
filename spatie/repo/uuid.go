package repo

import (
	"github.com/google/uuid"
)

func UUID() string {
	u7, err := uuid.NewV7()
	if err == nil {
		return u7.String()
	}
	return uuid.New().String()
}
