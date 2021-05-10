package main

import "fmt"

type Node struct {
	key   string
	count int
	next  *Node
	prev  *Node
}

type AllOne struct {
	head    *Node
	tail    *Node
	nodeMap map[string]*Node
}

func Constructor() AllOne {
	return AllOne{head: nil, tail: nil, nodeMap: make(map[string]*Node)}
}

func (all *AllOne) insert(key string) {
	n := Node{count: 1, key: key, next: all.head, prev: nil}
	if all.head != nil {
		all.head.prev = &n
	}
	all.head = &n
	if all.tail == nil {
		all.tail = &n
	}
	all.nodeMap[key] = &n
	return
}

func (all *AllOne) remove(n *Node) {
	if all.head == all.tail {
		all.head = nil
		all.tail = nil
	} else if n == all.head {
		n.next.prev = nil
		all.head = n.next
	} else if n == all.tail {
		n.prev.next = nil
		all.tail = n.prev
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
	}
}

func (all *AllOne) shiftRight(n *Node) {
	if n.next != nil && n.count > n.next.count {
		if n.prev != nil {
			n.prev.next = n.next
		} else {
			all.head = n.next
		}
		n.next.prev = n.prev
		curr := n.next
		for curr != nil && n.count > curr.count {
			curr = curr.next
		}
		if curr == nil {
			n.prev = all.tail
			all.tail.next = n
			all.tail = n
			all.tail.next = nil
		} else {
			n.prev = curr.prev
			n.next = curr
			curr.prev.next = n
			curr.prev = n
		}
	}
}

func (all *AllOne) shiftLeft(n *Node) {
	if n.prev != nil && n.count < n.prev.count {
		if n.next != nil {
			n.next.prev = n.prev
		} else {
			all.tail = n.prev
		}
		n.prev.next = n.next
		curr := n.prev
		for curr != nil && n.count < curr.count {
			curr = curr.prev
		}
		if curr == nil {
			n.next = all.head
			all.head.prev = n
			all.head = n
			all.head.prev = nil
		} else {
			n.prev = curr
			n.next = curr.next
			curr.next.prev = n
			curr.next = n
		}
	}
}

func (all *AllOne) Inc(key string) {
	if n, exists := all.nodeMap[key]; exists {
		n.count++
		all.shiftRight(n)
	} else {
		all.insert(key)
	}
}

func (all *AllOne) Dec(key string) {
	n := all.nodeMap[key]
	if n.count > 1 {
		n.count--
		all.shiftLeft(n)
	} else {
		all.remove(n)
	}
}

func (all *AllOne) GetMaxKey() string {
	if all.tail == nil {
		return ""
	}
	return all.tail.key
}

func (all *AllOne) GetMinKey() string {
	if all.head == nil {
		return ""
	}
	return all.head.key
}

func (all *AllOne) print() {
	for n := all.head; n != nil; n = n.next {
		if n != all.tail {
			fmt.Printf("%s,%d -> ", n.key, n.count)
		} else {
			fmt.Printf("%s,%d ", n.key, n.count)
		}
	}
	fmt.Println()
}

func main() {
	a := Constructor()
	a.Inc("hello")
	a.print()
	a.Inc("world")
	a.print()
	a.Inc("hello")
	a.print()

	a.Dec("world")
	a.print()

	a.Inc("hello")
	a.print()
	a.Inc("leet")
	a.print()

	a.Dec("hello")
	a.print()
	a.Dec("hello")
	a.print()
	a.Dec("hello")
	a.print()

	fmt.Print(a.GetMaxKey())
}
