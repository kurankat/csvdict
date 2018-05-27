package csvdict

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type CsvDictReader struct {
	cReader *csv.Reader
	header  []string
}

func NewCsvDictReader(r io.Reader) *CsvDictReader {
	cr := csv.NewReader(r)
	h, err := cr.Read()
	if err != nil {
		fmt.Println("I had trouble reading the header row")
		os.Exit(1)
	}

	return &CsvDictReader{
		cReader: cr,
		header:  h,
	}
}

func (d *CsvDictReader) Read() (csvMap map[string]string, err error) {
	csvMap = make(map[string]string)
	line, err := d.cReader.Read()
	if err == io.EOF {
		return csvMap, err
	}
	headers := d.header

	for index, field := range headers {
		csvMap[field] = strings.TrimSuffix(strings.TrimSpace(line[index]), "\n")
	}

	return csvMap, err
}

type CsvDictWriter struct {
	cWriter *csv.Writer
	header  []string
}

func NewCsvDictWriter(w io.Writer, h []string) *CsvDictWriter {

	return &CsvDictWriter{
		cWriter: csv.NewWriter(w),
		header:  h,
	}
}

func (d *CsvDictReader) GetHeaderRow() (header []string) {
	return d.header
}

func (d *CsvDictWriter) WriteHeaders() {
	d.cWriter.Write(d.header)
}

func (d *CsvDictWriter) Write(csvMap map[string]string) {
	var newLine []string
	for _, field := range d.header {
		newLine = append(newLine, csvMap[field])
	}
	err := d.cWriter.Write(newLine)
	if err != nil {
		fmt.Println("I had trouble writing to file. Are permissions correct?")
		os.Exit(1)
	}
}
