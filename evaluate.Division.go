// type Node struct {
// 	name    string
// 	partent *Node
// 	edge    float64
// 	node    *Node
// }

// func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
// 	store := make(map[string]*Node)
// 	for ind, eq := range equations {
// 		if _, ok2 := store[eq[1]]; !ok2 {
// 			store[eq[1]] = &Node{eq[1], 0, nil}
// 		}
// 		if _, ok := store[eq[0]]; !ok {
// 			store[eq[0]] = &Node{eq[0], values[ind], store[eq[1]]}

// 		} else {
// 			store[eq[0]].edge = values[ind]
// 			store[eq[0]].node = store[eq[1]]
// 		}

// 	}

// 	out := make([]float64, 0)
// 	for _, val := range queries {
// 		if _, ok := store[val[0]]; !ok {
// 			out = append(out, -1.0)
// 			continue
// 		}
// 		if _, ok := store[val[1]]; !ok {
// 			out = append(out, -1.0)
// 			continue
// 		}
// 		startNode := store[val[0]]
// 		res := float64(1)
// 		fmt.Println(store[val[1]])

// 		for startNode != nil && startNode != store[val[1]] {
// 			fmt.Println(startNode)

// 			res = res * startNode.edge
// 			startNode = startNode.node
// 		}
// 		if res == 0 {
// 			res = 1
// 			startNode = store[val[1]]

// 			for startNode != nil && startNode != store[val[0]] {
// 				fmt.Println(startNode)

// 				res = res / startNode.edge
// 				startNode = startNode.node
// 			}
// 		}
// 		out = append(out, res)
// 		fmt.Println("RES: ", res)

// 	}
// 	return out

// }

type Rate struct {
	name string
	val  float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	//create graph
	graph := make(map[string][]Rate)
	for i, eq := range equations {
		val := values[i]
		a, b := eq[0], eq[1]
		//add itself
		graph[a] = append(graph[a], Rate{a, 1.0})
		//add both ways
		graph[a] = append(graph[a], Rate{b, val})
		graph[b] = append(graph[b], Rate{a, 1.0 / val})
	}

	//for each query find shortest path using BFS
	res := make([]float64, len(queries))
	for i, q := range queries {
		visited := make(map[string]bool, 0)
		source, target := q[0], q[1]

		//start BFS with source
		queue := []Rate{Rate{source, 1}}

		res[i] = -1
		for len(queue) != 0 && res[i] == -1 {
			//pull from queue
			rate := queue[0]
			queue = queue[1:]

			//mark as visited
			visited[rate.name] = true

			//get the availables rates for current node
			for _, foundRate := range graph[rate.name] {
				if target == foundRate.name {
					//found the target
					res[i] = rate.val * foundRate.val
					break
				} else {
					//will try at next level
					if !visited[foundRate.name] {
						queue = append(queue, Rate{foundRate.name, rate.val * foundRate.val})
					}
				}
			}
		}
	}

	return res

}