package card

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSuit(t *testing.T) {
    assert.Less(t, int(Club), int(Diamond), "♣ < ♦")
    assert.Less(t, int(Diamond), int(Heart), "♦ < ♥")
    assert.Less(t, int(Heart), int(Spade), "♥ < ♠")
}

func TestSuitString(t *testing.T) {
    assert.Equal(t, Club.String(), "♣", "Club should be represented by ♣")
    assert.Equal(t, Diamond.String(), "♦", "Diamond should be represented by ♦")
    assert.Equal(t, Heart.String(), "♥", "Heart should be represented by ♥")
    assert.Equal(t, Spade.String(), "♠", "Spade should be represented by ♠")
}

func TestValueString(t *testing.T) {
    assert.Equal(t, Value(2).String(), "2", "The value 2 should be represented by 2")
    assert.Equal(t, Value(7).String(), "7", "The value 7 should be represented by 2")
    assert.Equal(t, Value(10).String(), "10", "The value 10 should be represented by 10")
    assert.Equal(t, Queen.String(), "Q", "The queen should be represented by Q")
    assert.Equal(t, Ace.String(), "A", "The ace should be represented by A")
}

func TestCardString(t *testing.T) {
    var tests = []struct {
        c Card
        want string
    }{
        {Card{Heart, 2, true}, "♥2"},
        {Card{Club, 10, false}, "_♣10"},
        {Card{Spade, Jack, false}, "_♠J"},
        {Card{Diamond, Queen, true}, "♦Q"},
    }

    for _, tt := range tests {
        non := ""
        if !tt.c.Visible {
            non += "non-"
        }
        assert.Equal(
            t, tt.c.String(), tt.want,
            "A %vvisible %v of %v should be represented by %v",
            non, tt.c.Value, tt.c.Suit, tt.want,
        )
    }
}

func TestCardTurn(t *testing.T) {
    c1 := Card{Heart, 4, false}
    c1.Turn()
    assert.True(t, c1.Visible, "When we turn a non-visible card, it should become visible")
    c2 := Card{Spade, Jack, true}
    c2.Turn()
    assert.False(t, c2.Visible, "When we turn a visible card, it should become non-visible")
}
