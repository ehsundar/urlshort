
sqlc: sqlc.yaml storage/pg/sql/schema.sql storage/pg/sql/query.sql
	sqlc generate
