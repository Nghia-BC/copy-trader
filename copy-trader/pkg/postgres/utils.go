package postgres

const (
	DATABASE_HOST_WRITE_ENV = "DATABASE_HOST_WRITE"
	DATBASE_HOST_ENV        = "DATABASE_HOST"
	DATABASE_PORT_ENV       = "DATABASE_PORT"
	DATABASE_USERNAME_ENV   = "DATABASE_USERNAME"
	DATBASE_PASSWORD_ENV    = "DATABASE_PASSWORD"
	NODE_ENV                = "NODE_ENV"
)

// Structure to define Postgres client

type pgConfig struct {
	host     string
	port     uint16
	user     string
	database string
	password string
	sslMode  string
}
