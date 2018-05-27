package csvdict

// csvdict implements a dictionary reader for CSV files

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// Basic reader
type DictReader struct {
	cReader *csv.Reader
	header  []string
}

func NewDictReader(r io.Reader) *DictReader {
	cr := csv.NewReader(r)
	h, err := cr.Read()
	if err != nil {
		fmt.Println("I had trouble reading the header row. Is the file in the correct format?")
		os.Exit(1)
	}

	return &DictReader{
		cReader: cr,
		header:  h,
	}
}

func (d *DictReader) Read() (csvMap map[string]string, err error) {
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

func (d *DictReader) GetHeaderRow() (header []string) {
	return d.header
}

type DictWriter struct {
	cWriter *csv.Writer
	header  []string
}

func NewDictWriter(w io.Writer, h []string) *DictWriter {

	return &DictWriter{
		cWriter: csv.NewWriter(w),
		header:  h,
	}
}

func (w *DictWriter) Flush() {
	w.cWriter.Flush()
}

func (w *DictWriter) WriteHeaders() {
	w.cWriter.Write(w.header)
}

func (w *DictWriter) Write(csvMap map[string]string) error {
	var newLine []string
	for _, field := range w.header {
		newLine = append(newLine, csvMap[field])
	}
	err := w.cWriter.Write(newLine)
	if err != nil {
		return err
	}
	return err
}
