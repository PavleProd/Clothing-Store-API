package db

type PreparedQuery struct {
	query  string
	values []any
}

func NewPreparedQuery(query string, values []any) *PreparedQuery {
	return &PreparedQuery{
		query:  query,
		values: values,
	}
}
