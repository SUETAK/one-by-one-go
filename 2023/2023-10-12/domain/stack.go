package domain

// NewStack は型パラメータTを宣言する
// T は任意の型である
func NewStack[T any]() *TwinKataStack[T] {
	return &TwinKataStack[T]{
		items: []T{},
	}
}

// T は型パラメータ
type TwinKataStack[T any] struct {
	items []T
}

func (s *TwinKataStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *TwinKataStack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
