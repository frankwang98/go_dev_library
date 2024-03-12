package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 创建一个写入CSV文件的示例
	writeExample()

	// 创建一个读取CSV文件的示例
	readExample()
}

func writeExample() {
	// 创建一个写入CSV文件的示例
	file, err := os.Create("example.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "25", "New York"},
		{"Bob", "30", "Los Angeles"},
		{"Charlie", "20", "Chicago"},
	}

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("CSV file created successfully")
}

func readExample() {
	// 创建一个读取CSV文件的示例
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, row := range records {
		fmt.Println(row)
	}
}
