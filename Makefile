migrate_database:
	migrate -database postgres://postgres:@localhost:5432/affiliate?sslmode=disable -path db/migrations up