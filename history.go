package main

import(
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

type History struct {
	List []*HistoryItem
	loadReq func(req *Request)
}

type HistoryItem struct {
	name string
	req *Request
}

//Create
func NewHistory() *History {
	return &History{}
}

func (h *History) Render() *fyne.Container {
	title := widget.NewLabel("History")
	button := widget.NewButton("Clear", h.clearHistory)
	historyTop := container.NewBorder(nil, nil, nil, button, title)

	list := widget.NewList(
		func() int {
			return len(h.List)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Name")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(h.List[i].name)
		},
	)
	
	list.OnSelected = func(id widget.ListItemID) {
		item := h.List[int(id)]
		h.loadReq(item.req)
	}
	
	
	return container.NewBorder(historyTop, nil, nil, nil, list)
}

//Actions
func (h *History) Add(req *Request) {
	item := &HistoryItem {
		name:getHostname(req.url),
		req:req,
	}
	h.List = append(h.List, item)
}

func (h *History) clearHistory() {
	h.List = []*HistoryItem{}
}
