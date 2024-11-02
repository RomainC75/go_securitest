package queue

type SQueue struct {
	strategy IStrategy
}

var queueInstance *SQueue

func SetQueue(strategy IStrategy) {
	queueInstance = &SQueue{
		strategy: strategy,
	}
}

func GetQueue() *SQueue {
	return queueInstance
}
