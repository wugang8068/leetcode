package tests

import (
	"fmt"
	"testing"
)

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	     Val int
	     Next *ListNode
}


func initialNodes() (l1 *ListNode, l2 *ListNode){
	return &ListNode{
		Val: 0,
		Next: nil,
	}, &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 3,
			Next: nil,
		},
	}
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 == nil {
			if l2 != nil {
				l1 = l2
				l2 = nil
			}
		}
		return l1
	}
	l1.Val = l1.Val + l2.Val
	if l1.Val > 9 {
		l1.Val = l1.Val - 10
		if l1.Next != nil && l2.Next != nil {
			l1.Next.Val += 1
		}else{
			if l1.Next == nil {
				l1.Next = &ListNode{
					Val: 1,
					Next: nil,
				}
			}else if l2.Next == nil {
				l2.Next = &ListNode{
					Val: 1,
					Next: nil,
				}
			}
		}
	}
	l1.Next = addTwoNumbers(l1.Next, l2.Next)
	return l1
}

func TestTo(t *testing.T)  {
	l1, l2 := initialNodes()
	l := addTwoNumbers(l1, l2)
	fmt.Println(*l)
}
