package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()       //新建builder
	builder.AddFromFile("test.glade") //读取glade文件

	// 获取窗口控件指针，注意"window1"要和glade里的标志名称匹配
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	b1 := gtk.ButtonFromObject(builder.GetObject("button_previous")) //获取按钮1
	b2 := gtk.ButtonFromObject(builder.GetObject("button_next"))     //获取按钮2

	//信号处理
	b1.Connect("clicked", func() {
		//获取按钮内容
		fmt.Println("button txt = ", b1.GetLabel())
	})

	b2.Connect("clicked", func() {
		//获取按钮内容
		fmt.Println("button txt = ", b2.GetLabel())
		gtk.MainQuit() //关闭窗口
	})

	//按窗口关闭按钮，自动触发"destroy"信号
	window.Connect("destroy", gtk.MainQuit)

	window.ShowAll()

	gtk.Main()
}
