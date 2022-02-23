package _33

type Node struct {
	Val       int
	Neighbors []*Node
}

//给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
//
//图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
func cloneGraph(node *Node) *Node {
 //注意被访问过的结点要标记
	visited := make(map[*Node]*Node)
	return clone(node, visited)

}
//用来复制图
func clone(node *Node, visited map[*Node]*Node) *Node {
    if node == nil {
    	return nil
	}
	if v, ok := visited[node]; ok {
		//说明访问过了
		return v
	}
	//创建一个新节点
	newNode := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode

}
