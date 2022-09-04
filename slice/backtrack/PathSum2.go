package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func backtrack4PathSum2(root *TreeNode, pathSum, targetSum int, path []int, res *[][]int)  {
	pathSum += root.Val					//计算是否满足条件前，需要加上当前这个节点的值
	path = append(path, root.Val)

	if pathSum == targetSum && root.Left == nil && root.Right == nil {
		t := make([]int, len(path))
		copy(t, path)
		*res = append(*res, t)		//没有直接return，需要回退当前节点的值
	} else {
		if root.Left != nil {
			backtrack4PathSum2(root.Left, pathSum, targetSum, path, res)
		}
		if root.Right != nil {
			backtrack4PathSum2(root.Right, pathSum, targetSum, path, res)
		}
	}

	pathSum -= path[len(path)-1]	//不管是否满足条件，两种情况都需要回退
	path = path[:len(path)-1]
}

func pathSum2(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	pathSum := 0
	var path []int
	var res [][]int
	backtrack4PathSum2(root, pathSum, targetSum, path, &res)
	return res
}
