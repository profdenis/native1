package list

import (
	"errors"
	"testing"
)

func TestIntList_IsEmpty(t *testing.T) {
	type fields struct {
		start *intNode
		end   *intNode
		size  int
	}
	testCases := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test IsEmpty True",
			fields: fields{
				start: nil,
				end:   nil,
				size:  0,
			},
			want: true,
		},
		{
			name: "Test IsEmpty False",
			fields: fields{
				start: &intNode{value: 10},
				end:   &intNode{value: 20},
				size:  2,
			},
			want: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			list := &IntList{
				start: tt.fields.start,
				end:   tt.fields.end,
				size:  tt.fields.size,
			}
			if got := list.IsEmpty(); got != tt.want {
				t.Errorf("IntList.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntList_Size(t *testing.T) {
	testCases := []struct {
		name string
		list *IntList
		want int
	}{
		{
			name: "Empty List",
			list: NewIntList(),
			want: 0,
		},
		{
			name: "Single Element",
			list: NewIntList(1),
			want: 1,
		},
		{
			name: "Multiple Elements",
			list: NewIntList(1, 2, 3, 4, 5),
			want: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.Size()
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIntList_String(t *testing.T) {
	tests := []struct {
		name string
		list *IntList
		want string
	}{
		{
			name: "Empty List",
			list: NewIntList(),
			want: "",
		},
		{
			name: "Single Element List",
			list: NewIntList(5),
			want: "5 --| ",
		},
		{
			name: "Multiple Element List",
			list: NewIntList(1, 2, 3, 4, 5),
			want: "1 --> 2 --> 3 --> 4 --> 5 --| ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.String(); got != tt.want {
				t.Errorf("IntList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntList_AddStart(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		element  int
		expected *IntList
	}{
		{
			name:     "TestIntList_AddStart_WithEmptyList",
			input:    []int{},
			element:  5,
			expected: NewIntList(5),
		},
		{
			name:     "TestIntList_AddStart_WithNonEmptyList",
			input:    []int{2, 3, 4},
			element:  1,
			expected: NewIntList(1, 2, 3, 4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewIntList(tt.input...)
			list.AddStart(tt.element)

			if !equal(list, tt.expected) {
				t.Errorf("got %v, want %v", list, tt.expected)
			}
		})
	}
}

func TestIntList_AddEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		element  int
		expected *IntList
	}{
		{
			name:     "TestIntList_AddEnd_WithEmptyList",
			input:    []int{},
			element:  5,
			expected: NewIntList(5),
		},
		{
			name:     "TestIntList_AddEnd_WithNonEmptyList",
			input:    []int{2, 3, 4},
			element:  1,
			expected: NewIntList(2, 3, 4, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewIntList(tt.input...)
			list.AddEnd(tt.element)

			if !equal(list, tt.expected) {
				t.Errorf("got %v, want %v", list, tt.expected)
			}
		})
	}
}

func TestIntList_RemoveStart(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedValue int
		expectedList  *IntList
		err           error
	}{
		{
			"Remove from empty list",
			[]int{},
			0,
			NewIntList(),
			errors.New("empty list, cannot remove"),
		},
		{
			"Remove from list with one element",
			[]int{10},
			10,
			NewIntList(),
			nil,
		},
		{
			"Remove from list with multiple elements",
			[]int{10, 20, 30},
			10,
			NewIntList(20, 30),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewIntList(tt.input...)

			got, err := list.RemoveStart()

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}

			if got != tt.expectedValue {
				t.Errorf("expected element: %d, got: %d", tt.expectedValue, got)
			}

			if tt.err == nil && list.Size() != len(tt.input)-1 {
				t.Errorf("expected size after removal: %d, got: %d", len(tt.input)-1, list.Size())
			}

			if !equal(list, tt.expectedList) {
				t.Errorf("got %v, want %v", list, tt.expectedList)
			}
		})
	}
}

func TestIntList_RemoveEnd(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedValue int
		expectedList  *IntList
		err           error
	}{
		{
			"Remove from empty list",
			[]int{},
			0,
			NewIntList(),
			errors.New("empty list, cannot remove"),
		},
		{
			"Remove from list with one element",
			[]int{10},
			10,
			NewIntList(),
			nil,
		},
		{
			"Remove from list with multiple elements",
			[]int{10, 20, 30},
			30,
			NewIntList(10, 20),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewIntList(tt.input...)

			got, err := list.RemoveEnd()

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}

			if got != tt.expectedValue {
				t.Errorf("expected element: %d, got: %d", tt.expectedValue, got)
			}

			if tt.err == nil && list.Size() != len(tt.input)-1 {
				t.Errorf("expected size after removal: %d, got: %d", len(tt.input)-1, list.Size())
			}

			if !equal(list, tt.expectedList) {
				t.Errorf("got %v, want %v", list, tt.expectedList)
			}
		})
	}
}

