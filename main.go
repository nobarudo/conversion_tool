package main

import (
	"log"
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
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "エンコード",
						OnClicked: func() {
							outTE.SetText(url.QueryEscape(inTE.Text()))
						},
					},
					PushButton{
						Text: "デコード",
						OnClicked: func() {
							str, err := url.QueryUnescape(inTE.Text())
							if err != nil {
								log.Println(err)
							}
							outTE.SetText(str)
						},
					},
					PushButton{
						Text: "クリア",
						OnClicked: func() {
							inTE.SetText("")
							outTE.SetText("")
						},
					},
				},
			},
		},
	}.Run()
}
