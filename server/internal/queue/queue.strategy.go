package queue

type IStrategy interface {
	Push()
	Listen()
}
