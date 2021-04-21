package validation

const (
	OK = iota
	NotValid
)

type Status int

type Result struct {
	Message string `json:"message"`
	Status  Status `json:"status"`
}
