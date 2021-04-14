/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res map[int]bool
var nodes []*TreeNode
var nodeParents map[*TreeNode]*TreeNode

func distanceK(root *TreeNode, target *TreeNode, dis int) []int {
	if dis == 0 {
		return []int{target.Val}
	}
	res = make(map[int]bool, 0)
	nodes = make([]*TreeNode, 0)
	nodeParents = make(map[*TreeNode]*TreeNode)
	nodeParents[root] = nil
	fillNodes(root)
	for _, node := range nodes {
		if node.Right == target {
			if dis == 1 {
				res[node.Val] = true
			}
			cascade(node.Left, dis-2)
			cascadeBack(nodeParents[node], node, dis-2)
		} else if node.Left == target {
			if dis == 1 {
				res[node.Val] = true
			}
			cascade(node.Right, dis-2)
			cascadeBack(nodeParents[node], node, dis-2)
		} else if node == target {
			cascade(node.Right, dis-1)
			cascade(node.Left, dis-1)
			cascadeBack(nodeParents[node], node, dis-1)
		} else {

		}
	}

	fmt.Println(res)
	keys := make([]int, 0, len(res))
	for k := range res {
		keys = append(keys, k)
	}
	return keys
}

// func findTarget(node *TreeNode, target *TreeNode,k int) int {
//     if (node == nil){
//         return  -1
//     }
//     if node == target {
//         return 0
//     }
//     if k == 0  && node != target{
//         return -1
//     }

//     fmt.Println(node)
//     out1 := findTarget(node.Right,target, k -1)
//     if out1 == 0 {
//         return 0
//     }
//     out2 := findTarget(node.Left,target, k -1)
//     if out2 == 0 {
//         return 0
//     }
//     fmt.Println("-------------")
//     return -1
// }

func fillNodes(root *TreeNode) {
	if root == nil {
		return
	}
	nodes = append(nodes, root)
	nodeParents[root.Right] = root
	nodeParents[root.Left] = root

	fillNodes(root.Right)
	fillNodes(root.Left)
}

func cascade(root *TreeNode, dis int) {
	if root == nil {
		return
	}
	if dis == 0 {
		res[root.Val] = true
		return
	}
	if dis < 0 {
		return
	}
	cascade(root.Right, dis-1)
	cascade(root.Left, dis-1)

}

func cascadeBack(root *TreeNode, prev *TreeNode, dis int) {
	if root == nil {
		return
	}
	if dis == 0 {
		res[root.Val] = true
		return
	}
	if dis < 0 {
		return
	}
	if nodeParents[root] == nil {
		if root.Right == prev {
			cascade(root.Left, dis-1)
		} else {
			cascade(root.Right, dis-1)
		}
	}
	cascadeBack(nodeParents[root], root, dis-1)
	if root.Right == prev {
		cascade(root.Left, dis-1)
	} else {
		cascade(root.Right, dis-1)
	}

}