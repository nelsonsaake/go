package factory

import (
	"github.com/bxcodec/faker/v3"
	"github.com/nelsonsaake/go/spatie/models"
)

func FakeRole() *models.Role {
	return &models.Role{
		ID:   faker.UUIDDigit(),
		Name: faker.Word(),
	}
}

func FakePermission() *models.Permission {
	return &models.Permission{
		ID:   faker.UUIDDigit(),
		Name: faker.Word(),
	}
}

// Add similar functions for other models as needed
