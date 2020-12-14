package ds

type Node struct {
	Next *Node
	Val  int
}

type LinkedList struct {
	head *Node
}
