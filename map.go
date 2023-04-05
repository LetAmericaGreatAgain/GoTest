package main

import "fmt"

func main(){
	var siteMap map[string]string

	siteMap = make(map[string]string)

	siteMap["google"] = "谷歌"
	siteMap["apple"] = "苹果"
	siteMap["unknown"] = "unknown"

	for key,value := range siteMap{
		fmt.Printf("key=%s,value=%s\n",key,value)
	}

	fmt.Println(siteMap["google"])

	value,isExist := siteMap["YouTube"]
	// if(isExist == nil){//(mismatched types bool and untyped nil)
	// 	fmt.Println("value is not exist")
	// }
	if(isExist){
		fmt.Printf("value exist and is %s\n",value)
	}else{
		fmt.Println("value is not exist")
	}

    delete(siteMap,"unknown")
    fmt.Println("unknown 已经被删除")
    for key,value := range siteMap{
		fmt.Printf("key=%s,value=%s\n",key,value)
	}

	fmt.Printf("通过递归求1-n之和，和为：%d\n",SumByRecursion(100))
}

func SumByRecursion(n int)int{
	if(n == 0){
		return 0
	}else{
		return n+SumByRecursion(n-1)
	}

}

/**
key=unknown,value=unknown
key=google,value=谷歌
key=apple,value=苹果
谷歌
value is not exist
unknown 已经被删除
key=google,value=谷歌
key=apple,value=苹果
通过递归求1-n之和，和为：5050
*/
