package face

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestLoadFaces(t *testing.T) {
    require.Nil(t, Faces)
    require.NoError(t, LoadFaces())
    require.Len(t, Faces, 52)
}
