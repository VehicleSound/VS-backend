package interfaces

type Metrics interface {
	AddHttpRequest(status int)
	AddNewUser()
}
