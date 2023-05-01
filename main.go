package main

import(
	"log"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var fyneApp fyne.App
var window fyne.Window
var content *fyne.Container

var display *Display
var history *History

func setup() {
	log.Println("Starting GoReq")
	fyneApp = app.NewWithID("com.sunshine.goreq")
	window = fyneApp.NewWindow("Simple HTTP Request")
	
	history = NewHistory()
	history.loadReq = loadRequest
	
	display = NewDisplay()
	display.addToHistory = addToHistory
}

func main() {
	setup()
	
	content = Render()
	
	window.SetContent(content)
	
	window.CenterOnScreen()
	
	window.Resize(fyne.NewSize(700, 400))
	window.SetFixedSize(true)
	
	window.SetMaster()
	
	window.ShowAndRun()
}

func Render() *fyne.Container {
	return container.NewBorder(nil, nil, history.Render(), nil, display.Render())
}

func loadRequest(req *Request) {
	display.Req = req
	display.Update()
}

func addToHistory(req *Request) {
	history.Add(req)
	content = Render()
	window.SetContent(content)
}