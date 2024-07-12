package httpadapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckImage(t *testing.T) {
	assert.True(t, checkImage("something.png"), "should be true")
	assert.True(t, checkImage("something.jpg"), "should be true")
	assert.True(t, checkImage("something.jfif"), "should be true")
	assert.True(t, checkImage("something.pjpeg"), "should be true")
	assert.True(t, checkImage("something.pjp"), "should be true")
	assert.True(t, checkImage("something.avif"), "should be true")
	assert.True(t, checkImage("something.gif"), "should be true")
	assert.True(t, checkImage("something.png"), "should be true")

	assert.False(t, checkImage("something.txt"), "should be false")
	assert.False(t, checkImage("something.zip"), "should be false")
}
