package main

import (
	"net/url"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "convert_tool",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "変換",
				OnClicked: func() {
					outTE.SetText(url.QueryEscape(inTE.Text()))
				},
			},
		},
	}.Run()
}
