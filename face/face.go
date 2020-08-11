package face

import (
    "fmt"
    . "github.com/luc65r/cards/card"
    "fyne.io/fyne"
)

var Faces map[Card]*fyne.Resource

func LoadFaces() error {
    Faces = make(map[Card]*fyne.Resource)

    suitNames := [...]string{"C", "D", "H", "S"}

    for s := Club; s <= Spade; s++ {
        for v := Value(2); v <= Ace; v++ {
            r, err := fyne.LoadResourceFromPath(
                fmt.Sprintf("/home/lucas/cards/face/%v%v.svg", suitNames[s], v),
            )
            if err != nil {
                return err
            }

            Faces[Card{s, v, true}] = &r
        }
    }

    return nil
}
