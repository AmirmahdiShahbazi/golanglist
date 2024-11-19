package list

type List[T any] struct {
	data []T
}

func NewList[T any]() *List[T] {
	return &List[T]{
		data: []T{},
	}
}

func (list *List[T]) Add(value T) *List[T] {
	list.data = append(list.data, value)
	return list
}

func (list *List[T]) Find(index int) T {
	if index >= len(list.data) {
		panic("Index out of Range")
	}
	return list.data[index]
}

func (list *List[T]) InsertAt(index int, value T) *List[T] {
	if index >= len(list.data) {
		panic("Index out of Range")
	}
	list.data = append(list.data, value)
	copy(list.data[index+1:], list.data[index:])
	list.data[index] = value
	return list
}

func (list *List[T]) Get() []T {
	return list.data
}

func (list *List[T]) RemoveFrom(index int) *List[T] {
	if index < 0 || index >= len(list.data) {
		panic("Index out of range")
	}
	copy(list.data[index:], list.data[index+1:])
	list.data = list.data[:len(list.data)-1]
	return list
}

func (list *List[T]) Pop() T {
	if len(list.data) == 0 {
		panic("pop from empty list")
	}
	lastIndex := len(list.data) - 1  
	lastElement := list.data[lastIndex]  
	list.data = list.data[:lastIndex]  
	return lastElement
}