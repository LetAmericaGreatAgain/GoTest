package main

import "fmt"
/**
 * range关键字用于迭代遍历数组、切片、通道和集合
 */ 
func main(){
	pow := []int{1,2,4,6,8,16}

	for index,value := range pow{
		fmt.Printf("pow[%d]=%d\n",index,value)
	}

	map1 := make(map[int]float32)
	map1[1] = 1.1
	map1[2] = 2.2
	map1[3] = 3.3
	map1[4] = 4.4
	for key,value := range map1{
		fmt.Printf("map1[%d]=%f\n",key,value)
	}

    //遍历字符串，i为索引，c为unicode值
	for i,c := range "go"{
		fmt.Println(i,c)
	}
}
/**
pow[0]=1
pow[1]=2
pow[2]=4
pow[3]=6
pow[4]=8
pow[5]=16
map1[2]=2.200000
map1[3]=3.300000
map1[4]=4.400000
map1[1]=1.100000
0 103
1 111
*/