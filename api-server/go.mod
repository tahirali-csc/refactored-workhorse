module github.com/workhorse/apiserver

go 1.14

replace github.com/workhorse/api => ../api/

require (
	github.com/lib/pq v1.8.0
	github.com/workhorse/api v0.0.0
	gopkg.in/yaml.v2 v2.3.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
