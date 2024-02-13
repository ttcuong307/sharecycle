## Sharecycle-application-backend

## Summary of service

This is the backend of sharecycle application.
Allow you to post items that you don't need and the one that need will find

### Versions:

- `go`: 1.20
- `mySQL`:

## API Documentation Reference

## Running the service locally

- navigate to the backend folder
- go mod tidy
- set env in local.env
- go run cmd/main.go

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
