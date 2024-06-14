package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckImage(t *testing.T) {
	assert.True(t, CheckImage("something.png"), "should be true")
	assert.True(t, CheckImage("something.jpg"), "should be true")
	assert.True(t, CheckImage("something.jfif"), "should be true")
	assert.True(t, CheckImage("something.pjpeg"), "should be true")
	assert.True(t, CheckImage("something.pjp"), "should be true")
	assert.True(t, CheckImage("something.avif"), "should be true")
	assert.True(t, CheckImage("something.gif"), "should be true")
	assert.True(t, CheckImage("something.png"), "should be true")

	assert.False(t, CheckImage("something.txt"), "should be false")
	assert.False(t, CheckImage("something.zip"), "should be false")
}

func TestExtractExtension(t *testing.T) {
	assert.Equal(t, ".png", ExtractExtension("something.png"), "should be equal")
	assert.Equal(t, ".jpg", ExtractExtension("something.jpg"), "should be equal")
	assert.Equal(t, ".jfif", ExtractExtension("something.jfif"), "should be equal")
	assert.Equal(t, ".pjpeg", ExtractExtension("something.pjpeg"), "should be equal")
	assert.Equal(t, ".pjp", ExtractExtension("something.pjp"), "should be equal")
	assert.Equal(t, ".avif", ExtractExtension("something.avif"), "should be equal")
	assert.Equal(t, ".gif", ExtractExtension("something.gif"), "should be equal")
	assert.Equal(t, ".png", ExtractExtension("something.png"), "should be equal")
}
