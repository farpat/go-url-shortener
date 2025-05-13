package errors

type InvalidUrlError struct {
	Url string
}

func (e *InvalidUrlError) Error() string {
	return "Invalid URL: " + e.Url
}
