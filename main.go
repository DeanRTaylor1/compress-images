package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/h2non/bimg"
)

func main() {
	// Read the images directory
	dir := "./images"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please make sure your images are saved in the images folder")
		log.Fatal(err)
	}
	for _, file := range files {
		compressImage(dir, file)
	}
}

func compressImage(dir string, image fs.DirEntry) {
	// Read the image
	// Resize the image
	// Check if the image is compressed
	// Create a new directory for the compressed images
	// Get the file size
	// Write the image to the new directory
	// Print the results
	buffer, err := bimg.Read(dir + "/" + image.Name())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

	}

	newImage, err := bimg.NewImage(buffer).Resize(531, 800)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

	}

	size, err := bimg.NewImage(newImage).Size()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if size.Width == 800 || size.Height == 531 {
		fmt.Println("Image is compressed")
	}

	newDir := "./compressed"
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		err := os.Mkdir(newDir, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println("Directory created: ", newDir)
	}

	fileSize := bimg.NewImage(newImage).Length()

	fmt.Printf("Successfully resized %s. \nNew image height: %d, resized image width: %d.\nFsile size: %d bytes\n", image.Name(), size.Height, size.Width, fileSize)
	fmt.Println("Writing image to: ", newDir+"/"+strings.Replace(image.Name(), ".jpg", "", -1)+"-compressed.jpg")
	bimg.Write(newDir+"/"+strings.Replace(image.Name(), ".jpg", "", -1)+"-compressed.jpg", newImage)

}
