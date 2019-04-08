# 基础
## 语法
```scala
import java.awt._// 引入包内所有成员
```

+ 如果浮点数后面有f或者F后缀时，表示这是一个Float类型，否则就是一个
Double类型的。
+ 声明元组 val pa = (40,"Foo")
```scala
//var定义变量  val定义常量  可以显示定义类型:
val a:Int = 123;
```

```scala
//for循环:
for( a <- 1 to 10)有10   for( a <- 1 until 10) //没10

for( a <- 1 to 3; b <- 1 to 3)//遍历所有值

for( var x <- List )

for(a <- 1 to 10 if a>5;if a<8)
```

```scala
//yield操作: 将for循环后的值生成变量
var returnValue=
    for(a <- 1 to 10 if a>5;if a<8)
        yield a

    for(a<-returnValue){
        println(a)
    }
```

```scala
//定义方法 不要返回值时用Unit就行,不写等号和定义,方法会被隐式声明为抽象
def _print(a:Int,b:String):Unit={
        println(a)
        println(b)
    }
```
+ 函数用var声明,方法用def
+ 闭包是一个函数，返回值依赖于声明在函数外部的一个或多个变量。闭包通常来讲可以简单的认为是可以访问一个函数里面局部变量的另外一个函数。
```scala
//例如：
var factor = 2
val multiplier = (i:Int)=>i*factor
println(multiplier(3))
```

```scala
//StringBuilder类
val buf = new StringBuilder;
buf += 'a'
buf ++= "bcdef"
println( "buf is : " + buf.toString );
```
```scala
//定义数组
var z = Array("Runoob", "Baidu", "Google")
//或
var a:Array[String] = new Array[String](3)
a(2) = "zj"

//合并数组
var myList3 =  concat( myList1, myList2)

//创建区间
var myList1 = range(10, 20, 2)
//或
var myList1 = range(10, 20) //默认步长为1

//创建二维数组
var index = new Array[Array[Int]](10)
for(i <- 0 until index.length){
    index(i) = new Array[Int](4)
}
index(3)(2) = 23;
```

```scala
//Scala集合

// 定义整型 List
val x = List(1,2,3,4)

// 定义 Set
val x = Set(1,3,5,7)

// 定义 Map
val x = Map("one" -> 1, "two" -> 2, "three" -> 3)

// 创建两个不同类型元素的元组
val x = (10, "Runoob")

// 定义 Option
val x:Option[Int] = Some(5)
```

可选值,List和Map等的返回类型:  
+ 有值时用some包裹，没有时返回None
+ some和None都是option的子类

```scala
//强制拆包
println(x.get("one") get)

//getOrElse() 方法:返回存在的元素或者输入值

//模式匹配
val sites = Map("runoob" -> "www.runoob.com", "google" -> "www.google.com")
println(show(sites.get( "runoob")) )
println(show(sites.get( "baidu")) )

def show(x: Option[String]) = x match {
      case Some(s) => s
      case None => "?"
   }
/**输出：
www.runoob.com
?
**/
```

```scala
//迭代器基本操作
val it = Iterator("Baidu", "Google", "Runoob", "Taobao")
while (it.hasNext){
    println(it.next())
}

// 迭代器
val ita = Iterator(20,40,2,50,69, 90)
println("最大元素是：" + ita.max )
println("最小元素是：" + ita.min )
println("ita.size 的值: " + ita.size )//size和length一样
println("ita.length 的值: " + ita.length )
```

```scala
// 模式匹配,匹配了一个就不会再匹配其他的了
def matchTest(x: Int): String = x match {
      case 1 => "one"
      case 2 => "two"
      case _ => "many"//其余所有的
   }
println(matchTest(3))

//不同数据类型的模式匹配
def matchTest(x: Any): Any = x match {
      case 1 => "one"
      case "two" => 2
      case y: Int => "scala.Int"
      case _ => "many"
   }
```

```scala
//提取器
def unapply(str: String): Option[(String, String)] = {
      val parts = str split "@"
      if (parts.length == 2){
         Some(parts(0), parts(1)) 
      }else{
         None
      }
   }
println ("Unapply 方法 : " + unapply("Zara@gmail.com"));
```

