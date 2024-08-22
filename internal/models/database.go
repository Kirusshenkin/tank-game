package models

// DatabaseConfig содержит параметры для подключения к базе данных
type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}
