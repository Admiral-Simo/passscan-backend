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

func TestExtractExtension(t *testing.T) {
	assert.Equal(t, ".png", extractExtension("something.png"), "should be equal")
	assert.Equal(t, ".jpg", extractExtension("something.jpg"), "should be equal")
	assert.Equal(t, ".jfif", extractExtension("something.jfif"), "should be equal")
	assert.Equal(t, ".pjpeg", extractExtension("something.pjpeg"), "should be equal")
	assert.Equal(t, ".pjp", extractExtension("something.pjp"), "should be equal")
	assert.Equal(t, ".avif", extractExtension("something.avif"), "should be equal")
	assert.Equal(t, ".gif", extractExtension("something.gif"), "should be equal")
	assert.Equal(t, ".png", extractExtension("something.png"), "should be equal")
}
