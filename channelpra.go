package main

import (
	"fmt"
	"time"
)
func sender_one(myCh chan int){
	for i:=0;i<=50;i++ {
		myCh <- i
		time.Sleep(5 * time.Second)
	}

}
func sender_two(myCh chan int){
	for i:=0;i<=50;i++ {
		myCh <- i
		time.Sleep(2 * time.Second)
	}
}
func receiver(myCh chan int,c int){
	fmt.Println(<-myCh)
	fmt.Println(c)

}
func main(){
	fmt.Println("golang Channels")
	myCh1 := make(chan int)
	myCh2 := make(chan int)
	go sender_one(myCh1)
	go sender_two(myCh2)
	select {
	case res := <-myCh1:
		receiver(myCh1,res)
	
	case res := <-myCh2:
		receiver(myCh2,res)
	}
}

	
