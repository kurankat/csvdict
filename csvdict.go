// Package csvdict is a Short Go library that extends encoding/csv to handle importing CSV content with headers as a map (dictionary)
package csvdict

// csvdict implements a dictionary reader for CSV files

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// DictReader contains a reference to a CSV reader and a header row slice, used to map
// each CSV row to its headers
type DictReader struct {
	cReader *csv.Reader
	header  []string
}

// NewDictReader takes an io.Reader as an argument and returns a reference to a DictReader
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

// Read reads a single line of a CSV file and returns a map where each field is mapped to its header
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

// ReadAll reads all the remaining records from r. Each record is a map of fields.
// A successful call returns err == nil, not err == io.EOF. Because ReadAll is defined
// to read until EOF, it does not treat end of file as an error to be reported.
func (d *DictReader) ReadAll() (records []map[string]string, err error) {
	for {
		record, err := d.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}
}

// GetHeaderRow returns the header row of the DictReader object
func (d *DictReader) GetHeaderRow() (header []string) {
	return d.header
}

// DictWriter contains a reference to a CSV Writer object and a header row slice
type DictWriter struct {
	cWriter *csv.Writer
	header  []string
}

// NewDictWriter takes an io.Writer and a slice containing the header row
// as arguments and returns a reference to a DictWriter
func NewDictWriter(w io.Writer, h []string) *DictWriter {

	return &DictWriter{
		cWriter: csv.NewWriter(w),
		header:  h,
	}
}

// Flush ensures that the writer is flushed to file
func (w *DictWriter) Flush() {
	w.cWriter.Flush()
}

// WriteHeaders writes the header row to file
func (w *DictWriter) WriteHeaders() {
	w.cWriter.Write(w.header)
}

// Write takes a map with headers as indices and field contents as values, and writes it
// to file as CSV
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

// WriteAll writes multiple CSV records to w using Write and then calls Flush.
func (w *DictWriter) WriteAll(records []map[string]string) error {
	for _, record := range records {
		err := w.Write(record)
		if err != nil {
			return err
		}
	}

	w.cWriter.Flush()
	return w.cWriter.Error()
}
