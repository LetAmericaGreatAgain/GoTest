package main

import "fmt"

func main(){
	/**
	 * 切片是对数组的抽象,是一种内置类型，相当于动态数组
	 * 切片slice
	 */
	var slice1 []int //可以通过未指定数组大小的形式来声明一个切片

    // slice1[0] = 999//报错
    fmt.Printf("len=%d,capacity=%d\n",len(slice1),cap(slice1))

    if(slice1 == nil){
    	fmt.Println("slice1是一个空切片，还未进行初始化")
    }
	//也可以通过make函数来创建切片
	var slice2 []int = make([]int,7,9) //7为切片的初始长度，9为切片的容量 
    fmt.Printf("len=%d,capacity=%d\n",len(slice2),cap(slice2))

	for i := 0;i<7;i++{//i范围到9的时候会执行异常
		slice2[i] = i
	}
	for i:=0;i<7;i++{
		fmt.Printf("slice2[%d]=%d\n",i,slice2[i])
	}

    //创建切片
    numbers := []int{0,1,2,3,4,5,6,7,8,9}

    //切片截取
    fmt.Printf("len=%d,cap=%d,slice=%v\n",len(numbers),cap(numbers),numbers)
    fmt.Println("numbers[1:4] == ",numbers[1:4])
    fmt.Println("numbers[:4] == ",numbers[:4])//上限
    fmt.Println("numbers[1:] == ",numbers[1:])//下限
    
    //append和copy函数，追加和复制切片
    PrintSlice(slice1)

    slice1 = append(slice1,0)//追加一个值
    PrintSlice(slice1)
    
    slice1 = append(slice1,1,2,3)//追加多个值
    PrintSlice(slice1)

    slice3 := make([]int,len(slice1),cap(slice2)*2)
    slice4 := make([]int,len(slice1),cap(slice2)*2)

    copy(slice3,slice1)
    PrintSlice(slice3)

    copy(slice3,slice4)
    PrintSlice(slice3)//len=4,cap=18,slice=[0 0 0 0]
}
func PrintSlice(arr []int){
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(arr),cap(arr),arr)
}
/**
len=0,capacity=0
slice1是一个空切片，还未进行初始化
len=7,capacity=9
slice2[0]=0
slice2[1]=1
slice2[2]=2
slice2[3]=3
slice2[4]=4
slice2[5]=5
slice2[6]=6
len=10,cap=10,slice=[0 1 2 3 4 5 6 7 8 9]
numbers[1:4] ==  [1 2 3]
numbers[:4] ==  [0 1 2 3]
numbers[1:] ==  [1 2 3 4 5 6 7 8 9]
len=0,cap=0,slice=[]
len=1,cap=1,slice=[0]
len=4,cap=4,slice=[0 1 2 3]
len=4,cap=18,slice=[0 1 2 3]
len=4,cap=18,slice=[0 0 0 0]
*/