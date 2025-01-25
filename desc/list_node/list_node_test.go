package list_node

import (
	"reflect"
	"testing"
)

func TestListNodeify(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "all is ok",
			args: args{
				list: []int{1, 2, 3, 4, 5},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "nil list",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListNodeify(tt.args.list...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNodeify() = %v, want %v", got, tt.want)
			}
		})
	}
}
