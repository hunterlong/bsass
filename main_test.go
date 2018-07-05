package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrueValue(t *testing.T) {

	val, px := TrueValue("27px;")

	t.Log(val, px)

}

func TestCollectVariables(t *testing.T) {
	vars := CollectVariables("scss/variables")
	assert.Equal(t, "15px", vars["$full"])
}
