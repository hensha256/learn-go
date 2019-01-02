package dfs

type Node struct {
    value int
    children []*Node
}

func DfsSumOverFive(root *Node) (sum int) {
	if root.value > 5 {
		sum = root.value
	}

	for _, pointer := range root.children {
		sum += DfsSumOverFive(pointer)
	}
	return
}