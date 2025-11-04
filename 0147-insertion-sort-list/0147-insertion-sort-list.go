/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	swapped := true
	for swapped {
		swapped = false
		current := head

		for current.Next != nil {
			if current.Val > current.Next.Val {
				current.Val, current.Next.Val = current.Next.Val, current.Val
				swapped = true
			}
			current = current.Next
		}
	}

	return head
    
}


// func InsertSort(arr []int) {

// 	for i, v := range arr {
// 		j := i - 1
		
// 		if i >= 1 {
// 			for j >= 0 && arr[j] > v {
// 				arr[j+1] = arr[j]
// 				j--
// 			}
// 			arr[j+1] = v

// 		}
// 		fmt.Println("arr", arr)

// 	}
// }