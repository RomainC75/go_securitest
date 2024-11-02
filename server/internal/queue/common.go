package queue

type SQueue struct {
	Strategy IStrategy
}

var queueInstance *SQueue

func SetQueue(strategy IStrategy) {
	queueInstance = &SQueue{
		Strategy: strategy,
	}
}

func GetQueue() *SQueue {
	return queueInstance
}
