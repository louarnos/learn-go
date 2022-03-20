package main

import (
	//"errors"
	"fmt"
)

func main() {
//["MyLinkedList","addAtHead","addAtTail","addAtIndex"]
//[[],[1],[3],[3,2]]
	ll := Constructor();
	ll.AddAtHead(1)
	ll.ShowMeTheMoney()
	ll.AddAtTail(3)
	ll.ShowMeTheMoney()
	ll.AddAtIndex(3,2)
	ll.ShowMeTheMoney()
}

type MyLinkedList struct {
    Val *int
    Next *MyLinkedList
}

func Constructor() MyLinkedList {
    head := MyLinkedList{}
    return head
}

func (this *MyLinkedList) ShowMeTheMoney() {

	var a []int
	currentNode := this
	for currentNode != nil {
		a = append(a, *currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Printf("%+v\n", a)
}

func (this *MyLinkedList) Get(index int) int {
    i := 0
    current := this
    for i != index && current != nil {
        current = current.Next
        i++
    }

    if current == nil || current.Val == nil {
        return -1
    }
    return *current.Val
}

func (this *MyLinkedList) AddAtHead(val int)  {
    if this.Val != nil {
        newSecond := MyLinkedList{Val: this.Val, Next: this.Next }
        this.Val = &val
		this.Next = &newSecond
    } else {
        this.Val = &val
    }
}

func (this *MyLinkedList) AddAtTail(val int)  {

	if this.Val == nil {
		this.Val = &val
		return
	}

    curNode := this
    for curNode.Next != nil {
        curNode = curNode.Next
    }
    tailNode := MyLinkedList{Val: &val}
    curNode.Next = &tailNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {

    if index == 0 {
        if this.Val != nil {
            newSecondNode := *this
            this.Next = &newSecondNode
        }
        this.Val = &val
        return
    }

    curIndex := 0
    currentNode := this
    var prevNode *MyLinkedList

    for curIndex != index && currentNode != nil {
        prevNode = currentNode
        currentNode = currentNode.Next
        curIndex++
    }

    newNode := MyLinkedList{ Val: &val, Next: currentNode }
    prevNode.Next = &newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index == 0 {
        if this.Next == nil {
            this.Val = nil
        } else {
            this.Val = this.Next.Val
            newHead := this.Next
            this.Next = newHead.Next
        }
        return
    }

    curIndex := 0
    currentNode := this
    var prevNode *MyLinkedList

    for curIndex != index {

		if currentNode.Next == nil {
			return
		}

        prevNode = currentNode
        currentNode = currentNode.Next
        curIndex++
    }
    prevNode.Next = currentNode.Next
}

func hasCycle(head *MyLinkedList) bool {
    current := head
    m := make(map[*ListNode]int)
    for current != nil && current.Next != nil {
        if _ , ok := m[current]; ok {
            return true
        }
        m[current]= 1
        current = current.Next
    }
    return false
}

func detectCycle(head *ListNode) *ListNode {
    current := head
    m := make(map[*ListNode]int)
    for current != nil && current.Next != nil {
        if _ , ok := m[current]; ok {
            return current
        }
        m[current]= 1
        current = current.Next
    }
    return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    currentA := headA
    currentB := headB

    for {
        for {
            if currentA == currentB {
                return currentA
            }
            if currentB.Next == nil {
                break
            }
            currentB = currentB.Next
        }
        currentB = headB
        if currentA.Next == nil {
            break
        }
        currentA = currentA.Next
    }
    return nil
}

func reverseList(head *ListNode) *ListNode {
    cur := head
    var prev *ListNode
    var nx   *ListNode

    for cur != nil {
        nx = cur.Next
        cur.Next = prev
        prev = cur
        cur = nx
    }

    return prev
}
