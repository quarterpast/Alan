package main

import (
	"github.com/quarterto/go-gtk/gtk"
	"os"
	"fmt"
)

type Alan struct {
	*gtk.GtkWindow
	tabs *gtk.GtkNotebook
	toolbar *gtk.GtkHBox
}

func TabToolbar() *gtk.GtkHBox {
	hbox := gtk.HBox(true,0)
	return hbox
}
func _Alan() Alan {
	alan := Alan{}

	alan.GtkWindow = gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	alan.tabs = gtk.Notebook()
	alan.toolbar = TabToolbar()

	alan.tabs.SetShowTabs(false)
	alan.tabs.Connect("change-current-page",func(pos int) {
		tab := alan.tabs.GetNthPage(pos)
		alan.SetTitle(alan.tabs.GetMenuLabelText(tab))
	})
	alan.tabs.Connect("page-removed",func() {
		if alan.tabs.GetNPages() == 1 {
			alan.tabs.SetShowTabs(false)
		}
	})
	alan.tabs.Connect("page-added",func() {
		if alan.tabs.GetNPages() > 1 {
			alan.tabs.SetShowTabs(true)
		}
	})
	alan.tabs.SetActionWidget(alan.toolbar)
	alan.SetDefaultSize(400,500)
	alan.Connect("destroy",gtk.MainQuit)
	alan.Add(alan.tabs)
	return alan
}
func (a Alan) OpenDialog() {
	dialog := gtk.FileChooserDialog("Open",a.GtkWindow,gtk.GTK_FILE_CHOOSER_ACTION_OPEN,"Open",1)
	dialog.ShowAll()
}
func (a Alan) Open(files []string) {
	for _,f := range files {
		fmt.Printf("%d",a.tabs.AppendPage(gtk.Label(f),gtk.Label(f)));
	}
}
func main() {
	gtk.Init(&os.Args)
	defer gtk.Main()

	app := _Alan()
	if(len(os.Args) > 1) {
		app.Open(os.Args[1:])
	}
	app.ShowAll()
}