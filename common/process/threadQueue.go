package process

import (
	"sync"
)

type ThreadQueue struct {
	queue map[string]uint32
	mutex *sync.Mutex
}

//添加新的map项
func (q *ThreadQueue) Add(applicationID string, processID uint32) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue[applicationID] = processID
}

//获取进程id
func (q ThreadQueue) GetProcessID(applicationId string) uint32 {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if v, ok := q.queue[applicationId]; !ok {
		return 0
	} else {
		return v
	}
}
func (q ThreadQueue) GetApplicationIDByProcessID(proccessId uint32) string {
	for k, v := range q.queue {
		if v == proccessId {
			return k
		}
	}
	return ""
}

//获取所有的进程
func (q ThreadQueue) GetAllProcess() map[string]uint32 {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.queue
}

//清空所有的元素
func (q ThreadQueue) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for k := range q.queue {
		delete(q.queue, k)
	}
}

//删除指定的线程，通过应用id
func (q *ThreadQueue) Delete(applicationId string) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if _, ok := q.queue[applicationId]; !ok {
	} else {
		delete(q.queue, applicationId)
	}
	return true
}

func NewThreadQueue() *ThreadQueue {
	queue := ThreadQueue{
		queue: make(map[string]uint32),
		mutex: &sync.Mutex{},
	}
	return &queue
}
