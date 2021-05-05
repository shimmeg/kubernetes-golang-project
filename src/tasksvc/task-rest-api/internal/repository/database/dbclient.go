package database

type DB interface {
	Open()

	Close()
}
