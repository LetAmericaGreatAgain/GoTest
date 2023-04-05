package main

import(
 "fmt"
 "time"
 "strconv"
) 

func main(){
	/**
	 * go并发
	 */
	//可以使用go开启一个新的运行期线程
	go say("first")
	say("second")

	/**
	 * 通道channel是用来传递数据的一个数据结构
	 */
	c := make(chan int)
	arr := []int{1,3,4,2,-9,3,5,10,9}
	go sum(arr[:len(arr)/2],c)
	go sum(arr[len(arr)/2:],c)
	x , y := <-c , <-c
	fmt.Printf("x=%d,y=%d,sum=%d\n",x,y,x+y)

	//带缓存的通道
	ch := make(chan int,2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	//fmt.Println(<-ch) 
	v,ok := <- ch
	if(ok){
		fmt.Println("通道中还有值，值为：",v)
	}else{
		fmt.Println("通道中已经没有数据了，即将手动关闭通道")
		close(ch)//fatal error: all goroutines are asleep - deadlock!
	}

    mychannel := make(chan int,3)
    go channelTravel(arr,mychannel)
    fmt.Println("接收到的通道中的数据为:")
    for i := range mychannel{
    	fmt.Println(i)
    }
}

func sum(numbers []int,c chan int){
	sum := 0
	for _,v := range numbers{
		sum += v
	}
	c <- sum
}

func say(word string){
	for i := 0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(word+strconv.Itoa(i))
	}
}

func channelTravel(arr []int,ch chan int){
	for _,value := range arr{
		ch <- value
	}
	close(ch)//如果close不关闭，主函数会在接收完arr中数据后还一直等待接受，主函数会一直阻塞
}