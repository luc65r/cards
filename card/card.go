package card

type Suit int

const (
    Club Suit = iota
    Diamond
    Heart
    Spade
)

func (s Suit) String() string {
    return [...]string{"♣", "♦", "♥", "♠"}[s]
}


type Value int

const (
    Jack Value = iota + 11
    Queen
    King
    Ace
)

func (v Value) String() string {
    v -= 2
    return [...]string{
        "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A",
    }[v]
}


type Card struct {
    Suit Suit
    Value Value
    Visible bool
}

func (c *Card) String() (s string) {
    if !c.Visible {
        s = "_"
    }
    s += c.Suit.String() + c.Value.String()
    return
}

func (c *Card) Turn() {
    c.Visible = !c.Visible
}
