package hexgrid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//    _ _           _ _
//  /     \       /     \
// /( 0,0) \ _ _ /(2,-1) \
// \       /     \       /
//  \ _ _ / (1,0) \ _ _ /
//  /     \       /     \
// / (0,1) \ _ _ / (2,0) \
// \       /     \       /
//  \ _ _ / (1,1) \ _ _ /
//        \       /
//         \ _ _ /
func TestHexagonalGrid(t *testing.T) {
	assert.Len(t, HexagonalGrid(3), 1+6+12+18)
}
