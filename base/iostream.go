package base

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func readCSVFile(path string) ([][]string, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {

		}
	}(inputFile)

	reader := csv.NewReader(inputFile)
	return reader.ReadAll()
}

func Write2CsvFile() {
	records, err := readCSVFile(".\\base\\testdata.csv")
	if err != nil {
		fmt.Println("error open file:", err)
		return
	}
	outputFile, err := os.Create("newfile.csv")
	if err != nil {
		fmt.Println("error open new file:", err)
		return
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(outputFile)

	writer := csv.NewWriter(outputFile)
	// 将数据刷新到文件中
	defer writer.Flush()

	// records 二维数组
	for i, record := range records {
		if i == 0 {
			record = append(record, "composedField")
		} else {
			record = append(record, strings.Join(record, "-"))
		}

		if err := writer.Write(record); err != nil {
			fmt.Printf("error writing record to CSV:%v", err)
		}
	}

	fmt.Println("Process CSV File successfully")
}
