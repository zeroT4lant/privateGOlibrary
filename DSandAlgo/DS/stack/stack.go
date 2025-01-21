package main

func main() {

}

type Stack struct {
	data []int
}

func (s *Stack) Push(el int) {
	s.data = append(s.data, el)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	//Последний элемент
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
