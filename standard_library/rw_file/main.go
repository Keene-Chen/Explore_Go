package main



func main() {
	readFile := "./res/data.json"
	writeFile := "./res/sorted_data.json"
	writeFile1 := "./res/sorted_data1.json"

	SortJSON(readFile, writeFile)
	SortJSON1(readFile, writeFile1)
}
