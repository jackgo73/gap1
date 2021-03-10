package p_00002_add_two_numbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1",
			args{
				&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
				&ListNode{5, &ListNode{6, &ListNode{4, nil}}}},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}}},
		{"case2",
			args{
				&ListNode{0, nil},
				&ListNode{0, nil}},
			&ListNode{0, nil}},
		{"case3",
			args{
				&ListNode{9, &ListNode{9, &ListNode{9, nil}}},
				&ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			&ListNode{8, &ListNode{9, &ListNode{9, &ListNode{1, nil}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
