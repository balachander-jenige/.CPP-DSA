/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    // Step 1: Inorder traversal to get sorted values
    var nums []int
    inorder(root, &nums)

    // Step 2: Build balanced BST from sorted array
    return buildBST(nums, 0, len(nums)-1)
}

// Inorder traversal
func inorder(root *TreeNode, nums *[]int) {
    if root == nil {
        return
    }

    inorder(root.Left, nums)
    *nums = append(*nums, root.Val)
    inorder(root.Right, nums)
}

// Build balanced BST
func buildBST(nums []int, left int, right int) *TreeNode {
    if left > right {
        return nil
    }

    mid := left + (right-left)/2

    node := &TreeNode{
        Val: nums[mid],
    }

    node.Left = buildBST(nums, left, mid-1)
    node.Right = buildBST(nums, mid+1, right)

    return node
}
