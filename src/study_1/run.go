package main

import "fmt"
import "./Node"

func main() {

	head := &Node.Tree{10, nil, nil}
	head.Left = &Node.Tree{6, nil, nil}
	head.Right = &Node.Tree{14, nil, nil}
	head.Left.Left = &Node.Tree{4, nil, nil}
	head.Left.Right = &Node.Tree{8, nil, nil}
	head.Right.Left = &Node.Tree{12, nil, nil}
	head.Right.Right = &Node.Tree{16, nil, nil}


	//head = head.Convert()
	head = head.ConvertBySlice()


	for head != nil{
		fmt.Println(head.Value)
		head = head.Left
	}



}

