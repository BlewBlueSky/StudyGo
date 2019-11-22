package main

import(
	"fmt"
	"os"
	"github.com/mattn/go-gtk/gtk"
)


func main()  {
	// 1、初始化
	gtk.Init(&os.Args)

	// 2、用户代码
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("go gtk")
	win.SetSizeRequest(480,320)
	win.Show()

	// 3、主事件循环
	fmt.Println("before")
	gtk.Main()
	fmt.Println("over")
}