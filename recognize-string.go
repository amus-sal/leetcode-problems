func reorganizeString(S string) string {
	dict := make(map[rune]int)
	for _, c := range S {
		dict[c]++
	}
	pq := &pqueue{}
	heap.Init(pq)
	for k, v := range dict {
		heap.Push(pq, item{k, v})
	}
	res := make([]rune, 0)
	for pq.Len() > 0 {
		item1 := heap.Pop(pq).(item)
		if pq.Len() == 0 {
			if len(res) > 0 && res[len(res)-1] == item1.val || item1.count > 1 {
				return ""
			}
			res = append(res, item1.val)
			break
		}
		item2 := heap.Pop(pq).(item)
		res = append(res, item1.val)
		res = append(res, item2.val)

		if item1.count > 1 {
			heap.Push(pq, item{item1.val, item1.count - 1})
		}

		if item2.count > 1 {
			heap.Push(pq, item{item2.val, item2.count - 1})
		}

	}
	return string(res)

}

type item struct {
	val   rune
	count int
}

type pqueue []item

func (pq pqueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq pqueue) Len() int {
	return len(pq)
}

func (pq *pqueue) Push(x interface{}) {
	*pq = append(*pq, x.(item))
}

func (pq *pqueue) Pop() interface{} {
	res := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return res
}