func TestIntList_Insert(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		element  int
		pos      int
		expected *IntList
		err      error
	}{
		{
			name:     "FirstElement",
			elements: []int{},
			element:  3,
			pos:      0,
			expected: NewIntList(3),
			err:      nil,
		},
		{
			name:     "MiddleElement",
			elements: []int{1, 2, 4},
			element:  3,
			pos:      2,
			expected: NewIntList(1, 2, 3, 4),
			err:      nil,
		},
		{
			name:     "LastErrorFree",
			elements: []int{1, 2, 3},
			element:  4,
			pos:      3,
			expected: NewIntList(1, 2, 3, 4),
			err:      nil,
		},
		{
			name:     "OutOfBounds",
			elements: []int{1, 2, 3},
			element:  5,
			pos:      4,
			expected: NewIntList(1, 2, 3),
			err:      errors.New("invalid position"),
		},
		{
			name:     "NegativeIndex",
			elements: []int{1, 2, 3},
			element:  5,
			pos:      -1,
			expected: NewIntList(1, 2, 3),
			err:      errors.New("invalid position"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewIntList(tt.elements...)
			err := list.Insert(tt.element, tt.pos)

			if (err != nil || tt.err != nil) && (err == nil || tt.err == nil || err.Error() != tt.err.Error()) {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}

			if !equal(list, tt.expected) {
				t.Fatalf("expected list %v, got %v", tt.expected, list)
			}
		})
	}
}

func TestIntList_Remove(t *testing.T) {
	tests := []struct {
		name           string
		elements       []int
		position       int
		expectedSize   int
		expectedError  string
		expectedResult int
		expectedList   *IntList
	}{
		{
			name:           "Empty List",
			elements:       []int{},
			position:       0,
			expectedSize:   0,
			expectedError:  "invalid position",
			expectedResult: 0,
			expectedList:   NewIntList(),
		},
		{
			name:           "Removing Start",
			elements:       []int{1, 2, 3},
			position:       0,
			expectedSize:   2,
			expectedError:  "",
			expectedResult: 1,
			expectedList:   NewIntList(2, 3),
		},
		{
			name:           "Removing End",
			elements:       []int{1, 2, 3},
			position:       2,
			expectedSize:   2,
			expectedError:  "",
			expectedResult: 3,
			expectedList:   NewIntList(1, 2),
		},
		{
			name:           "Removing from Middle",
			elements:       []int{1, 2, 3},
			position:       1,
			expectedSize:   2,
			expectedError:  "",
			expectedResult: 2,
			expectedList:   NewIntList(1, 3),
		},
		{
			name:           "Removing from Invalid Position",
			elements:       []int{1, 2, 3},
			position:       4,
			expectedSize:   3,
			expectedError:  "invalid position",
			expectedResult: 0,
			expectedList:   NewIntList(1, 2, 3),
		},
		{
			name:           "Removing from middle of bigger list",
			elements:       []int{1, 2, 3, 4, 5, 6, 7},
			position:       5,
			expectedSize:   6,
			expectedError:  "",
			expectedResult: 6,
			expectedList:   NewIntList(1, 2, 3, 4, 5, 7),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewIntList(tt.elements...)
			actualResult, err := list.Remove(tt.position)

			if (err != nil) != (tt.expectedError != "") || (err != nil && err.Error() != tt.expectedError) {
				t.Errorf("Expected error %v, but got %v", tt.expectedError, err)
			}

			if tt.expectedResult != actualResult {
				t.Errorf("Expected result %v, but got %v", tt.expectedResult, actualResult)
			}

			if tt.expectedSize != list.Size() {
				t.Errorf("Expected size %v, but got %v", tt.expectedSize, list.Size())
			}

			if !equal(list, tt.expectedList) {
				t.Fatalf("expected list %v, got %v", tt.expectedList, list)
			}
		})
	}
}

func equal(list1, list2 *IntList) bool {
	if list1 == nil && list2 == nil {
		return true
	}
	if list1 != nil && list2 != nil {
		if list1.Size() != list2.Size() {
			return false
		}
		start1, start2 := list1.start, list2.start
		for start1 != nil && start2 != nil {
			if start1.value != start2.value {
				return false
			}
			start1 = start1.next
			start2 = start2.next
		}
		if start1 != nil || start2 != nil {
			return false
		}
		return true
	}
	return false
}
