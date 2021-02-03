package mergeLinkedList

import (
	"fmt"
	"testing"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	// 连上l1剩余的链,连上l2剩余的链
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	var listNode *ListNode
	for _, list := range lists {
		listNode = mergeList(listNode, list)
	}
	return listNode
}

func TestMerge(t *testing.T)  {
  list := mergeKLists([]*ListNode{
  	&ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	},
	  &ListNode{
		  Val:  1,
		  Next: &ListNode{
			  Val:  2,
			  Next: &ListNode{
				  Val:  4,
				  Next: nil,
			  },
		  },
	  },
	  &ListNode{
		  Val:  3,
		  Next: &ListNode{
			  Val:  4,
			  Next: &ListNode{
				  Val:  5,
				  Next: nil,
			  },
		  },
	  },
  })
  fmt.Println(list)
}