package v1

const (
	ApiVersion     string = "v1"
	SuccessMessage string = "success"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
