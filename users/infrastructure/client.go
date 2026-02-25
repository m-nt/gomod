package infrastructure

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/m-nt/gomod/users/infrastructure/ent"
)

func NewEntClient(dsn string) *ent.Client {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run auto-migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
