package main

import "fmt"

func main() {
	intChan:=make(chan int, 10)
	for i:=0;i< 10;i++{
		intChan<-i
	}
	close(intChan)
	syncChan:=make(chan struct{}, 1)
	go func() {
		/**
		代码中的Loop:指明了其下方的for语句才是真正跳出语句。如果没有Loop，那么只会从select中break出来，程序就死循环了。
		 */
	Loop:
		for {
			select {
			case e, ok:=<-intChan:
				if !ok{
					fmt.Println("End.")
					break Loop
				}
				fmt.Printf("Received: %v\n", e)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
}
