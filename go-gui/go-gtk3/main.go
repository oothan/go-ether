package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func main() {
	// initialize GTK without parsing any command line argument
	gtk.Init(nil)

	// create a new toplevel window, set its title, and collect it to the "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window: ", err)
	}

	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new label widget to show in the window
	l, err := gtk.LabelNew("Hello, GTK3")
	if err != nil {
		log.Fatal("Unable to create label : ", err)
	}

	// add label to the window
	win.Add(l)

	// set the window size
	win.SetDefaultSize(800, 600)

	// recursively show all widgets contained in this window
	win.ShowAll()

	// begin executing the GTK main loop.
	// This blocks until gtk.MainQuit() is run.
	gtk.Main()
}
