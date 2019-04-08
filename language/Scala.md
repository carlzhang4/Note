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
```