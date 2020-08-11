package war

import (
    "fmt"
    "errors"
    . "github.com/luc65r/cards/card"
)

type game struct {
    players map[string]*player
}

type player struct {
    deck *Deck
    front *Deck
    out bool
}

func Play(names []string) error {
    if hasDuplicates(names) {
        return errors.New("Duplicates names!")
    }

    g := initGame(names)

    for {
        sum := 0
        for _, p := range(g.players) {
            sum += len(p.deck.Cards)
        }
        if sum != 52 {
            return errors.New(fmt.Sprintf("There are %v cards", sum))
        }

        fmt.Println("\n_____ Decks _____")
        for n, p := range(g.players) {
            if p.out {continue}
            fmt.Println(n, p.deck)

            p.front.Push(p.deck.Pop())
            p.front.Cards[0].Turn()
        }

        winners := g.compareFront(names)
        if len(winners) == 1 {
            g.pickUp(winners[0])
        } else {
            g.pickUp(g.war(winners))
        }

        alive := make([]string, 0, len(names))
        for n, p := range(g.players) {
            if p.out {
                continue
            }
            if len(p.deck.Cards) == 0 {
                fmt.Printf("\nPlayer %v is out!\n", n)
                p.out = true
                continue
            }
            alive = append(alive, n)
        }
        if len(alive) == 1 {
            fmt.Printf("\nPlayer %v has won the game!\n", alive[0])
            return nil
        }
    }

    return nil
}

func (g *game) compareFront(names []string) []string {
    fmt.Println("\n____ Fronts _____")
    max := 0
    win := make([]string, 0, len(g.players))
    for _, n := range names {
        p := g.players[n]
        if p.out {continue}
        fmt.Println(n, p.front)
        v := int(p.front.Cards[len(p.front.Cards)-1].Value)
        if v == max {
            win = append(win, n)
        } else if v > max {
            max = v
            win = win[:1]
            win[0] = n
        }
    }

    return win
}

func (g *game) pickUp(winner string) {
    fmt.Printf("Player %v wins the cards!\n", winner)
    for _, p := range(g.players) {
        if p.out {continue}
        cs := p.front
        for _, c := range cs.Cards {
            c.Visible = false
        }
        g.players[winner].deck.Insert(cs.Cards...)
        p.front.Cards = p.front.Cards[:0]
    }
}

func (g *game) war(names []string) string{
    fmt.Println("War between", names, "!")
    for _, n := range names {
        p := g.players[n]
        if len(p.deck.Cards) < 2 {
            continue
        }
        p.front.Push(p.deck.Pop())
        p.front.Push(p.deck.Pop())
        p.front.Cards[len(p.front.Cards)-1].Turn()
    }

    winners := g.compareFront(names)
    if len(winners) > 1 {
        return g.war(winners)
    }
    return winners[0]
}

func hasDuplicates(names []string) bool {
    for i := 0; i < len(names); i++ {
        for j := 0; j < i; j++ {
            if names[i] == names[j] {
                return true
            }
        }
    }
    return false
}

func initGame(names []string) *game {
    g := new(game)
    g.players = make(map[string]*player)
    decks := make([]*Deck, len(names))

    for i := range names {
        decks[i] = new(Deck)
    }

    tmpDeck := NewDeck()
    tmpDeck.Shuffle()
    tmpDeck.DealAll(decks...)

    for i, n := range names {
        g.players[n] = &player{
            deck: decks[i],
            front: new(Deck),
            out: false,
        }
    }
    tmpDeck, decks = nil, nil

    return g
}
