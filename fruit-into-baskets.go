import (
	"fmt"
)

func totalFruit(tree []int) int {
	ans := 0
	for i := 0; i < len(tree); i++ {
		treeTypeStore := make(map[int]bool)
		n := 0
		ians := 0
		if len(tree)-i-1 < ans {
			return ans
		}

		for j := i; j < len(tree); j++ {
			if ok, _ := treeTypeStore[tree[j]]; !ok && n == 2 {
				break
			}
			if ok, _ := treeTypeStore[tree[j]]; !ok {
				treeTypeStore[tree[j]] = true
				n++
			}
			ians++
		}
		ans = max(ans, ians)
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}