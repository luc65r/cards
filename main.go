package main

import (
    "fmt"
    /*
    _ "github.com/luc65r/cards/card"
    "github.com/luc65r/cards/face"

    "fyne.io/fyne"
    "fyne.io/fyne/app"
    "fyne.io/fyne/widget"
    */
    "github.com/luc65r/cards/war"
)

func main() {
    /*
    err := face.LoadFaces()
    if err != nil {
        fmt.Println(err)
        return
    }

    a := app.New()
    w := a.NewWindow("Test")
    w.SetContent(&widget.Box{Children: []fyne.CanvasObject{
        &widget.Label{Text: "Hello test!"},
        &widget.Button{Text: "Press", OnTapped: func() {
            a.Quit()
        }},
    }})

    w.ShowAndRun()
    */
    err := war.Play([]string{"test0", "test1", "test2", "test3"})
    if err != nil {
        fmt.Println(err)
    }
}
