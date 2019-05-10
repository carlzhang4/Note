# Go

### 基本
+ go run hello.go 运行
+ go build hello.go 生成二进制文件
+ 每个go程序都要有一个main包
+ "fmt"包包含了io函数
+ 标识符大写才能被包外引用,小写在包内全局可见
+ 字符串可以通过➕连接
+ len(a)得到长度
+ if条件不用括号
+ for有三种,都不要括号
  + for init;condition;post{}
  + for condition{}
  + for {}
    ```go
    numbers := [6]int{1,2,3,5}
	for i,x:=range numbers{
		fmt.Println(i,x)
    }
    ```
+ 强制类型转换 int(variable)
+ 结构体,作为参数是值传递
    ```go
    type Books struct{
		title string
		id int
    }
    新建结构体
    fmt.Println(Books{"oop",1})
	fmt.Println(Books{title:"oop",id:1})
    fmt.Println(Books{title:"oop"})
    ```
+ range slice会返回对应的index和value
+ range string会返回index和字符的unicode
+ range map 返回键值对


---
### 指针
+ 指针 var ip *int,空指针的值为nil
+ 指针数组
    ```go
    var ptr [3]*int
	a := [3]int{1,2,3}
	for i:=0;i<3;i++{
	    ptr[i] = &a[i]
	}
	for i:=0;i<3;i++{
	    fmt.Println(*ptr[i])
    }
    ```
+ 指针指针
    ```go
    a := 100
	var ptr *int
	var pptr **int
	ptr = &a
	pptr = &ptr
    fmt.Println(**pptr)
    ```
---
### 变量
+ 变量类型
    + 布尔类型:
    ```go
    var b bool = true
    ```
    + 数字类型,int,float64
    ```go
    var b, c int = 1, 2
    ```
    + 字符串
+ 变量只声明不初始化就为默认值,数值类型为0,bool为false,字符串为空,其他为nil
+ 变量声明
    ```go
    var a int = 12

    var a = 12

    var a int
    a = 12

    a:=12

    var a int
    a,b := 1,2 // 使用这种语法必须有新变量被声明带旧变量
    ```
+ 多变量声明
    ```go
    var a,b,c type
    a,b,c = va,vb,vc

    var a,b = va,vb //自动推断

    a,b := va,vb

    var(
        a type_a
        b type_b
    )//一般用于声明全局变量
    ```
+ &a来取地址
+ 交换同类型变量值 a,b=b,a
+ _用来抛弃值,只写变量 _,a = 1,2
+ 常量用const声明
+ const枚举
    ```go
    const(
        a = 0
        b = 1
    )
    ```
---
### 数组
+ 声明 
    ```go
    var name [size]type
    var balance = [5]float64{1,2,3}
    var balance = []float64{1,2,3}

    a:=[3][4]int{       //d多维数组
        {1,2},
        {2,3},
    }

    a:=[3][4]int{       //d多维数组
        {1,2},
        {2,3}}
    ```

---
### 切片
+ 切片
    ```go
    slice := make([]int,3,10) //声明 make([]T,length,capacity) capacity可选
    slice := []int{1,2,3}
    slice := arr[:]
    len()//长度
    cap()//容量
    append(slice,1,2,3)//添加三个元素
    copy(slice1,slice2)//拷贝2中内容到1
---
### Map
+ Map
    ```go
    var Map map[string]string //只声明的话为nil,不能存放键值对
    Map = make(map[string]string)

    Map := map[string]string{key:value,...}

    value,ok := Map[key] //ok是bool变量,可以用来判断是否存在

    delete(Map,key)删除某个条目
    ```
---
### 并发
+ 开启goroutine go funcion()
+ 通道
    定义通道变量 c := make(chan int)
    设置缓冲区 c:= make(chan int,100),不带缓冲区发送方会阻塞直到接收方从通道接收值,接收方在有值前一直阻塞
    v,ok <-c 没值ok为false
    可以用close关闭 close(c)
    ```go
    func sum(s []int, c chan int){
        sum := 0
        for _,v :=range s{
            sum+=v
        }
        c <- sum
    }
    func main() {
        s := []int{1,2,3,4}
        c := make(chan int)
        go sum(s[:len(s)/2],c)
        go sum(s[len(s)/2:],c)
        x,y := <-c,<-c
        fmt.Println(x,y)
    }
---
### 函数
+ func(参数)返回类型{}
    ```go
    func test(a int,b float64){
	    fmt.Println(a,b)
    }
    ```
+ go语言默认值传递
+ 方法
    ```go
    type Circle struct{
	    radius float64
    }

    func main() {
        var c1 Circle
        c1.radius = 10.00
        fmt.Println(c1.getArea())

    }
    func (c Circle)getArea()float64{
        return 3.14 * c.radius * c.radius
    }
    ```
+ 闭包
    ```go
    func getSequence() func(string)int{
        i:=0
        return func(str string)int{
            i+=1
            return i
        }
    }
    func main(){
        nextNumber := getSequence()
        fmt.Println(nextNumber("a"))
        fmt.Println(nextNumber("b"))
    }
    ```
+ 接口
    ```go
    type Phone interface{
	    call()
    }

    type Iphone struct{

    }

    func (iPhone Iphone)call(){
        fmt.Println("I am Iphone")
    }
    func main() {
        var phone Phone
        phone = new(Iphone)
        phone.call()
    }
    ```
