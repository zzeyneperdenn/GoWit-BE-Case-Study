echo "Running migration on: ${DATABASE_HOST}:${DATABASE_PORT}"
migrate -path ./db/migrations \
  -database "postgresql://postgres:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable"\
  down 1
echo "Migration Finished"
