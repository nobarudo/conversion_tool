package main

import (
	"encoding/base64"
	"html"
	"log"
	"net/url"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit
	var menu *walk.ComboBox

	var menuItem = []string{
		"URL",
		"Base64",
		"HTML",
		"Base64→URL(デコードはURL→Base64)",
	}

	MainWindow{
		Title:   "conversion_tool",
		MinSize: Size{600, 400},
		Size:    Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			ComboBox{
				AssignTo:     &menu,
				Model:        menuItem,
				CurrentIndex: 0,
			},
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE, VScroll: true},
					TextEdit{AssignTo: &outTE, ReadOnly: true, VScroll: true},
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "エンコード",
						OnClicked: func() {
							switch menu.CurrentIndex() {
							case 0:
								outTE.SetText(url.QueryEscape(inTE.Text()))
							case 1:
								outTE.SetText(base64.StdEncoding.EncodeToString([]byte(inTE.Text())))
							case 2:
								outTE.SetText(html.EscapeString(inTE.Text()))
							case 3:
								outTE.SetText(url.QueryEscape(base64.StdEncoding.EncodeToString([]byte(inTE.Text()))))
							}
						},
					},
					PushButton{
						Text: "デコード",
						OnClicked: func() {
							switch menu.CurrentIndex() {
							case 0:
								str, err := url.QueryUnescape(inTE.Text())
								if err != nil {
									log.Println(err)
								}
								outTE.SetText(str)
							case 1:
								str, err := base64.StdEncoding.DecodeString(inTE.Text())
								if err != nil {
									log.Println("error: ", err)
								}
								outTE.SetText(string(str))
							case 2:
								outTE.SetText(html.UnescapeString(inTE.Text()))
							case 3:
								str1, err := url.QueryUnescape(inTE.Text())
								if err != nil {
									log.Println(err)
								}
								str2, err := base64.StdEncoding.DecodeString(str1)
								if err != nil {
									log.Println("error: ", err)
								}
								outTE.SetText(string(str2))
							}
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
