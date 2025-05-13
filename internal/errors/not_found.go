package errors

type NotFoundError struct {
	Slug string
}

func (e *NotFoundError) Error() string {
	return "URL not found for slug '" + e.Slug + "'"
}
