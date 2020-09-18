package api

type BuildNodeBinding struct {
	Id      int
	BuildId int
	NodeId  int
}

type BuildStepNodeBinding struct {
	Id        int
	StepId    int
	IpAddress string
}
