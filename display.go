package main

import(
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
	
	headerListContainer *fyne.Container
	fieldListContainer *fyne.Container
	
	sendButton *widget.Button
	
	status *widget.Label
	response *widget.Label
	
	addToHistory func(r *Request)
}

//Create
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

//Render
func (d *Display) Render() *fyne.Container {
	optionTitle := widget.NewLabel("Options")
	optionsButton := widget.NewButton("Clear", d.clearOptions)
	optionsTop := container.NewBorder(nil, nil, nil, optionsButton, optionTitle)
	
	optionLabels := container.New(layout.NewVBoxLayout(), 
		widget.NewLabel("Method"),
		widget.NewLabel("URL"),
		widget.NewLabel("Headers"),
		widget.NewLabel(" "),
		widget.NewLabel("Fields"),
		widget.NewLabel(" "),
		widget.NewLabel("Secure"),
	)
	
	//Header
	newHeaderButton := widget.NewButton("+", func() {d.addItem("header")})
	newHeaderContainer := container.NewBorder(nil, nil, nil, newHeaderButton, d.newHeader)
	headerList := d.RenderList(d.Req.headers, "header")
	d.headerListContainer = container.New(layout.NewMaxLayout(), headerList)
	headerFinal := container.New(layout.NewVBoxLayout(), newHeaderContainer, d.headerListContainer)
	
	//Field
	newFieldButton := widget.NewButton("+", func() {d.addItem("field")})
	newFieldContainer := container.NewBorder(nil, nil, nil, newFieldButton, d.newField)
	fieldList := d.RenderList(d.Req.fields, "field")
	d.fieldListContainer = container.New(layout.NewMaxLayout(), fieldList)
	fieldFinal := container.New(layout.NewVBoxLayout(), newFieldContainer, d.fieldListContainer)
	
	options := container.New(layout.NewVBoxLayout(), d.method, d.url, headerFinal, fieldFinal, display.secure, d.sendButton)
	optionCombined := container.New(layout.NewHBoxLayout(), optionLabels, options)
	optionContainer := container.NewBorder(optionsTop, nil, nil, nil, optionCombined)
	
	d.status.SetText("Status:")
	statusButton := widget.NewButton("Clear", d.clearStatus)
	statusContainer := container.NewBorder(nil, nil, nil, statusButton, d.status)
	
	responseScroll := container.NewVScroll(d.response)
	responseContainer := container.NewBorder(statusContainer, nil, nil, nil, responseScroll)
	
	combinedContainer := container.NewBorder(nil, nil, optionContainer, nil, responseContainer)
	finalContainer := container.New(layout.NewMaxLayout(), combinedContainer)
	
	return finalContainer
}

func (d *Display) RenderList(items []string, listType string) *widget.List {
	list := widget.NewList(
		func() int {
			return len(items)
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("Name")
			button := widget.NewButton("-", nil)
			return container.NewBorder(nil, nil, nil, button, label)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {			
			//Set Name
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(items[i])
			button := o.(*fyne.Container).Objects[1].(*widget.Button)	
			button.OnTapped = func() {d.deleteItem(listType, i)}
		},
	)
	
	return list
}

func (d *Display) Update() {
	d.method.Selected = d.Req.method
	d.method.Refresh()
	d.url.Text = d.Req.url
	d.url.Refresh()
	
	newHeaders := d.RenderList(d.Req.headers, "header")
	d.headerListContainer.Objects[0] = newHeaders
	d.headerListContainer.Refresh()
	
	newFields := d.RenderList(d.Req.fields, "header")
	d.fieldListContainer.Objects[0] = newFields
	d.fieldListContainer.Refresh()
	
	d.secure.Checked = d.Req.secure
	d.secure.Refresh()
	
	d.status.SetText("Status: " + d.Req.response.status)
	d.response.SetText(d.Req.response.body)
}

//Actions
func (d *Display) sendRequest() {
	err := d.Req.Send()
	if err != nil {
		d.status.SetText("Status: Error")
		d.response.SetText(err.Error())
		return
	}
	
	d.status.SetText("Status: " + d.Req.response.status)
	d.response.SetText(d.Req.response.body)
	
	d.addToHistory(d.Req)
}

func (d *Display) clearStatus() {
	d.status.SetText("Status:")
	d.response.SetText("")
}

func (d *Display) clearOptions() {
	d.method.Selected = ""
	d.method.Refresh()
	d.url.Text = ""
	d.url.Refresh()
	d.headerListContainer.Objects[0] = d.RenderList([]string{"Content-Type=application/x-www-form-urlencoded"}, "header")
	d.headerListContainer.Refresh()
	d.fieldListContainer.Objects[0] =d.RenderList([]string{}, "field")
	d.fieldListContainer.Refresh()
	d.secure.Checked = false
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

func (d *Display) addItem(listType string) {
	if listType == "header" {
		item := d.newHeader.Text
		if item == "" {
			return
		}
		
		d.Req.headers = append(d.Req.headers, item)
		newList := d.RenderList(d.Req.headers, "header")
		
		d.headerListContainer.Objects[0] = newList
		d.headerListContainer.Refresh()
		d.newHeader.SetText("")
		
	} else {
		item := d.newField.Text
		if item == "" {
			return
		}
		
		d.Req.fields = append(d.Req.fields, item)
		newList := d.RenderList(d.Req.fields, "field")
		
		d.fieldListContainer.Objects[0] = newList
		d.fieldListContainer.Refresh()
		d.newField.SetText("")
	}
}

func (d *Display) deleteItem(listType string, index int) {
	if listType == "header" {
		d.Req.headers = append(d.Req.headers[:index], d.Req.headers[index+1:]...)
		newList := d.RenderList(d.Req.headers, "field")
		d.headerListContainer.Objects[0] = newList
		d.headerListContainer.Refresh()
	} else {
		d.Req.fields = append(d.Req.fields[:index], d.Req.fields[index+1:]...)
		newList := d.RenderList(d.Req.fields, "field")
		d.fieldListContainer.Objects[0] = newList
		d.fieldListContainer.Refresh()
	}
}