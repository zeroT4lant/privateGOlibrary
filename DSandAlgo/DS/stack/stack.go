package main

type Stack struct {
	data []int
}

func main() {

}

func (s *Stack) Push(el int) {
	s.data = append(s.data, el)
}

func (s *Stack) Pop() int {
	if s.isEmpty() {
		return -1
	}
	//последний элемент слайса
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s Stack) isEmpty() bool {
	return len(s.data) == 0
}
