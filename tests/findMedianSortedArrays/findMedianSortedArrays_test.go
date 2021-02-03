package findMedianSortedArrays
//
//import (
//	"fmt"
//	"testing"
//)
//
////给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
////
////请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
////
////你可以假设 nums1 和 nums2 不会同时为空。
////
////示例 1:
////
////nums1 = [1, 3]
////nums2 = [2]
////
////则中位数是 2.0
////示例 2:
////
////nums1 = [1, 2]
////nums2 = [3, 4]
////
////则中位数是 (2 + 3)/2 = 2.5
////
////来源：力扣（LeetCode）
////链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
////著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	n1Len := len(nums1)
//	n2Len := len(nums2)
//	//计算中位数的位置
//	l := n1Len + n2Len
//	left := -1
//	right := -1
//	if l % 2 == 0 {
//		left = l / 2  - 1
//		right = l / 2
//	}else {
//		left = (l - 1) / 2
//	}
//	start := 0
//	for {
//		if start == left {
//			if start < n1Len && start < n2Len {
//
//			}
//		}
//
//		if start < n1Len {
//			if start < n2Len {
//				if nums1[start] < nums2[start] {
//					if start == left {
//						if right > -1 {
//							return
//						}
//					}
//					start++
//				}
//			}
//		}
//	}
//
//}
//
//func TestFindMedianSortedArrays(t *testing.T)  {
//	fmt.Println(findMedianSortedArrays([]int{
//		1,2,3,4,5,7,
//	}, []int{
//		3,4,6,
//	}))
//}
