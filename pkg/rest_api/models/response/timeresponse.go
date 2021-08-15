package response

type TimeResponse struct {
	UTC  string `json:"utc"`
	Unix int64  `json:"unix"`
}

type TimeResponseError struct {
	Error string `json:"error"`
}
