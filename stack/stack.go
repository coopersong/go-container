package stack

type Stack struct {
    slice []interface{}
}

func (s *Stack) Push(val interface{}) {
    s.slice = append(s.slice, val)
}

func (s *Stack) Pop() interface{} {
    if len(s.slice) > 0 {
        top := s.slice[len(s.slice)-1]
        s.slice = s.slice[:len(s.slice)-1]
        return top
    }

    return nil
}

func (s *Stack) Top() interface{} {
    if len(s.slice) > 0 {
        return s.slice[len(s.slice)-1]
    }

    return nil
}

func (s *Stack) Empty() bool {
    return len(s.slice) == 0
}
