package models

// SQLDataBase struct
type SQLDataBase struct {
	Server          string   `toml:"Server"`
	Database        string   `toml:"Database"`
	ApplicationName string   `toml:"ApplicationName"`
	User            string
	Password        string
}

type Material struct {
	Title       string
	Description string
	Ref         string
	DateCreated string
}
