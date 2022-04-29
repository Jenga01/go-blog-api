package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func FindImages(htm string, sendBase64String chan []string) <-chan []string {
	go func() {
		var imgRE = regexp.MustCompile(`;base64[^]+["']([^"']+)["'.]`)
		imgs := imgRE.FindAllStringSubmatch(htm, -1)
		out := make([]string, len(imgs))
		for i := range out {
			out[i] = imgs[i][1]
		}
		sendBase64String <- out
	}()
	return sendBase64String
}

func Base64toPng(data string, ch chan []string) <-chan []string {
	Base64String := make(chan []string)
	go FindImages(data, Base64String)
	result := <-Base64String
	out := make([]string, len(result)-len(result))
	go func() {
		for _, i := range result {
			rawDecodedImage, err := base64.StdEncoding.DecodeString(i)
			if err != nil {
				if _, ok := err.(base64.CorruptInputError); ok {
					panic("\nbase64 input is corrupt, check service Key")
				}
				fmt.Sprintln("Error:", err)
			}
			dir := "images/"
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				err := os.Mkdir(dir, 0755)
				check(err)
			}

			pngFilename := createImageName(5) + ".png"
			dirAndFilename := pngFilename
			f, err := os.Create(filepath.Join(dir, pngFilename))

			if err != nil {
				log.Fatal(err)
			}
			if _, err := f.Write(rawDecodedImage); err != nil {
				log.Fatal(err)
			}
			if err := f.Sync(); err != nil {
				log.Fatal(err)
			}
			out = append(out, dirAndFilename)
			fmt.Println("base64toPng sent data is:", dirAndFilename)
		}
		fmt.Println("base64ToPng out data: ", out)
		ch <- out
		close(ch)
	}()
	return ch
}

func ReplaceBase64ToDecodedImage(data string) string {
	imageSrc := make(chan []string)
	go Base64toPng(data, imageSrc)
	result := <-imageSrc
	re := regexp.MustCompile(`data:image/[^]+["']([^"']+)[+?"']`)
	count := 0
	imgPath := func(string) string {
		count++
		return result[count-1] + `"`
	}
	replacedImages := re.ReplaceAllStringFunc(data, imgPath)
	return replacedImages
}

func createImageName(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
