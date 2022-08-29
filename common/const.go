package common

const (
	DbTypeFlower     = 1
	DbTypeCategory   = 2
	DbTypeUser       = 3
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
