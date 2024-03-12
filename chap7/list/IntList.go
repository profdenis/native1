package list

import (
	"errors"
	"fmt"
	"strings"
)

type intNode struct {
	value int
	next  *intNode
}

type IntList struct {
	start *intNode
	end   *intNode
	size  int
}

func NewIntList(elements ...int) *IntList {
	l := &IntList{}
	for _, element := range elements {
		l.AddEnd(element)
	}
	return l
}
func (list *IntList) IsEmpty() bool {
	return list.start == nil
}

func (list *IntList) Size() int {
	return list.size
}

func (list *IntList) String() string {
	builder := strings.Builder{}
	current := list.start
	for current != nil {
		builder.WriteString(fmt.Sprintf("%v", current.value))
		if current.next == nil {
			builder.WriteString(" --| ")
		} else {
			builder.WriteString(" --> ")
		}
		current = current.next
	}
	return builder.String()
}

func (list *IntList) AddStart(element int) {
	newStart := &intNode{value: element, next: list.start}
	list.start = newStart
	list.size++
}

func (list *IntList) AddEnd(element int) {
	if list.start == nil {
		list.start = &intNode{value: element}
		list.end = list.start
	} else {
		list.end.next = &intNode{value: element}
		list.end = list.end.next
	}
	list.size++
}

func (list *IntList) RemoveStart() (int, error) {
	if list.start == nil {
		return 0, errors.New("empty list, cannot remove")
	}
	element := list.start.value
	list.start = list.start.next
	list.size--
	if list.size == 0 {
		list.end = nil
	}
	return element, nil
}

func (list *IntList) RemoveEnd() (int, error) {
	var element int
	// empty list
	if list.end == nil {
		return element, errors.New("empty list, cannot remove")
	}

	if list.end == list.start {
		// only 1 element input the list
		element = list.end.value
		list.start = nil
		list.end = nil
	} else {
		// at least 2 elements input the list
		newEnd := list.start
		for newEnd.next != list.end {
			newEnd = newEnd.next
		}
		element = list.end.value
		list.end = newEnd
		list.end.next = nil
	}
	list.size--
	return element, nil
}

func (list *IntList) Insert(element int, pos int) error {
	switch {
	case pos < 0 || pos > list.size:
		return errors.New("invalid position")
	case pos == 0:
		list.AddStart(element)
	case pos == list.size:
		list.AddEnd(element)
	default:
		current := list.start
		for i := 0; i < pos-1; i++ {
			current = current.next
		}
		newNode := &intNode{value: element, next: current.next}
		current.next = newNode
		list.size++
	}
	return nil
}

func (list *IntList) Remove(pos int) (int, error) {
	switch {
	case pos < 0 || pos >= list.size:
		return 0, errors.New("invalid position")
	case pos == 0:
		return list.RemoveStart()
	case pos == list.size-1:
		return list.RemoveEnd()
	default:
		current := list.start
		for i := 0; i < pos-1; i++ {
			current = current.next
		}
		element := current.next.value
		current.next = current.next.next
		list.size--
		return element, nil
	}
}
