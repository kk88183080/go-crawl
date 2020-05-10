package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"log"
	"os"
)

const Add string = "添加"
const Del string = "删除"
const Exit string = "退出"

var win *gtk.Window

func main() {
	// 初始化环境
	gtk.Init(&os.Args)

	/**
	主窗口
	*/
	win = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)

	configView(win)

	// 创建固定布局
	layout := gtk.NewFixed()
	// 添加布局
	win.Add(layout)
	//显示所有的控件
	win.ShowAll()

	// 添加按钮
	addButton(layout, "第一个", 400, 400, Add)
	addButton(layout, "第二个", 200, 200, Del)
	addButton(layout, "第三个", 100, 100, Exit)

	gtk.Main()
}

func configView(win *gtk.Window) {
	// 设置窗口显示的位置，默认的
	win.SetPosition(gtk.WIN_POS_NONE)
	// 设置窗口显示的位置，居中
	win.SetPosition(gtk.WIN_POS_CENTER)
	//win.SetPosition(gtk.WIN_POS_MOUSE)
	//win.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)
	//win.SetIconFromFile("face.png")     //设置icon
	// 设置标题
	win.SetTitle("第一个窗口")
	// 设置默认大小
	win.SetDefaultSize(500, 500)
	// 设置大小
	//win.SetSizeRequest(300, 500)
	// 设置不可大小伸缩
	//win.SetResizable(false)
	//win.SetDeletable(false) // 取消关闭按钮,不起作用
	// 按窗口关闭按钮，自动触发"destroy"信号
	win.Connect("destroy", func() {
		log.Println("点击了窗口关闭按钮")
		gtk.MainQuit()
	})

	win.Connect("configure_event", func() {
		width, height := win.GetSize()
		log.Println("width:%s, heigth:%s", width, height)
	})
}

func addButton(layout *gtk.Fixed, lable string, width int, height int, btn interface{}) {
	button := gtk.NewButtonWithLabel(lable)
	layout.Put(button, width, height)
	//layout.Move(button, width, 200)
	button.Show()
	// 名字
	button.SetName(btn.(string))
	button.SetTooltipText("哈哈")

	//按下按钮时触发
	button.Connect("pressed", buttonHandler, button.GetName(), 10)
	//按下按钮时触发
	button.Connect("clicked", buttonHandler, button.GetName(), 10)
	//释放按钮时触发
	button.Connect("released", buttonHandler, button.GetName(), 10)
}

func buttonHandler(ctx *glib.CallbackContext) {
	args := ctx.Data()
	s := args.(string)
	if s == "添加" {
		log.Println("我点击了添加按钮")
	} else if s == "删除" {
		log.Println("我点击了删除按钮")
	} else if s == "退出" {
		log.Println("我点击了退出按钮")
		// 整个窗口不可见
		//win.SetVisible(false)
		// 延时退出时间
		//time.Sleep(time.Millisecond * 1500)

		//gtk.MainQuit()
		gtk.MainQuit()
		os.Exit(0)
	}
}
