package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func powerrename(directory *string, match *string, replace *string) {
	fmt.Println("--Power Rename--")
	if *match == "null" {
		fmt.Println("Please provide a string to match with.")
		fatal()
	}
	if *replace == "null" {
		fmt.Println("Please provide a string to replace matching text with.")
		fatal()
	}
	if (*directory)[len(*directory)-1:] != "/" {
		*directory += "/"
	}
	files, err := ioutil.ReadDir(*directory)
	if err != nil {
		fatal()
	}
	for _, f := range files {
		filename := f.Name()
		if strings.Contains(filename, *match) {
			output := strings.Replace(filename, *match, *replace, -1)
			fmt.Println(filename + " => " + output)
			err = os.Rename(*directory+filename, *directory+output)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
