module github.com/workhorse/scheduler

go 1.14

replace github.com/workhorse/api => ../api/

replace github.com/workhorse/client => ../client/

replace github.com/workhorse/commons => ../commons/

require (
	github.com/workhorse/api v0.0.0
	github.com/workhorse/commons v0.0.0 // indirect
	github.com/workhorse/client v0.0.0
)
