package main

import "fmt"

var userId int
var userPwd string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	var loop = true
	for {
		fmt.Println("-----欢迎登陆多人聊天系统-----")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}

		// 更多用户的输入，显示新的提示信息
		if key == 1 {
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n", &userPwd)
			// 登录逻辑处理
			err := login(userId, userPwd)
			if err != nil {
				fmt.Println("登录失败\nerr:", err)
				continue
			}
			fmt.Println("登录成功", loop)

		} else if key == 2 {

		} else {

		}

	}
}
