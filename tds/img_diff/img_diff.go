package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"crypto/sha256"
)

func main() {
	// read 1st image
	content, err := ioutil.ReadFile("image_1.jpg")
	if err != nil{
		log.Fatal(err)
	}
	// hash & print 1st image
	h := sha256.New()
	h.Write([]byte(content))
	fmt.Printf("%x", h.Sum(nil))
	// read 2nd image
	content2, err := ioutil.ReadFile("image_2.jpg")
	if err != nil{
		log.Fatal(err)
	}
	// hash & print 2nd image
	i := sha256.New()
	i.Write([]byte(content2))
	fmt.Printf("%x", i.Sum(nil))
	// read 3rd image
	content3, err := ioutil.ReadFile("image_3.jpg")
	if err != nil{
		log.Fatal(err)
	}
	// hash & print 3rd image
	j := sha256.New()
	j.Write([]byte(content3))
	fmt.Printf("%x", j.Sum(nil))
}
