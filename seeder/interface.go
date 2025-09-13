package seeder

type Seeder interface {
	Seed(args ...int) error
}
