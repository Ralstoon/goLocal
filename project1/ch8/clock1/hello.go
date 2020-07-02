package main

func main() {
	var ch = make(chan struct{})
	// ch = nil
	// ch <- struct{}{}
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	// ch <- struct{}{}
	// }()
	<-ch
}
