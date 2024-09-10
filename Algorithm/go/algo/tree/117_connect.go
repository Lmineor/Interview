package tree

/*
NextNode
116和117类似
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。
*/
type NextNode struct {
	Val   int
	Left  *NextNode
	Right *NextNode
	Next  *NextNode
}
func CBuild(){
	t := &NextNode{
		Val:   1,
		Left:  &NextNode{
			Val:   2,
			Left:  &NextNode{Val: 4},
			Right: &NextNode{Val: 5},
		},
		Right: &NextNode{Val: 3,
			Left: &NextNode{Val: 6},
		Right: &NextNode{Val: 7}},
	}
	connect(t)
}
func connect(root *NextNode) *NextNode {
	var allLevel []*NextNode
	var currentLevelLen int
	allLevel = append(allLevel, root) // 根节点入队列
	levelNodes := make([]*NextNode, 0)
	for len(allLevel) > 0 {
		currentLevelLen = len(allLevel)
		levelNodes = []*NextNode{}
		for i := 0; i < currentLevelLen; i++ {
			node := allLevel[0]
			allLevel = allLevel[1:]
			levelNodes = append(levelNodes, node)
			if node.Left != nil {
				allLevel = append(allLevel, node.Left)
			}
			if node.Right != nil {
				allLevel = append(allLevel, node.Right)
			}
			linkLevelNodes(levelNodes)
		}
	}
	return root
}

func linkLevelNodes(nodes []*NextNode){

	for i := 0; i<len(nodes)-1;i++{
		nodes[i].Next = nodes[i+1]
		print(nodes[i].Val)
	}
	nodes[len(nodes)-1].Next = nil
}