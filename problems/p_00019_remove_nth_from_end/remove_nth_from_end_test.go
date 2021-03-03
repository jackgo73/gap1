package p_00019_remove_nth_from_end

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test__removeNthFromEnd__0(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	expect := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}}
	n := 2

	actual := removeNthFromEnd(head, n)
	for ; expect != nil && actual != nil; actual, expect = actual.Next, expect.Next {
		require.Equal(t, expect.Val, actual.Val)
	}

	require.Nil(t, actual)
	require.Nil(t, expect)
}

func Test__removeNthFromEnd__1(t *testing.T) {
	head := &ListNode{1, nil}
	expect := &ListNode{1, nil}
	n := 1

	actual := removeNthFromEnd(head, n)
	for ; expect != nil && actual != nil; actual, expect = actual.Next, expect.Next {
		require.Equal(t, expect.Val, actual.Val)
	}

	require.Nil(t, actual)
	require.Nil(t, expect)
}