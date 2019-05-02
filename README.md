# csvdict

Golang short library that extends encoding/csv to handle importing CSV content with headers as a dictionary (map)

## Usage

`NewDictReader` takes an argument of an io.Reader object and returns a pointer to a `DictReader` object that contains both a CSV Reader and a separate slice with the first line of the CSV file, assumed to be the header row.

Similarly, `NewDictWriter` takes two arguments, an io.Writer and a slice of headers, returning a pointer to a `DictWriter` object (and an error value, hopefully nil), a CSV Writer that will write the value corresponding to each header in the slice to file, in order:

`dictReader, err := csvdict.NewDictReader(ioReader)`
`dictWriter := csvdict.NewDictWriter(ioWriter, headerSlice)`

To read a line, use `Read()`, which returns a map with the header row contents as the indices and the contents of the row being read as the value. For example:

`rowMap, err := dictReader.Read()`

Optionally, to read all the contents, you can use `ReadAll()`:

`records, err := r.ReadAll()`

To write to file, first write the header row line to file:

`dictWriter.WriteHeaders()`

Then to write records one line at a time, use `Write(csvMap)`, which takes as an argument a map with indices matching the headers of the CSV file to be written:

`err = dictWriter.Write(csvMap)`

f writing line by line, to ensure that everything has been written to file before the end of execution, call `Flush()`:

`dictWriter.Flush()`

Optionally, you can use `WriteAll(csvMapSlice)`, which takes a slice of CSV maps, and calls `Flush()` automatically:

`dictWriter.Writeall(csvMapSlice)`