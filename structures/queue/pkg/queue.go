package queue

import "slices"

type Queue struct {
	items []any
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (any, bool) {
	if len(q.items) == 0 {
		return nil, false
	}
	item := q.items[0]
	q.items[0] = nil
	q.items = q.items[1:]
	return item, true
}

func (q *Queue) Peek() (any, bool) {
	if len(q.items) == 0 {
		return nil, false
	}
	return q.items[0], true
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) Items() []any {
	return slices.Clone(q.items)
}
