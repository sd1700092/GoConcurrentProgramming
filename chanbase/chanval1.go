package main

import (
	"fmt"
	"time"
)


// 这段代码演示了，mapChan元素类型是引用类型。因此，接收方对元素值的副本的修改会影响到发送方持有的原始值。
// 但是如果Channel里的元素是值类型，那么接收方对元素的任何修改都不会影响发送方发送的值。
var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan:=make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok:=<-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() {
		countMap:=make(map[string]int)
		for i:=0;i<5;i++{
			mapChan<-countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
