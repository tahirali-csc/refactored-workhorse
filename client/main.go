package main

import (
	"sync"
)



func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//
	//b := Builds{}
	//go func() {
	//	b.Watch("http://localhost:8084/events", func(obj interface{}) {
	//		build := obj.(*api.Build)
	//		log.Println("Build::::", build)
	//	})
	//}()
	//

	//go func(){
	//	b := api.Builds{}
	//	b.Watch("http://localhost:8084/events", func(obj interface{}) {
	//		//		build := obj.(*api.Build)
	//		//		log.Println("Build::::", build)
	//		//	})
	//}()
	//wg.Wait()
}
