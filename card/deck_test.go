package card

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
    d1, d2, d3 := new(Deck), new(Deck), new(Deck)
    d1.Cards = []*Card{
        {Club, 9, true},
        {Diamond, King, false},
    }
    d2.Cards = []*Card{
        {Club, 9, true},
        {Diamond, King, false},
        {Spade, Ace, true},
    }
    d3.Cards = []*Card{
        {Club, 9, true},
        {Diamond, King, false},
        {Spade, Ace, true},
        {Spade, Ace, false},
        {Diamond, 5, true},
    }

    d1.Push(&Card{Spade, Ace, true})
    assert.Equal(t, d1, d2)

    d1.Push()
    d1.Push(&Card{Spade, Ace, false}, &Card{Diamond, 5, true})
    assert.Equal(t, d1, d3)

    d1.Push(d2.Cards...)
    assert.Len(t, d1.Cards, 8)
}

func TestPop(t *testing.T) {
    d1, d2 := new(Deck), new(Deck)
    d1.Cards = []*Card{
        {Club, 9, true},
        {Diamond, King, false},
        {Spade, Ace, true},
    }
    d2.Cards = []*Card{
        {Club, 9, true},
        {Diamond, King, false},
    }

    assert.Equal(t, d1.Pop(), &Card{Spade, Ace, true})
    assert.Equal(t, d1, d2)
    assert.Equal(t, d1.Pop(), &Card{Diamond, King, false})
    assert.Equal(t, d1.Pop(), &Card{Club, 9, true})
    assert.Panics(t, func() {d1.Pop()}, "Should panic when there are no cards")
}

func TestInsert(t *testing.T) {
    d1, d2, d3 := new(Deck), new(Deck), new(Deck)
    d1.Cards = []*Card{
        {Club, 9, true},
        {Diamond, King, false},
    }
    d2.Cards = []*Card{
        {Spade, Ace, true},
        {Club, 9, true},
        {Diamond, King, false},
    }
    d3.Cards = []*Card{
        {Spade, Ace, false},
        {Diamond, 5, true},
        {Spade, Ace, true},
        {Club, 9, true},
        {Diamond, King, false},
    }

    d1.Insert(&Card{Spade, Ace, true})
    assert.Equal(t, d1, d2)

    d1.Insert()
    d1.Insert(&Card{Spade, Ace, false}, &Card{Diamond, 5, true})
    assert.Equal(t, d1, d3)

    d1.Insert(d2.Cards...)
    assert.Len(t, d1.Cards, 8)
}

func TestNewDeck(t *testing.T) {
    d1 := NewDeck()
    assert.Len(t, d1.Cards, 52, "A new deck should have 52 cards")
    for _, c := range d1.Cards {
        assert.NotNil(t, c, "No pointer should be nil")
    }
    d2 := NewDeck()
    assert.Equal(t, d1, d2, "Two non-shuffled new decks should be identical")
}

func TestSuffle(t *testing.T) {
    d1 := NewDeck()
    d1.Shuffle()
    assert.Len(t, d1.Cards, 52, "A new shuffled deck should have 52 cards")
    for _, c := range d1.Cards {
        assert.NotNil(t, c, "No pointer should be nil")
    }
    d2 := NewDeck()
    d2.Shuffle()
    assert.NotEqual(t, d1, d2, "Two shuffled decks should be different")
}

func BenchmarkShuffle(b *testing.B) {
    d := NewDeck()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        d.Shuffle()
    }
}

func TestReverse(t *testing.T) {
    d := new(Deck)
    d.Cards = []*Card{}
    d.Reverse()
    assert.Empty(t, d.Cards)

    d.Cards = []*Card{
        {Heart, 7, true},
        {Spade, Ace, false},
    }
    d.Reverse()
    assert.Equal(t, d.Cards, []*Card{
        {Spade, Ace, true},
        {Heart, 7, false},
    })
    
    d.Cards = []*Card{
        {Heart, Queen, true},
        {Club, 10, true},
        {Club, 3, false},
    }
    d.Reverse()
    assert.Equal(t, d.Cards, []*Card{
        {Club, 3, true},
        {Club, 10, false},
        {Heart, Queen, false},
    })

    d1 := NewDeck()
    d1.Shuffle()
    d2 := new(Deck)
    d2.Cards = make([]*Card, 52)
    for i, c := range d1.Cards {
        d2.Cards[i] = c
    }
    d1.Reverse()
    d1.Reverse()
    assert.Equal(t, d1, d2, "Reversing two times should change nothing")
}

func BenchmarkReverse(b *testing.B) {
    d := NewDeck()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        d.Reverse()
    }
}

func TestString(t *testing.T) {
    d := new(Deck)
    d.Cards = []*Card{
        {Spade, 6, false},
    }
    assert.Equal(t, d.String(), "_♠6")

    d.Cards = []*Card{
        {Heart, Queen, true},
        {Club, 10, true},
        {Club, 3, false},
    }
    assert.Equal(t, d.String(), "♥Q ♣10 _♣3")
}

func TestDealAll(t *testing.T) {
    d := NewDeck()
    d.Shuffle()
    d1, d2, d3 := new(Deck), new(Deck), new(Deck)
    d.DealAll(d1, d2)
    assert.Len(t, d.Cards, 0)
    assert.Len(t, d1.Cards, 26)
    assert.Len(t, d2.Cards, 26)
    d = NewDeck()
    d.Shuffle()
    d.DealAll(d1, d2, d3)
    assert.Len(t, d.Cards, 1)
    assert.Len(t, d1.Cards, 43)
    assert.Len(t, d2.Cards, 43)
    assert.Len(t, d3.Cards, 17)
}
