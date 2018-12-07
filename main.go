package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	const appID = "com.github.mernstack.weather-desktop-application"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("Could not create application:", err)
	}
	application.Connect("activate", func() {
		win := newWindow(application)

		aNew := glib.SimpleActionNew("new", nil)
		aNew.Connect("activate", func() {
			newWindow(application).ShowAll()
		})
		application.AddAction(aNew)

		aQuit := glib.SimpleActionNew("quit", nil)
		aQuit.Connect("activate", func() {
			application.Quit()
		})
		application.AddAction(aQuit)

		win.ShowAll()
	})

	os.Exit(application.Run(os.Args))
}

func newWindow(application *gtk.Application) *gtk.ApplicationWindow {
	win, err := gtk.ApplicationWindowNew(application)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle("Weather Application")



	// Create a header bar
	header, err := gtk.HeaderBarNew()
	if err != nil {
		log.Fatal("Could not create header bar:", err)
	}
	header.SetShowCloseButton(true)
	header.SetTitle("Weather App")
	header.SetSubtitle("Weather App")

	// Create a new menu button
	mbtn, err := gtk.MenuButtonNew()
	if err != nil {
		log.Fatal("Could not create menu button:", err)
	}

	// Set up the menu model for the button
	menu := glib.MenuNew()
	if menu == nil {
		log.Fatal("Could not create menu (nil)")
	}
	// Actions with the prefix 'app' reference actions on the application
	// Actions with the prefix 'win' reference actions on the current window (specific to ApplicationWindow)
	// Other prefixes can be added to widgets via InsertActionGroup
	menu.Append("New Window", "app.new")
	menu.Append("Close Window", "win.close")
	menu.Append("Custom Panic", "custom.panic")
	menu.Append("Quit", "app.quit")

	// Create the action "win.close"
	aClose := glib.SimpleActionNew("close", nil)
	aClose.Connect("activate", func() {
		win.Close()
	})
	win.AddAction(aClose)

	// Create and insert custom action group with prefix "custom"
	customActionGroup := glib.SimpleActionGroupNew()
	win.InsertActionGroup("custom", customActionGroup)

	mbtn.SetMenuModel(&menu.MenuModel)

	// add the menu button to the header
	header.PackStart(mbtn)

	// Assemble the window
	win.Add(SearchBar())
	win.Add(SearchButton())
	win.SetTitlebar(header)
	win.SetPosition(gtk.WIN_POS_MOUSE)
	win.SetDefaultSize(600, 300)
	return win
}


// A search Bar

func SearchBar() *gtk.Entry {
	// create a search bar
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry:", err)
	}

	entry.SetPlaceholderText("Enter the Location")
	return entry
}

func SearchButton() *gtk.Button{
	buttonHello, err := gtk.ButtonNewWithLabel("Search")
	if err != nil {
		log.Fatal("Unable to create Button:", err)
	}
	return buttonHello
}