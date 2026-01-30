type Pod struct {
	Name string
}

type Node struct {
	Name string
	Load int
	Pods []Pod
}

func main() {
	nodes := []Node{
		{"node1", 2, nil},
		{"node2", 1, nil},
		{"node3", 3, nil},
	}

	pods := []Pod{
		{"p1"}, {"p2"}, {"p3"},
		{"p4"}, {"p5"}, {"p6"},
	}

	for _, pod := range pods {
		min := 0
		for i := range nodes {
			if nodes[i].Load < nodes[min].Load {
				min = i
			}
		}
		nodes[min].Pods = append(nodes[min].Pods, pod)
		nodes[min].Load++
	}

	for _, node := range nodes {
		fmt.Print(node.Name, ": ")
		for _, pod := range node.Pods {
			fmt.Print(pod.Name, " ")
		}
		fmt.Println()
	}
} 
