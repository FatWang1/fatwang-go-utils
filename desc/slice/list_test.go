package slice

import (
	"reflect"
	"testing"
)

func TestContain(t *testing.T) {
	type args[T comparable] struct {
		list []T
		item T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "contain",
			args: args[int]{
				list: []int{0},
				item: 0,
			},
			want: true,
		},
		{
			name: "not contain",
			args: args[int]{
				list: []int{1},
				item: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.list, tt.args.item); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	type args[T any] struct {
		idx  int
		list []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	list := []string{"name"}
	tests := []testCase[string]{
		{
			name: "all is ok",
			args: args[string]{
				idx:  0,
				list: list,
			},
			want: "name",
		},
		{
			name: "out of range",
			args: args[string]{
				idx:  1,
				list: list,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetItem(tt.args.idx, tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveListElement(t *testing.T) {
	tests := []struct {
		list []int
		val  int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{[]int{5, 5, 5, 5, 5}, 5, []int{}},
		{[]int{1, 2, 3, 4, 5}, 6, []int{1, 2, 3, 4, 5}},
		{[]int{2, 4, 6, 8, 10}, 4, []int{2, 6, 8, 10}},
	}

	for _, test := range tests {
		got := RemoveItemByValue(test.list, test.val)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("RemoveItemByValue(%v, %v) = %v; want %v", test.list, test.val, got, test.want)
		}
	}
}

func TestRemoveListElementInPlace(t *testing.T) {
	tests := []struct {
		list []int
		val  int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{[]int{5, 5, 5, 5, 5}, 5, []int{}},
		{[]int{1, 2, 3, 4, 5}, 6, []int{1, 2, 3, 4, 5}},
		{[]int{2, 4, 6, 8, 10}, 4, []int{2, 6, 8, 10}},
	}

	for _, test := range tests {
		got := RemoveItemByValueInPlace(test.list, test.val)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("RemoveItemByValueInPlace(%v, %v) = %v; want %v", test.list, test.val, got, test.want)
		}
	}
}
