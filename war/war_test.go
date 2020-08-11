package war

import (
    "testing"
    "github.com/stretchr/testify/assert"
    . "github.com/luc65r/cards/card"
)

func TestCompareFront(t *testing.T) {
    g := game{[]*player{
        {nil, &Deck{[]*Card{{Spade, 3, true}}}},
        {nil, &Deck{[]*Card{{Heart, King, true}}}},
        {nil, &Deck{[]*Card{{Diamond, 7, true}}}},
    }}
    assert.Equal(t, g.compareFront(), []int{1})
    g = game{[]*player{
        {nil, &Deck{[]*Card{{Heart, 10, true}}}},
        {nil, &Deck{[]*Card{{Spade, 6, true}}}},
        {nil, &Deck{[]*Card{{Diamond, 2, true}}}},
        {nil, &Deck{[]*Card{{Club, 10, true}}}},
    }}
    assert.Equal(t, g.compareFront(), []int{0, 3})
}
