# Shell Script
## 基础
+ 第一行开头需要是如下,代表使用的shell解释器
    ```shell
    #! bin/bash
    ```
+ 执行脚本
    ```shell
    sh fileName
    #或者chmod +x fileNmae后
    ./fileName
    ```

+ 变量
    + shell变量赋值时候不能有空格符,只能用字母数字下划线
    + 使用定义过的变量要加美元符
    + 变量名外加花括号可以方便识别边界,和后面的字符就可以不加空格了
    ```shell
    name="zhang"
    echo $name jie
    name=wang #重新赋值
    echo ${name}jie
    readonly name #以后不能修改了
    unset variable_name #删除变量

+ 字符串
    + 可以使用单引号或者双引号,单引号会原样输出,双引号可以使用变量和转义字符
    + 使用双引号拼接字符串
        ```shell
        a=zhang
        b=jie
        c="$a$b"
        ```
    + 获取字符串长度
        `echo ${#c}`

+ 数组
    + 只有一维数组
    + 定义
    ```shell
    array_name=(value0 value1 value2 value3)
    #或者
    array_name=(
    value0
    value1)
    #或者
    array_name[0]=value0
    array_name[1]=value1
    ```
    + 使用
    ```shell
    valuen=${array_name[n]}

    echo ${array_name[@]}#获取所有元素
    #或者
    echo ${array_name[*]}

    length=${#array_name[@]}#获取数组长度
    #或者
    length=${#array_name[*]}
    ```
+ 获取命令行参数
    ```shell
    echo "执行的文件名:$0";

    echo "第一个参数为:$1";

    echo "参数个数:$#";

    echo "脚本运行进程ID:$$";

    echo "传递的参数作为一个字符串显示：$*";

    echo "传递的参数作为多个字符串显示：$@";
    for i in "$@"; do
    echo $i
    ```

---
#基本运算符
+ 算数运算符
    + (+ - * / %)使用expr和反引号,运算符左右要有空格
    ```shell
    val=`expr 2 + 2`
    echo "两数之和为 : $val" #输出4
    `expr $a \* $b` #乘法必须要反斜杠
    ```
    + 等于和不等,必须要有空格,必须要在中括号里(支持数字和字符串)
    ```shell
    [ $a == $b ]
    [ $a != $b ]
    ```
+ 关系运算符
    + 只支持数字,或者是数字的字符串
    ```shell
    [ $a -eq $b ]#检测是否相等
    [ $a -ne $b ]#不等
    [ $a -gt $b ]#大于
    [ $a -lt $b ]
    [ $a -ge $b ]#大于等于
    [ $a -le $b ]

    [ ! false ]#取反
    [ $a -lt 20 -o $b -gt 100 ]#-o or
    [ $a -lt 20 -a $b -gt 100 ]#-a and
    ```

+ 逻辑运算符
    ```shell
    [[ $a -lt 100 && $b -gt 100 ]] 
    [[ $a -lt 100 || $b -gt 100 ]] 
    ```
+ 字符串运算符
    ```shell
    [ $a = $b ] #判断字符串相等
    [ -z $a ] #判断是否为0
    [ -n "$a" ] #判断是否不为0
    [ $a ] #判断字符串是否为空
    ```

+ 文件测试运算符
    ```shell
    file="/var/www/runoob/test.sh"
    [ -e $file ] #是否存在
    [ -s $file ] #是否为空,不空返回true
    [ -x $file ] #是否可执行
    [ -w $file ]#是否可写
    [ -r $file ] #是否可读
    [ -d $file ] #是否是目录
---
##重定向
+ 每个linux/unix运行会打开三个文件:标准输入stdin(0),标准输出stdout(1),标准错误输出stderr(2)
    ```shell
    command > file #将stdout重定向到 file
    command < file #将stdin 重定向到 file
    command 2>file #重定向stderr
    command > /dev/null #起到"禁止输出"的效果

---
#环境变量
+ /etc/profile文件中添加变量,对所有用户永久生效
+ ~/.bash_profile,对当前用户永久生效
+ 临时,开启子shell,重新打开后就没了
    ```shell
    export CLASSPATH=./JAVA_HOME/lib
    ```
+ 新加PATH,完成后 source fileName立即生效,否则要等到重新进入生效
    ```shell
    export PATH=$PAHT:<new path>
    ```
+ env显示所有环境变量
---
## 命令
+ date命令
    ```shell
    date "+%Y年 %m月 %d日 %H点 %M分 %S秒 星期%w"
    ```
+ echo(会自动添加换行符)
    ```shell
    echo "It is a test" > myfile #显示结果定向到文件
    echo "It is a test" >> myfile #以追加的形式
    echo -e "OK! \n" #e选项开启转义就能换行了
    echo -e "OK! \c" #\c是不换行意思
    echo `date` #显示命令执行结果
    ```
+ printf
    ```shell
    printf "%d %s\n" 1 "abc" #单引号也行,没有引号也行
    printf %s abc def #格式只指定一个参数,被重用
    printf "\a" #响铃

+ SCP(-r 文件夹)
    ```shell
    scp fileName zhangjie@gpu8.fabu.ai:/home/zhangjie
    ```
+ find
    ```shell
    find /path -name "name"
    ```
+ grep
    ```shell
    grep pattern fileName
    grep "pattern" /path -r
    ls -l | grep staff
    ```
+ whoami 显示用户
