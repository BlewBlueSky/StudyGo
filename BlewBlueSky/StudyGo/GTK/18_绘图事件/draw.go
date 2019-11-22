package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()       //新建builder
	builder.AddFromFile("test.glade") //读取glade文件

	window := gtk.WindowFromObject(builder.GetObject("window1")) //获取窗口控件
	window.SetSizeRequest(480, 320)                              //设置窗口大小
	window.SetAppPaintable(true)                                 //允许窗口能绘图(重要)

	var w, h int //保存窗口的宽度和高度

	//改变窗口大小时，触发"configure-event"，然后手动刷新绘图区域，否则图片会重叠
	window.Connect("configure-event", func() {
		window.QueueDraw() //刷新绘图区域

		//获取窗口的宽度和高度
		window.GetSize(&w, &h)
		fmt.Println(w, h)
	})

	x := 0 //画笑脸起点

	//绘图（曝光）事件，其回调函数做绘图操作
	window.Connect("expose-event", func() {
		//指定窗口为绘图区域，在窗口上绘图
		painter := window.GetWindow().GetDrawable()
		gc := gdk.NewGC(painter)

		//设置背景图的pixbuf，其宽高和窗口一样，最后一个参数固定为false
		bk, _ := gdkpixbuf.NewPixbufFromFileAtScale("./image/bk.jpg", w, h, false)
		//设置笑脸的pixbuf，其宽高为50，最后一个参数固定为false
		face, _ := gdkpixbuf.NewPixbufFromFileAtScale("./image/face.png", 50, 50, false)

		//画图
		//bk：需要绘图的pixbuf，第5、6个参数为画图的起点（相对于窗口而言）
		//第3、4个参数习惯为0，第7、8个参数为-1则按pixbuf原来的尺寸绘图
		//gdk.RGB_DITHER_NONE不用抖动，最后2个参数习惯写0
		painter.DrawPixbuf(gc, bk, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		painter.DrawPixbuf(gc, face, 0, 0, x, 100, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

		//释放pixbuf资源
		bk.Unref()
		face.Unref()
	})

	button := gtk.ButtonFromObject(builder.GetObject("button1")) //获取窗口控件
	//按钮"clicked"信号处理
	button.Clicked(func() {
		x += 50 //增加笑脸起点坐标

		if x >= w { //如果超过窗口宽度，置0
			x = 0
		}

		window.QueueDraw() //刷新绘图区域
	})

	//按窗口关闭按钮，自动触发"destroy"信号
	window.Connect("destroy", gtk.MainQuit)

	window.ShowAll()

	gtk.Main()
}
