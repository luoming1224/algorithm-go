package main

// 参考 https://zhuanlan.zhihu.com/p/441757437 稍作改动
// https://blog.csdn.net/a6661314/article/details/122918223
func inOrder(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	var res []int
	var stack []*TreeNode
	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		curNode = stack[len(stack)-1]
		res = append(res, curNode.Val)
		stack = stack[:len(stack)-1]
		curNode = curNode.Right
	}

	return res
}
