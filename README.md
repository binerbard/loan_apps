# Loan Apps

This is a simple loan apps built with Golang.

## Features

- User can register and login
- User can apply for loan
- User can see the list of loan
- User can see the detail of loan
- User can make payment for loan
- User can see the list of payment

## Technologies

- Golang
- Gorm
- Fiber
- Minio
- MySQL

## How to run

1. Clone this repository
2. Create a new file named `.env` and add the following variables:
   - DB_USER
   - DB_PASSWORD
   - DB_HOST
   - DB_PORT
   - DB_NAME
   - STORAGE_ACCESS_KEY
   - STORAGE_SECRET_KEY
   - STORAGE_BUCKET
3. Run the command `go run cmd/api/main.go`
4. Open the browser and go to `localhost:3000`
