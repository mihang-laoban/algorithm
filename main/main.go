package main

import (
	"dp/ds"
)

func main() {
	tmp := ds.LinkedList{}
	tmp.Insert(1)
	tmp.Insert(2)
	tmp.Insert(3)
	tmp.Insert(4)
	tmp.Insert(5)
	tmp.Display()
	tmp.Display()
	tmp.Display()
}
