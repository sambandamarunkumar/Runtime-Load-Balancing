type Pod struct {
	Name string
}
type Node struct {
	Name string
	Pods []Pod
}
func main() {
	nodes := []Node{
		{Name: "node1"},
		{Name: "node2"},
		{Name: "node3"},
	}
	pods := []Pod{
		{"p1"}, {"p2"}, {"p3"}, {"p4"},
		{"p5"}, {"p6"}, {"p7"}, {"p8"},
	}
	for i, pod := range pods {
		index := i % len(nodes)
		nodes[index].Pods = append(nodes[index].Pods, pod)
	}
	for _, node := range nodes {
		fmt.Print(node.Name, ": ")
		for _, pod := range node.Pods {
			fmt.Print(pod.Name, " ")
		}
		fmt.Println()
	}
}
