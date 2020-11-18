package dp

import "testing"

func TestReverse(t *testing.T){

}

func initList(){
	list := []int{1,2,3,4,5}
	head := Node{}
	for _, v := range list {
		tmp := Node{Val: v}
		if head.Next == nil {
			head = tmp
		} else {

		}
	}

}