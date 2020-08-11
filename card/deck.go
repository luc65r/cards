package card

import (
    "math/rand"
    "time"
    "strings"
)

type Deck struct {
    Cards []*Card
}

func (d *Deck) Push(c ...*Card) {
    d.Cards = append(d.Cards, c...)
}

func (d *Deck) Pop() (c *Card) {
    l := len(d.Cards) - 1
    c = d.Cards[l]
    d.Cards = d.Cards[:l]
    return
}

func (d *Deck) Insert(c ...*Card) {
    d.Cards = append(c, d.Cards...)
}

func NewDeck() (d *Deck) {
    d = new(Deck)
    d.Cards = make([]*Card, 52)
    for s := Club; s <= Spade; s++ {
        for v := Value(2); v <= Ace; v++ {
            d.Cards[int(s) * 13 + int(v) - 2] = &Card{s, v, false}
        }
    }
    return
}

func (d *Deck) Shuffle() {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(d.Cards), func(i, j int) {
        d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
    })
}

func (d *Deck) Reverse() {
    for i, j := 0, len(d.Cards)-1; i < j; i, j = i+1, j-1 {
        d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
        d.Cards[i].Turn()
        d.Cards[j].Turn()
    }

    // Turn the middle card    
    if len(d.Cards) % 2 == 1 {
        d.Cards[len(d.Cards) / 2].Turn()
    }
}

func (d *Deck) String() string {
    s := make([]string, len(d.Cards))
    for i, c := range d.Cards {
        s[i] = c.String()
    }
    return strings.Join(s, " ")
}

func (d *Deck) DealAll(ds ...*Deck) {
    for len(d.Cards) >= len(ds) {
        for _, i := range ds {
            i.Push(d.Pop())
        }
    }
}
