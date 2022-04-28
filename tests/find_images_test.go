package tests

import (
	utils "first/Utils"
	"github.com/go-playground/assert/v2"
	"testing"
)

var data = `<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==" alt="Red dot" />
            <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==" alt="Green dot" />`

func TestFindImages(t *testing.T) {
	Base64String := make(chan []string)
	go utils.FindImages(data, Base64String)
	result := <-Base64String

	if len(result) > 0 {
		assert.NotEqual(t, result, "")
	} else {
		assert.Equal(t, result, "")
	}
}

func TestBase64ToPNG(t *testing.T) {
	base64ImageSrc := make(chan []string)
	go utils.Base64toPng(data, base64ImageSrc)
	decodedImagePath := <-base64ImageSrc

	assert.NotEqual(t, decodedImagePath, base64ImageSrc)
}

func TestReplacedImageSrc(t *testing.T) {
	replacedImages := utils.ReplaceBase64ToDecodedImage(data)

	assert.NotEqual(t, replacedImages, data)
}
