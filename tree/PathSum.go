package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	if root.Left != nil && hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}

// 注意：不能写成如下形式，如下形式会导致还没有遍历完就提前返回false退出了
// if root.Left != nil {
//   return hasPathSum(root.Left, targetSum-root.Val)
// }