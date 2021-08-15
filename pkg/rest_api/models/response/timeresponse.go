package response

type TimeResponse struct {
	Unix int64  `json:"unix"`
	UTC  string `json:"utc"`
}

type TimeResponseError struct {
	Error string `json:"error"`
}
