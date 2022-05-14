package server

import (
	"container/list"
	"time"

	"github.com/gerardjlr7/chat-server/protos"
)

type QueueMessages struct {
	queue *list.List
}

type Message struct {
	Timestamp time.Time
	Content   *protos.MessageRequest
}

func NewQueue() *QueueMessages {
	return &QueueMessages{
		queue: list.New(),
	}
}

func (q *QueueMessages) Enqueue(content *protos.MessageRequest) {
	timestamp := time.Now()

	q.queue.PushBack(Message{
		Timestamp: timestamp,
		Content:   content,
	})
}

func (q *QueueMessages) Dequeue() *protos.MessageReply {
	//Dequeue every element with timestamp before 24 hours
	if q.queue.Len() > 0 {
		front := q.queue.Front()
		reply := parseRequestToReply(front)
		q.queue.Remove(front)
		return reply
	}
	return nil
}

func parseRequestToReply(front *list.Element) *protos.MessageReply {
	msg := front.Value.(Message)
	request := msg.Content

	return &protos.MessageReply{
		Name:      request.Name,
		Message:   request.Message,
		Timestamp: msg.Timestamp.Format("2006-01-02 15:04"),
	}
}
