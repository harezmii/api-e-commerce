package seed

import (
	"api/ent"
	"context"
)

type Seeder struct {
	Client   *ent.Client
	Context  context.Context
	SeedInt  int
	Entities []string
}
type ISeeder interface {
	Seed()
}
