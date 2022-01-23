package leetcode

import (
    "fmt"
    "testing"
)

var nums = []int{2, 7, 11, 15}
var target int = 9

/**
 * leetcode 1
 * tow sum
 */
func TestTwoSum(t *testing.T) {
    assist := make(map[int]int)
    res := make([]int, 2)

    for i, num := range nums {
        if index, ok := assist[target-num]; ok {
            res[0] = index
            res[1] = i
            break
        }
        assist[num] = i
    }

    t.Logf("result: %v", res)
}

/**
 * leetcode 2
 * Add Two Numbers
 */
type ListNode struct {
    Val  int
    Next *ListNode
}

func TestAddTwoNum(t *testing.T) {
    l1, l2 := initTestCase2()

    l3 := &ListNode{0, nil}
    extra := 0
    cur := &ListNode{0, nil}
    l3.Next = cur

    for {
        if l1 == nil && l2 == nil {
            break
        }

        if l1 != nil && l2 != nil {
            cur.Val = l1.Val + l2.Val + extra
            l1 = l1.Next
            l2 = l2.Next
        } else if l1 != nil {
            cur.Val = l1.Val + extra
            l1 = l1.Next
        } else if l2 != nil {
            cur.Val = l2.Val + extra
            l2 = l2.Next
        }

        if cur.Val >= 10 {
            extra = 1
            cur.Val = cur.Val % 10
        } else {
            extra = 0
        }

        if l2 != nil || l1 != nil || extra != 0 {
            cur.Next = &ListNode{extra, nil}
            cur = cur.Next
        }
    }

    printListNode(l3.Next)
}

func initTestCase1() (*ListNode, *ListNode) {
    l1 := &ListNode{2, nil}
    l1.Next = &ListNode{4, nil}
    l1.Next.Next = &ListNode{3, nil}

    l2 := &ListNode{5, nil}
    l2.Next = &ListNode{6, nil}
    l2.Next.Next = &ListNode{4, nil}

    return l1, l2
}

func initTestCase2() (*ListNode, *ListNode) {
    l1 := &ListNode{0, nil}
    l2 := &ListNode{0, nil}

    return l1, l2
}

func initTestCase3() (*ListNode, *ListNode) {
    l1 := &ListNode{9, nil}
    l1.Next = &ListNode{9, nil}
    l1.Next.Next = &ListNode{9, nil}
    l1.Next.Next.Next = &ListNode{9, nil}
    l1.Next.Next.Next.Next = &ListNode{9, nil}
    l1.Next.Next.Next.Next.Next = &ListNode{9, nil}
    l1.Next.Next.Next.Next.Next.Next = &ListNode{9, nil}

    l2 := &ListNode{9, nil}
    l2.Next = &ListNode{9, nil}
    l2.Next.Next = &ListNode{9, nil}
    l2.Next.Next.Next = &ListNode{9, nil}

    return l1, l2
}

func printListNode(l *ListNode) {
    head := l

    for ; head != nil; head = head.Next {
        fmt.Print(head.Val)
        fmt.Print(" ")
    }

    fmt.Println()
}

/**
 * leetcode 3
 * Longest Substring Without Repeating Characters
 */
var s = "abcabcbb"
var s1 = "au"
var s2 = "bbbbb"
var s3 = "pwwkew"
var s4 = " "
var s5 = "dvdf"

func TestLeetCode(t *testing.T) {

    max := LongestSubstring(s5)

    t.Logf("len: %d", max)
}

func LongestSubstring(s string) int {
    i := 0
    j := i + 1
    max := 0

    if len(s) == 0 {
        return 0
    }

    for ; j < len(s); j++ {
        flag := checkDuplicate(s[i : j+1])
        if flag {
            val := j - i
            if val > max {
                max = val
            }
            i++
        }
    }
    val := j - i
    if val > max {
        max = val
    }

    return max
}

func checkDuplicate(s string) bool {
    dup := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        if _, ok := dup[s[i]]; ok {
            return true
        }
        dup[s[i]] = 0
    }

    return false
}
