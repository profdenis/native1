package list

import (
	"errors"
	"fmt"
	"strings"
)

type node[T any] struct {
	element T
	next    *node[T]
}

type SLList[T any] struct {
	size  int
	start *node[T]
	end   *node[T]
}

func NewSLList[T any](elements ...T) *SLList[T] {
	var list SLList[T]
	for _, element := range elements {
		list.AddEnd(element)
	}
	return &list
}

func (list *SLList[T]) String() string {
	if list.IsEmpty() {
		return ""
	}

	// 1 element in the list
	if list.start == list.end {
		return fmt.Sprintf("%v", list.start.element)
	}

	// more than 1 element in the list
	var builder strings.Builder
	current := list.start
	for current.next != nil {
		builder.WriteString(fmt.Sprintf("%v --> ", current.element))
		current = current.next
	}
	builder.WriteString(fmt.Sprintf("%v", current.element))
	return builder.String()
}

func (list *SLList[T]) IsEmpty() bool {
	return list.start == nil
}

func (list *SLList[T]) Size() int {
	return list.size
}

func (list *SLList[T]) AddEnd(element T) {
	newNode := &node[T]{element: element}
	if list.start == nil {
		list.start = newNode
		list.end = newNode
	} else {
		list.end.next = newNode
		list.end = newNode
	}
	list.size++
}

func (list *SLList[T]) AddStart(element T) {
	newNode := &node[T]{element: element}
	if list.start == nil {
		list.start = newNode
		list.end = newNode
	} else {
		newNode.next = list.start
		list.start = newNode
	}
	list.size++
}

func (list *SLList[T]) RemoveEnd() (T, error) {
	if list.start != nil {
		removed := list.end
		if list.start == list.end {
			list.start = nil
			list.end = nil
		} else {
			current := list.start
			for current.next != list.end {
				current = current.next
			}
			current.next = nil
			list.end = current
		}
		list.size--
		return removed.element, nil
	}
	var temp T
	return temp, errors.New("list is empty")
}

func (list *SLList[T]) RemoveStart() (T, error) {
	if list.start != nil {
		removed := list.start
		if list.start == list.end {
			list.start = nil
			list.end = nil
		} else {
			list.start = list.start.next
		}
		list.size--
		return removed.element, nil
	}
	var temp T
	return temp, errors.New("list is empty")
}
