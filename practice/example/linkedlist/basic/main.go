package main

import "fmt"

type HeroNode struct {
	no   int
	name string
	next *HeroNode
}

// 1. insert at the end
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil { // the temp now points to the last node
			break
		}
		temp = temp.next // let temp continue to find the last node address
	}
	temp.next = newHeroNode // add new node to the end of the linkedlist
}

// 2. insert node based on the number in assending order
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := false
	for {
		if temp.next == nil { // at the last node
			break
		} else if temp.next.no > newHeroNode.no { // new node should insert after the temp pointing node
			break
		} else if temp.next.no == newHeroNode.no { // This no has existed
			flag = true
			break
		}
		temp = temp.next // continue to find the proper position
	}
	if flag { // indicate the hero node has existed
		fmt.Printf("Sorry, the hero number no=%v has existed...", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

// 3. List all hero node info
func listHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("Empty Linkedlist")
		return
	}
	for {
		fmt.Printf("[%d, %s]===>\n", temp.next.no, temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

// 4 Delete a specific hero no
func DeleteHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil { // at the end of the node
			break
		} else if temp.next.no == id { // found the target to be deleted
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("id is not found")
	}
}

func main() {
	// create a empty headNode
	head := &HeroNode{}
	hero1 := &HeroNode{no: 1, name: "Ximeng"}
	hero3 := &HeroNode{no: 3, name: "Niu"}
	hero2 := &HeroNode{no: 2, name: "Mai"}
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)
	listHeroNode(head)
	fmt.Println("After deleting a hero node....")
	DeleteHeroNode(head, 3)
	listHeroNode(head)
}
