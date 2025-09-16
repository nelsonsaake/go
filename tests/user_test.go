package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/nelsonsaake/go/dsns"
	"github.com/nelsonsaake/go/spatie/models"
	"github.com/nelsonsaake/go/spatie/repo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {

	dsn := dsns.FromConfig(dsns.Config{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "admin",
		DBName:   "ag_spatie",
	})

	gorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	spatie := repo.New(gorm, context.Background())

	user := models.User{Base: models.Base{ID: "2fec39f7-4a62-4f96-b383-cacf3d7db13f"}}

	scope, err := spatie.MappedScope(user)
	if err != nil {
		t.Fatalf("error getting the scope: %v", err)
	}

	fmt.Print(scope)
}
