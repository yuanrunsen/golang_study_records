// 算法题：二叉搜索树与双向链表
// 题目：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求
// 不能创建任何新的结点，只能调整树中结点指针的指向。

package Node

type Tree struct {
	 Value int
	 Right *Tree
	 Left *Tree
}


//通过切片
func (node *Tree) convertBySlice(link []*Tree){
	if node == nil {
		return
	}

	node.Left.convertBySlice(link)
	if link[0] == nil{
		link[0] = node
	} else {
		link[0].Right = node
		node.Left = link[0]
		link[0] = link[0].Right
	}
	node.Right.convertBySlice(link)
}

func (node *Tree) ConvertBySlice() *Tree{
	link := []*Tree{nil}
	node.convertBySlice(link)
	return link[0]
}



//通过二级指针
func (node *Tree) convert(link **Tree){
	if node == nil {
		return
	}

	node.Left.convert(link)
	if *link == nil{
		*link = node
	} else {
		(*link).Right = node
		node.Left = *link
		*link = (*link).Right
	}
	node.Right.convert(link)
}


func (node *Tree) Convert() *Tree {
	var link *Tree
	node.convert(&link)
	return link
}