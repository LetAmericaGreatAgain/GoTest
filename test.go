package main

import "fmt"
import (
   "time"
)


var global_variable int 

 func main(){
    fmt.Println("hello , world!")
    fmt.Println("google"+"good")
    var stockcode = 123
    var enddate = "2020-12-31"
    var url = "Code = %d & endDate = %s"
    var target_url = fmt.Sprintf(url,stockcode,enddate)
    fmt.Println(target_url)
    /**
     * go语言有一下几种数据类型
     * 布尔型 bool
     * 数字类型 float32 float64 int
     * 字符串类型
     * 派生类型:1.指针类型Pointer 2.数组类型 3.结构化类型struct 4.Channel类型 5.函数类型 6.切片类型 7.接口类型interface 8.Map类型
     */
     var s string = "Runnable"
     fmt.Println(s)
     var a , b int = 1 , 2
     fmt.Println(a,b)

     ss := "run again"
     fmt.Println(ss)
     c , d := 3 , 4
     fmt.Println(c,d)

     // variable declared and not used
     //var variable int


     /**
      * 常量测试
      */
     const LENGTH int = 10
     const WIDTH int = 5
     fmt.Println(LENGTH*WIDTH)

     if true {
          fmt.Println("the condition is true")
     }

     /**
      * switch测试
      */
     var x interface{}
     switch i := x.(type){
         case nil:
            fmt.Printf("x的类型 :%T",i)
         case int:
            fmt.Printf("x的类型 :%T",i)
         default:
            fmt.Printf("未知型")
     }

     switch  {
     case true:
        /* code */
        fmt.Println()
        fmt.Println("first is true")
        fallthrough
     case true:
         fmt.Println("second are forced")
     default:
        /* code */
        fmt.Println("this is default case")
     }

     /**
      * for测试
      */
     var i int = 0;
     for true {//相当于while
         if i>=10 {break;}
         fmt.Printf("this is %d\n",i);
         i++;
     }

     for i:=0;i<3;i++{
         fmt.Printf("for loop %d\n",i)
     }

     /**
      * channel与select测试
      */
     c1 := make(chan string)
     c2 := make(chan string)

     go func(){
        time.Sleep(1*time.Second)
        c1 <- "one"
     }()
     go func(){
        time.Sleep(1*time.Second)
        c2 <- "two"
     }()

     for i:=0;i<2;i++{
         select{
         case msg1 := <-c1:
            fmt.Printf("%d received %s\n",i,msg1)
         case msg2 := <-c2:
            fmt.Printf("%d received %s\n",i,msg2)
         }
     }

     /**
      * 函数使用
      */
      fmt.Println(max(100,13849))
      fmt.Println(swap("google","apple"))
      s1 , s2 := swap("蟑螂","可怕的");
      fmt.Println(s1+s2)

      /**数组使用*/
      var arr [10]int;
      //var arr [10] int;
      for i := 0;i<10;i++{
         arr[i] = i;
      }
      for i := 0;i<10;i++{
         fmt.Printf("%d -- %d\n",i,arr[i])
      }

      /**
       * go语言指针
       */
      var iii int = 10;
      fmt.Printf("iii变量的地址为：%x\n",&iii)

      var p *int//p==nil
      fmt.Printf("p的值为%d\n",p)
      if p==nil{
         fmt.Println("p是一个空指针")
      }
      if(p!=nil){
         fmt.Println("p不是一个空指针")
      }

      /**
       * go语言结构体
       */
      var mybook Book
      mybook.title = "hello word"
      mybook.author = "www.kang.com"
      mybook.subject = "尘埃与世界"
      mybook.book_id = 1
      PrintBook(&mybook)

      
}

func PrintBook(book *Book){
   fmt.Printf("book.title=%s\n",book.title)
   fmt.Printf("book.author=%s\n",book.author)
   fmt.Printf("book.subject=%s\n",book.subject)
   fmt.Printf("book.book_id=%d\n",book.book_id)
}

type Book struct{
   title string
   author string
   subject string
   book_id int
}

func max(a int,b int) int{
   var c int;
   if(a>b){
      c = a;
   }else{
      c = b;
   }
   return c;
}

/**
 * go函数可以返回多个值
 */
func swap(a ,b string)(string , string){
   return b,a;
} 