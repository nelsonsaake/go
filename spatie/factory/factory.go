package factory

import (
	"github.com/bxcodec/faker/v3"
	"github.com/nelsonsaake/go/spatie/models"
)

func FakeRole() *models.Role {
	return &models.Role{
		Base: models.Base{
			ID: faker.UUIDDigit(),
		},
		Name: faker.Word(),
	}
}

func FakePermission() *models.Permission {
	return &models.Permission{
		Base: models.Base{
			ID: faker.UUIDDigit(),
		},
		Name: faker.Word(),
	}
}

// Add similar functions for other models as needed
func FakeUser() *models.User {
	return &models.User{
		Base: models.Base{
			ID: faker.UUIDDigit(),
		},
	}
}

func FakeRolePermission() *models.RolePermission {
	return &models.RolePermission{
		Base: models.Base{
			ID: faker.UUIDDigit(),
		},
		RoleID:       faker.UUIDDigit(),
		PermissionID: faker.UUIDDigit(),
	}
}

func FakeUserRole() *models.UserRole {
	return &models.UserRole{
		Base: models.Base{
			ID: faker.UUIDDigit(),
		},
		UserID: faker.UUIDDigit(),
		RoleID: faker.UUIDDigit(),
	}
}

func FakeUserPermission() *models.UserPermission {
	return &models.UserPermission{
		Base: models.Base{
			ID: faker.UUIDDigit(),
		},
		UserID:       faker.UUIDDigit(),
		PermissionID: faker.UUIDDigit(),
	}
}

func FakeUserPermissionRevoked() *models.UserPermissionRevoked {
	return &models.UserPermissionRevoked{
		Base: models.Base{
			ID: faker.UUIDDigit(),
		},
		UserID:       faker.UUIDDigit(),
		PermissionID: faker.UUIDDigit(),
	}
}
