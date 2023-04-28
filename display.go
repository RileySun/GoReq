package main

import(
	"log"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

var REQTYPES []string = []string{"GET", "POST"}

type Display struct {
	Req *Request

	method *widget.Select
	url *widget.Entry
	newHeader *widget.Entry
	newField *widget.Entry
	secure *widget.Check
	
	sendButton *widget.Button
	
	status *widget.Label
	response *widget.Label
}

func NewDisplay() *Display {
	display := &Display{
		Req:NewRequest(),
		url:widget.NewEntry(),
		newHeader:widget.NewEntry(),
		newField:widget.NewEntry(),
		response:widget.NewLabel(""),
		status:widget.NewLabel(""),
	}
	
	display.method = widget.NewSelect(REQTYPES, display.changeMethod)
	display.sendButton = widget.NewButtonWithIcon("Send", theme.MailForwardIcon(), display.sendRequest)
	display.url.OnChanged = display.changeURL
	display.secure = widget.NewCheck("", display.changeSecure)
	display.response.Wrapping = 2
	
	return display
}

func (d *Display) Render() *fyne.Container {
	optionTitle := widget.NewLabel("Options")
	optionLabels := container.New(layout.NewVBoxLayout(), 
		widget.NewLabel("Method"),
		widget.NewLabel("URL"),
		widget.NewLabel("Headers"),
		widget.NewLabel("Fields"),
		widget.NewLabel("Secure"),
	)
	options := container.New(layout.NewVBoxLayout(), d.method, d.url, d.newHeader, d.newField, display.secure, d.sendButton)
	optionCombined := container.New(layout.NewHBoxLayout(), optionLabels, options)
	optionContainer := container.NewBorder(optionTitle, nil, nil, nil, optionCombined)

	statusContainer := container.NewBorder(nil, nil, nil, nil, d.status)
	
	responseScroll := container.NewVScroll(d.response)
	responseContainer := container.NewBorder(statusContainer, nil, nil, nil, responseScroll)
	
	combinedContainer := container.NewBorder(nil, nil, optionContainer, nil, responseContainer)
	finalContainer := container.New(layout.NewMaxLayout(), combinedContainer)
	
	return finalContainer
}


//Actions
func (d *Display) sendRequest() {
	status, response, err := d.Req.Send()
	if err != nil {
		log.Println("Send Err")
		d.status.SetText("Error")
		d.response.SetText(err.Error())
		return
	}
	d.status.SetText(status)
	d.response.SetText(response)
}
func (d *Display) changeMethod(newMethod string) {
	d.Req.method = newMethod
}

func (d *Display) changeURL(newURL string) {
	d.Req.url = newURL
}

func (d *Display) changeSecure(secure bool) {
	d.Req.secure = secure
}