run:
    @ zig build run

test:
    @ zig build test

Migration t: (MigrationDev t)

DevMigration *t:
    @ dbmate -s "migrations/psql_schema.sql" -d "migrations/psql" -u "postgres://cpas:100400500@localhost:9005/cpasb_dev?sslmode=disable" {{t}}

ProdMigration *t:
    @ dbmate -s "migrations/psql_schema.sql" -d "migrations/psql" -u "postgres://cpas:100400500@psql:5432/cpasb_prod?sslmode=disable" {{t}}
