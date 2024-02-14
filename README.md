## Sharecycle-application-backend

## Summary of service

This is the backend of sharecycle application.
Connecting people who give things that they don't need and people who need them.

### Versions:

- `go`: 1.20
- `mySQL`: 8.0.35

## API Documentation Reference

## Running the service locally

- go mod tidy
- create a mySQL schema, name it sharecycle-local ( for local )
- set env in local.env (db's information, port)
- go run cmd/app/main.go

## Database Migration

- Migration live in `/migrations`
- They are [goose](https://github.com/pressly/goose) based
- Migration auto migrate up/down base on value on version.txt

## Database schema

## Manual Testing guide

- Using postman collection, set variable value inside collection configuration.

## Techical notes

- The project follow clean architect with custom modulized for each feature for easier maintain and scale

## TODO

## NOTES
