package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func uniq(strings []string) (ret []string) {
	return
}

func authors() []string {
	if b, err := exec.Command("git", "log").Output(); err != nil {
		lines := strings.Split(string(b), "\n")

		var a []string
		r := regexp.MustCompile(`^Author:\s*([^<]+).*$`)
		for _, e := range lines {
			ms := r.FindStringSubmatch(e)
			if ms == nil {
				continue
			}

			a = append(a, ms[1])
		}

		sort.Strings(a)
		var p string
		lines = []string{}
		for _, e := range a {
			if p == e {
				continue
			}
			lines = append(lines, e)
			p = e
		}

		return lines
	}

	return []string{"Yasuhiro Matsmoto <math.jp@gmail.com>"}
}

func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK GO!")
	window.SetIconName("gtk-dialog-info")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	vbox := gtk.NewVBox(false, 1)

	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	frame1 := gtk.NewFrame("Demo")
	framebox1 := gtk.NewVBox(false, 1)
	frame1.Add(framebox1)

	frame2 := gtk.NewFrame("Demo")
	framebox2 := gtk.NewVBox(false, 1)
	frame2.Add(framebox2)

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	dir, _ := filepath.Split(os.Args[0])
	imageFile := filepath.Join(dir, "../../data/test1.jpeg")

	label := gtk.NewLabel("Go Binding for GTK")
	label.ModifyFontEasy("Test")
	framebox1.PackStart(label, false, true, 0)

	entry := gtk.NewEntry()
	entry.SetText("Hello World")
	framebox1.Add(entry)

	image := gtk.NewImageFromFile(imageFile)
	framebox1.Add(image)

	scale := gtk.NewHScaleWithRange(0, 100, 1)
	scale.Connect("value-changed", func() {
		fmt.Println("scale:", int(scale.GetValue()))
	})
	framebox2.Add(scale)

	buttons := gtk.NewHBox(false, 1)

	button := gtk.NewButtonWithLabel("Button with Label")
	button.Clicked(func() {
		fmt.Println("button clicked:", button.GetLabel())
		messageDialog := gtk.NewMessageDialog(
			button.GetTopLevelAsWindow(),
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			entry.GetText(),
		)
		messageDialog.Response(func() {
			fmt.Println("Dialog Ok!")
			fileChooseDialog := gtk.NewFileChooserDialog(
				"Choose File ... ",
				button.GetTopLevelAsWindow(),
				gtk.FILE_CHOOSER_ACTION_OPEN,
				gtk.STOCK_OK,
				gtk.RESPONSE_ACCEPT,
			)

			filter := gtk.NewFileFilter()
			filter.AddPattern("*.go")

			fileChooseDialog.AddFilter(filter)
			fileChooseDialog.Response(func() {
				fmt.Println(fileChooseDialog.GetFilename())
				fileChooseDialog.Destroy()
			})

			fileChooseDialog.Run()
			messageDialog.Destroy()
		})

		messageDialog.Run()
	})
	buttons.Add(button)

	fontButton := gtk.NewFontButton()
	fontButton.Connect("font-set", func() {
		fmt.Println("title: ", fontButton.GetTitle())
		fmt.Println("fontName: ", fontButton.GetFontName())
		fmt.Println("userSize: ", fontButton.GetUseSize())
		fmt.Println("showSize: ", fontButton.GetShowSize())
	})
	buttons.Add(fontButton)
	framebox2.PackStart(buttons, false, false, 0)

	buttons = gtk.NewHBox(false, 1)

	toggleButton := gtk.NewToggleButtonWithLabel("Toggle with label")
	toggleButton.Connect("toggled", func() {
		if toggleButton.GetActive() {
			toggleButton.SetLabel("ToggleButton On")
		} else {
			toggleButton.SetLabel("ToggleButton Off")
		}
	})
	buttons.Add(toggleButton)

	checkButton := gtk.NewCheckButtonWithLabel("checkButton with label")
	checkButton.Connect("toggled", func() {
		if checkButton.GetActive() {
			checkButton.SetLabel("checkButton Checked")
		} else {
			checkButton.SetLabel("checkButton Unchecked")
		}
	})
	buttons.Add(checkButton)

	buttonBox := gtk.NewVBox(false, 1)
	radioFirst := gtk.NewRadioButtonWithLabel(nil, "Radio1")
	buttonBox.Add(radioFirst)
	buttonBox.Add(gtk.NewRadioButtonWithLabel(radioFirst.GetGroup(), "Radio2"))
	buttonBox.Add(gtk.NewRadioButtonWithLabel(radioFirst.GetGroup(), "Radio3"))
	radioFirst.Add(buttonBox)

	framebox2.PackStart(buttons, false, false, 0)

	vsep := gtk.NewVSeparator()
	framebox2.PackStart(vsep, false, false, 0)

	combos := gtk.NewHBox(false, 1)
	comboBoxEntry := gtk.NewComboBoxText()
	comboBoxEntry.AppendText("Money")
	comboBoxEntry.AppendText("Tiger")
	comboBoxEntry.AppendText("Elephant")
	comboBoxEntry.Connect("changed", func() {
		fmt.Println("value: ", comboBoxEntry.GetActiveText())
	})
	combos.Add(comboBoxEntry)

	comboBox := gtk.NewComboBoxText()
	comboBox.AppendText("Peach")
	comboBox.AppendText("Banana")
	comboBox.AppendText("Apple")
	comboBox.SetActive(1)
	comboBox.Connect("changed", func() {
		fmt.Println("value: ", comboBox.GetActiveText())
	})
	combos.Add(comboBox)

	framebox2.PackStart(combos, false, false, 0)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)
	textView := gtk.NewTextView()
	var start, end gtk.TextIter
	buffer := textView.GetBuffer()
	buffer.GetStartIter(&start)
	buffer.Insert(&start, "Hello ")
	buffer.GetEndIter(&end)
	buffer.Insert(&end, "world!")

	tag := buffer.CreateTag("bold", map[string]interface{}{
		"background": "#FF0000", "weight": "700",
	})
	buffer.GetStartIter(&start)
	buffer.GetEndIter(&end)
	buffer.ApplyTag(tag, &start, &end)
	swin.Add(textView)
	framebox2.Add(swin)

	buffer.Connect("changed", func() {
		fmt.Println("changed")
	})

	cascadeMenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(cascadeMenu)
	subMenu := gtk.NewMenu()
	cascadeMenu.SetSubmenu(subMenu)

	var menuItem *gtk.MenuItem
	menuItem = gtk.NewMenuItemWithMnemonic("E_xit")
	menuItem.Connect("activate", func() {
		gtk.MainQuit()
	})
	subMenu.Append(menuItem)

	cascadeMenu = gtk.NewMenuItemWithMnemonic("_View")
	menubar.Append(cascadeMenu)
	subMenu = gtk.NewMenu()
	cascadeMenu.SetSubmenu(subMenu)

	checkMenuItem := gtk.NewCheckMenuItemWithMnemonic("_Disable")
	checkMenuItem.Connect("activate", func() {
		vpaned.SetSensitive(!checkMenuItem.GetActive())
	})
	subMenu.Append(checkMenuItem)

	menuItem = gtk.NewMenuItemWithMnemonic("_Font")
	menuItem.Connect("activate", func() {
		fsd := gtk.NewFontSelectionDialog("Font")
		fsd.SetFontName(fontButton.GetFontName())
		fsd.Response(func() {
			fmt.Println(fsd.GetFontName())
			fontButton.SetFontName(fsd.GetFontName())
			fsd.Destroy()
		})
		fsd.SetTransientFor(window)
		fsd.Run()
	})
	subMenu.Append(menuItem)

	cascadeMenu = gtk.NewMenuItemWithMnemonic("_Help")
	menubar.Append(cascadeMenu)
	subMenu = gtk.NewMenu()
	cascadeMenu.SetSubmenu(subMenu)

	menuItem = gtk.NewMenuItemWithMnemonic("_About")
	menuItem.Connect("activate", func() {
		dialog := gtk.NewAboutDialog()
		dialog.SetName("Go-Gtk Demo!")
		dialog.SetProgramName("demo")
		dialog.SetAuthors(authors())
		dir, _ := filepath.Split(os.Args[0])
		imageFile := filepath.Join(dir, "../../data/test1.jpeg")
		pixBuf, _ := gdkpixbuf.NewPixbufFromFile(imageFile)
		dialog.SetLogo(pixBuf)
		dialog.SetLicense("he klsjdkjfdskjlsfdfsdlkjsfdkjs")
		dialog.SetWrapLicense(true)
		dialog.Run()
		dialog.Destroy()
	})
	subMenu.Append(menuItem)

	statusBar := gtk.NewStatusbar()
	context_id := statusBar.GetContextId("go-gtk")
	statusBar.Push(context_id, "GTK binding for Go")

	framebox2.PackStart(statusBar, false, false, 0)

	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
