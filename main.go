package main

import(
	"log"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var fyneApp fyne.App
var window fyne.Window
var display *Display

func setup() {
	log.Println("Starting GoReq")
	fyneApp = app.NewWithID("com.sunshine.goreq")
	window = fyneApp.NewWindow("Simple HTTP Request")
	display = NewDisplay()
}

func main() {
	setup()
	
	content := display.Render()//container.NewBorder(nil, nil, nil, nil, nil)
	
	window.SetContent(content)
	
	window.CenterOnScreen()
	
	window.Resize(fyne.NewSize(600, 300))
	window.SetFixedSize(true)
	
	window.SetMaster()
	
	window.ShowAndRun()
}