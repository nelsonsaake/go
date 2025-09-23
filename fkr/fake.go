package fkr

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/nelsonsaake/go/rand"
	"github.com/nelsonsaake/go/strs"
)

func init() {
	gofakeit.Seed(0)
}

// todo: break the dependency on strs package
func UUID() string {
	return strs.UUID()
}

func Word() string {
	return gofakeit.Word()
}

func Name() string {
	return gofakeit.Name()
}

func Bool() bool {
	return rand.Bool()
}

func Sentence() string {
	return gofakeit.Sentence(10)
}

func IPv4() string {
	return gofakeit.IPv4Address()
}

func PhoneNumber() string {

	// Ghanaian mobile prefixes
	prefixes := []string{"024", "054", "055", "056", "057", "059", "020", "050"}
	prefix := rand.Element(prefixes)

	// Generate 7 random digits and pad with zeros
	randNum := rand.Int(10000000)
	num := fmt.Sprintf("%07d", randNum)

	return "+233" + prefix[1:] + num
}

func Email() string {
	return gofakeit.Email()
}

func Address() string {
	return gofakeit.Address().Address
}

func ImageURL() string {
	return gofakeit.ImageURL(640, 480)
}

func FilePath() string {
	exts := []string{".txt", ".pdf", ".docx", ".xlsx", ".jpg", ".png", ".csv"}
	folders := []string{"documents", "downloads", "images", "music", "videos", "work"}
	folder := rand.Element(folders)
	filename := gofakeit.Word() + gofakeit.Word() + rand.Element(exts)
	return fmt.Sprintf("/home/user/%s/%s", folder, filename)
}

func ImageFilePath() string {
	exts := []string{".jpg", ".png", ".gif", ".bmp"}
	filename := gofakeit.Word() + rand.Element(exts)
	return fmt.Sprintf("/home/user/images/%s", filename)
}
