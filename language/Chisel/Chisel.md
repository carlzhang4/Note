# Chisel
### 基础
+ test:runMain examples.Launcher GCD 开启测试
+ test:runMain examples.Launcher GCD --backend-name verilator 生成verilog
+ Input可以悬空 默认为0, Output不可悬空

### 语法
+ 插值Print
    ```scala
    val myUInt = 33.U
    printf(p"myUInt = $myUInt") // myUInt = 33

    // Hexadecimal
    printf(p"myUInt = 0x${Hexadecimal(myUInt)}") //myUInt = 0x21

    // Binary
    printf(p"myUInt = ${Binary(myUInt)}") // myUInt = 100001

    // Character
    printf(p"myUInt = ${Character(myUInt)}") // myUInt = !
    ```

+ 定义寄存器变量
    ```scala
    val x = RegInit(0.U(max.getWidth.W))
	val y = RegInit(4.U(255.W))  //注意 括号里只能用 数字.W 的形式,直接用数字会有很大问题
    ```

+ 定义wire
    ```scala
    val wire = Wire(UInt(15.W))
    ```
+ 定义普通数据
    ```scala
    val t = UInt(4.W)
    ```
+ 定义IO的变量
    ```scala
    val wire = Output(Bool())
    val tot = Output(UInt(8.W))
    ```
+ wrap数据
    ```scala
    def wrapAround(n: UInt, max: UInt) =
		Mux(n > max, 0.U, n)
    ```
+ 单例模式
    ```scala
    //原来的代码
    package solutions
    import chisel3._
    class Adder extends Module {
        def wrapAround(n: UInt, max: UInt) =
            Mux(n > max, 0.U, n)

        def counter(max: UInt, en: Bool, amt: UInt): UInt = {
            val x = RegInit(0.U(max.getWidth.W))
            when (en) {
                x := x + amt
                }
            return x
        }
        val io = IO(new Bundle {
            val inc = Input(Bool())
            val amt = Input(UInt(4.W))
            val tot = Output(UInt(8.W))
        })
        io.tot := counter(255.U, io.inc, io.amt)
        printf(p"$io\n")
    }

    //修改为单例模式
    object Adder {
        def wrapAround(n: UInt, max: UInt) =
            Mux(n > max, 0.U, n)

        def counter(max: UInt, en: Bool, amt: UInt): UInt = {
            val x = RegInit(0.U(max.getWidth.W))
            when (en) {
                x := x + amt
            }
            return x
        }
    }
    class Adder extends Module {
        val io = IO(new Bundle {
            val inc = Input(Bool())
            val amt = Input(UInt(4.W))
            val tot = Output(UInt(8.W))
        })
        io.tot := Adder.counter(255.U, io.inc, io.amt)
        printf(p"$io\n")
    }
    ```


