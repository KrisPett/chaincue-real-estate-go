package utilities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatString(t *testing.T) {
	// Given
	testString1 := "VILLA"
	testString2 := "VACATION_HOME"
	testString3 := "CONDOMINIUM"

	expected1 := "Villa"
	expected2 := "Vacation Home"
	expected3 := "Condominium"

	// When
	actual1 := formatString(testString1)
	actual2 := formatString(testString2)
	actual3 := formatString(testString3)

	// Then
	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
	assert.Equal(t, expected3, actual3)
}
