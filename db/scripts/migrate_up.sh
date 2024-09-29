echo "Running migration on: ${DATABASE_HOST}:${DATABASE_PORT}"
migrate -path ./db/migrations \
  -database "postgresql://postgres:secret@gowit-be-case-study-db-1:5432/postgres?sslmode=disable"\
  up
echo "Migration Finished"
