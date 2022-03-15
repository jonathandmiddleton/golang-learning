package main

import "fmt"

func main() {

	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchett"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher)
	fmt.Println(year, pageNumber)
	fmt.Println(grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher)
	fmt.Println(year, pageNumber)
	fmt.Println(grade)

}

func calculateCost(pageNumber int, grade float32) float32 {
	pageNumber = 4
	grade = 9.0
	return float32(pageNumber) * grade
	fmt.Println(calculateCost)
}
