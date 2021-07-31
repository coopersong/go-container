package queue

type Queue struct {
    slice []interface{}
}

func (q *Queue) Front() interface{} {
    if len(q.slice) > 0 {
        return q.slice[0]
    }

    return nil
}

func (q *Queue) Pop() interface{} {
    if len(q.slice) > 0 {
        front := q.slice[0]
        q.slice = q.slice[1:]
        return front
    }

    return nil
}

func (q *Queue) Push(val interface{}) {
    q.slice = append(q.slice, val)
}

func (q *Queue) Empty() bool {
    return len(q.slice) == 0
}
