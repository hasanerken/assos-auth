package storages

import (
	"assos/ent"
	"context"
	"fmt"
	"log"
	"os"
)


func InitializeDatabase() (*ent.Client, error) {
	host := os.Getenv("PGHOST")
    port := os.Getenv("PGPORT")
    user := os.Getenv("PGUSER")
    password := os.Getenv("PGPASSWORD")
    database := os.Getenv("PGDATABASE")

    connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, database, password)

    client, err := ent.Open("postgres", connectionString)
    if err != nil {
        return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
    }

    // Run the auto migration tool.


    return client, nil
}

func MigrateDatabase(ctx context.Context, client *ent.Client) error {

	if err := client.Schema.Create(context.Background()); err != nil {
        return fmt.Errorf("failed creating schema resources: %v", err)
    }

	log.Println("Database schema migration successful")

	return nil
}
