package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v any) *ListItem
	PushBack(v any) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value any
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len         int
	first, last *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v any) *ListItem {
	l.len++

	item := newListItem(v)

	if l.first == nil {
		l.first = item
		l.last = item

		return item
	}

	item.Next = l.first
	l.first.Prev = item
	l.first = item

	return item
}

func (l *list) PushBack(v any) *ListItem {
	l.len++

	item := newListItem(v)

	if l.last == nil {
		l.first = item
		l.last = item

		return item
	}

	item.Prev = l.last
	l.last.Next = item
	l.last = item

	return item
}

func (l *list) Remove(i *ListItem) {
	if l.len == 1 {
		l.first = nil
		l.last = nil
		l.len = 0

		return
	}

	l.len--

	if i.Prev == nil {
		l.first = i.Next
		l.first.Prev = nil

		return
	}

	if i.Next == nil {
		l.last = i.Prev
		l.last.Next = nil

		return
	}

	i.Prev.Next = i.Next
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	i.Prev.Next = i.Next

	i.Next = l.first

	if l.first.Prev == nil {
		l.first.Prev = i
	}

	i.Prev = nil

	l.first = i
}

func newListItem(v any) *ListItem {
	return &ListItem{
		Value: v,
	}
}

func NewList() List {
	return new(list)
}
