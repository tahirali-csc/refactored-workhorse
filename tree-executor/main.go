package main

type Node interface {
	Add(s string, exec Node)
}

type ExecNode struct {
}

func (sn ExecNode) Add(cmd string, exec Node) {

}

type SerialNode struct {
}

func (sn SerialNode) Add(cmd string, exec Node) {

}

type Option struct {
}

func Exec(cmd string, options *Option) Node {
	return ExecNode{}
}

func Serial() Node {
	return SerialNode{}
}

type ParallelNode struct {
}

func (sn ParallelNode) Add(cmd string, exec Node) {

}

func Parallel() Node {
	return ParallelNode{}
}

func pipeline1() Node {
	root := Serial()
	root.Add("", Exec("pyton -v", nil))
	root.Add("", Exec("pyton -v", &Option{}))
	return root
}

func pipeline2() Node {
	root := Parallel()
	root.Add("", Exec("pyton -v", nil))
	root.Add("", Exec("pyton -v", &Option{}))
	return root
}

func pipeline3() Node {
	root := Serial()
	pNode := Parallel()
	pNode.Add("maui", Exec("echo valley isle", nil))
	pNode.Add("maui", Exec("echo valley isle", nil))

	root.Add("")

}

func main() {
	pipeline1()
	pipeline2()
}
