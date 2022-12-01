#!/bin/sh

echo "Waiting for postgres database to start..."
./wait-for postgres:5432

echo "Migrating the database..."
migrate -path migration/ -database "postgres://farzaneh:3971231050@postgres/ukit?sslmode=disable" -verbose up

echo "Starting the server..."
/app/main