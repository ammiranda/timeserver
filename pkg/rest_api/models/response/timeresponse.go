package response

type TimeResponse struct {
	UTC string  `json:"utc"`
	Unix int	`json:"unix"`
}