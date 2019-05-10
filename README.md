# fizzbuzz
main.go是主程序，提示输入两个10以内的整数。这里没有判断非法。

fizzlib目录是具体逻辑，带TDD 测试case

下面是case的设计记录：
https://docs.qq.com/sheet/DZlBOSGtjSGd1ZEp0?tab=BB08J2&coord=A25%24A25%240%240%241%240


文件说明：
main.go 为FizzBuzz程序入口，提示输入两个10以内的整数，其中第一个特殊数字设置为3，第二个特殊数字为5；
fizzlib目录是具体逻辑：Fizzgame类函数实现返回数的判断

执行方法:
1、go mian.go，按照提示输出两个小于10的整数且用空格分开，如3 5，回车。会打印1-100的输出结果。
2、运行go test -v可以运行所有测试，并看到测试结果。
