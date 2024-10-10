# GoNav Backend

Version: 1.0.0

This is the backend part of the GoNav project, built with Go and the Gin framework.

## Requirements

- Go (version 1.16 or higher recommended)
- SQLite3

## Installation and Running

1. After cloning the project, enter the backend directory:

   ```bash
   cd backend
   ```
2. Install dependencies:

   ```bash
   go mod tidy
   ```
3. Configure the application:

   - Open `config.yaml` and adjust the settings as needed.
4. Run the server:

   ```bash
   go run main.go
   ```
5. The server will run at `http://localhost:8080` by default

## Configuration

The `config.yaml` file contains the following settings:

- `server`:
  - `port`: The port on which the server runs (default: 8080)
  - `mode`: Gin mode, can be "debug" or "release"
- `database`:
  - `type`: Database type (currently only supports "sqlite")
  - `name`: Database file name
  - `path`: Path to the database file
- `cors`:
  - `allowOrigins`: List of allowed origins for CORS
- `jwt`:
  - `secret`: Secret key for JWT token generation
  - `expirationHours`: JWT token expiration time in hours

## Main Features

- User authentication
- Bookmark CRUD operations
- Category management
- Settings management
- Public bookmark sharing
