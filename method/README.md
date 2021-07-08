Golang 方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)

- 只能为当前包内命名类型定义方法
- 参数 `receiver` 可任意命名。如方法中未曾使用 ，可省略参数名
- 参数 `receiver` 类型可以是 `T` 或 `*T`。基类型 `T` 不能是接口或指针。
- 不支持方法重载，`receiver` 只是参数签名的组成部分
- 可用实例 `value` 或 `pointer` 调用全部方法，编译器自动转换

一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。

## 普通函数与方法的区别
1. 对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
2. 对于方法（如 `struct` 的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以