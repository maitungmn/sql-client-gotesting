package sqlclient

import "database/sql"

type sqlRows struct {
	rows *sql.Rows
}

type rows interface {
	HasNext() bool
	Close() error
	Scan(destinations ...interface{}) error
}

func (r *sqlRows) HasNext() bool {
	return r.rows.Next()
}
func (r *sqlRows) Close() error {
	return r.rows.Close()
}
func (r *sqlRows) Scan(destinations ...interface{}) error {
	return r.rows.Scan(destinations...)
}
