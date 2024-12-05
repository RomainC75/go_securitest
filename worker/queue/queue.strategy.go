package queue

type IStrategy interface {
	Push(key string, data []byte) error
	Listen()
}
