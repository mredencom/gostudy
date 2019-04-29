# gostudy# 2019-01-01 致力于go语言学习Personal study#### 最近几天没有学习#### 过年这几天已经忘记学习了~_~### go语言的流程控制:    1. 没有do和while循环,只有更广义的for循环    2. switch语句灵活多变,还可以用于类型判断    3. if语句和switch语句都可以包含一天初始化子语句    4. break语句和continue语句可以后跟一条标签(label)语句,以标识需要终止或继续的代码块    5. defer语句可以使用我们的更加方便的执行异常的捕获和资源回收任务    6. select语句也用于多分支的选择,但只与通道配合使用    7. go语句用于异步启用goroutine并执行指定函数### 小结 v1.0    1. 每个代码文件都属于一个包，而包名应该与代码文件所在的文件夹同名。    2. Go 语言提供了多种声明和初始化变量的方式。如果变量的值没有显式初始化，编译器会      将变量初始化为零值。    3. 使用指针可以在函数间或者 goroutine 间共享数据。    4. 通过启动 goroutine 和使用通道完成并发和同步。    5. Go 语言提供了内置函数来支持 Go 语言内部的数据结构。    6. 标准库包含很多包，能做很多很有用的事情。    7. 使用 Go 接口可以编写通用的代码和框架。### 小结 v2.0    1.在 Go 语言中包是组织代码的基本单位。          2.环境变量 GOPATH 决定了 Go 源代码在磁盘上被保存、编译和安装的位置。       3.可以为每个工程设置不同的 GOPATH，以保持源代码和依赖的隔离。       4.go 工具是在命令行上工作的最好工具。       5.开发人员可以使用 go get 来获取别人的包并将其安装到自己的 GOPATH 指定的目录。       6.想要为别人创建包很简单，只要把源代码放到公用代码库，并遵守一些简单规则就可以了。       7.Go 语言在设计时将分享代码作为语言的核心特性和驱动力。       8.推荐使用依赖管理工具来管理依赖。       9.有很多社区开发的依赖管理工具，如 godep、 vender 和 gb### 小结 v3.0    1.数组是构造切片和映射的基石。    2.Go 语言里切片经常用来处理数据的集合，映射用来处理具有键值对结构的数据。    3.内置函数 make 可以创建切片和映射，并指定原始的长度和容量。也可以直接使用切片和映射字面量，或者使用字面量作为变量的初始值。    4.切片有容量限制，不过可以使用内置的 append 函数扩展容量。    5.映射的增长没有容量或者任何限制。    6.内置函数 len 可以用来获取切片或者映射的长度。    7.内置函数 cap 只能用于切片。    8.通过组合，可以创建多维数组和多维切片。也可以使用切片或者其他映射作为映射的值。但是切片不能用作映射的键。    9.将切片或者映射传递给函数成本很小，并且不会复制底层的数据结构。### 小结 v3.1    *和&的区别 :        & 是取地址符号 , 即取得某个变量的地址 , 如 ; &a        *是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .### 小结 v4.0    1.使用关键字 struct 或者通过指定已经存在的类型，可以声明用户定义的类型。    2.方法提供了一种给用户定义的类型增加行为的方式。    3.设计类型时需要确认类型的本质是原始的，还是非原始的。    4.接口是声明了一组行为并支持多态的类型。    5.嵌入类型提供了扩展类型的能力，而无需使用继承。    6.标识符要么是从包里公开的，要么是在包里未公开的。### 小结v5.0    1.并发是指 goroutine 运行的时候是相互独立的。    2.使用关键字 go 创建 goroutine 来运行函数。    3.goroutine 在逻辑处理器上执行，而逻辑处理器具有独立的系统线程和运行队列。    4.竞争状态是指两个或者多个 goroutine 试图访问同一个资源。    5.原子函数和互斥锁提供了一种防止出现竞争状态的办法。    6.通道提供了一种在两个 goroutine 之间共享数据的简单方法。    7.无缓冲的通道保证同时交换数据，而有缓冲的通道不做这种保证                            