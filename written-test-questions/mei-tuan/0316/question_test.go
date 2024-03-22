// @Author: Ciusyan 3/21/24

package _316

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLegalWord(t *testing.T) {

	assert.Equal(t, 0, legalWord("Abc"))
	assert.Equal(t, 0, legalWord("abcd"))
	assert.Equal(t, 0, legalWord("ABCCDA"))
	assert.Equal(t, 0, legalWord("A"))
	assert.Equal(t, 0, legalWord(""))
	assert.Equal(t, 1, legalWord("AbC"))
	assert.Equal(t, 1, legalWord("AbCC"))
	assert.Equal(t, 1, legalWord("aCCC"))
	assert.Equal(t, 2, legalWord("ACCcd"))
	assert.Equal(t, 3, legalWord("ACCcdd"))
	assert.Equal(t, 1, legalWord("ACcdd"))
	assert.Equal(t, 2, legalWord("aACcdd"))

}
