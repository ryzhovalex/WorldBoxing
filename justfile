set shell := ["nu", "-c"]
set dotenv-load
dbmate := if os_family() == "windows" { "dbmate.cmd" } else { "dbmate" }

run:
    @ go run .

lint:
    @ go fmt

test t="":
    @ if "{{t}}" == "" { go test ./... } else { go test -run {{t}} }

check: lint test

db *t:
    @ {{dbmate}} -s "Database/SqliteSchema.sql" -d "Database/Migrations" {{t}}
