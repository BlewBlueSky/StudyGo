package main

import (
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()       //新建builder
	builder.AddFromFile("test.glade") //读取glade文件

	// 获取窗口控件指针，注意"window1"要和glade里的标志名称匹配
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.SetSizeRequest(480, 320) //设置窗口大小

	//获取布局控件
	table := gtk.TableFromObject(builder.GetObject("table1"))

	button := gtk.NewButtonWithLabel("新按钮")  //新建按钮
	table.AttachDefaults(button, 2, 3, 2, 3) //指定位置添加控件

	//按窗口关闭按钮，自动触发"destroy"信号
	window.Connect("destroy", gtk.MainQuit)

	window.ShowAll()

	gtk.Main()
}
