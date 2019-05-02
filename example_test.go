package csvdict_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kurankat/csvdict"
)

func ExampleDictReader() {
	in := `name, age, occupation
"Mark Smith",33,"Jack of all trades"
"Douglas Adams",42,Writer
Methuselah,969,Patriarch
`

	r, err := csvdict.NewDictReader(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	headers := r.GetHeaderRow()

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v:%v %v:%v %v:%v\n", headers[0], record[headers[0]],
			headers[1], record[headers[1]],
			headers[2], record[headers[2]])
	}
	// Output:
	// name:Mark Smith  age:33  occupation:Jack of all trades
	// name:Douglas Adams  age:42  occupation:Writer
	// name:Methuselah  age:969  occupation:Patriarch
}

func ExampleDictReader_ReadAll() {
	in := `name, age, occupation
"Mark Smith",33,"Jack of all trades"
"Douglas Adams",42,Writer
Methuselah,969,Patriarch
`

	r, err := csvdict.NewDictReader(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	headers := r.GetHeaderRow()

	for _, record := range records {
		fmt.Printf("%v:%v %v:%v %v:%v\n", headers[0], record[headers[0]],
			headers[1], record[headers[1]],
			headers[2], record[headers[2]])
	}

	// Output:
	// name:Mark Smith  age:33  occupation:Jack of all trades
	// name:Douglas Adams  age:42  occupation:Writer
	// name:Methuselah  age:969  occupation:Patriarch
}

func ExampleDictWriter() {
	headers := []string{"name", "age", "occupation"}
	records := []map[string]string{
		{"name": "Mark Smith", "age": "33", "occupation": "Jack of all trades"},
		{"name": "Douglas Adams", "age": "42", "occupation": "Writer"},
		{"name": "Methuselah", "age": "969", "occupation": "Patriarch"},
	}

	w := csvdict.NewDictWriter(os.Stdout, headers)

	w.WriteHeaders()

	for _, record := range records {
		err := w.Write(record)
		if err != nil {
			panic(err)
		}
	}

	w.Flush()
	// Output:
	// name,age,occupation
	// Mark Smith,33,Jack of all trades
	// Douglas Adams,42,Writer
	// Methuselah,969,Patriarch
}

func ExampleDictWriter_WriteAll() {
	headers := []string{"name", "age", "occupation"}
	records := []map[string]string{
		{"name": "Mark Smith", "age": "33", "occupation": "Jack of all trades"},
		{"name": "Douglas Adams", "age": "42", "occupation": "Writer"},
		{"name": "Methuselah", "age": "969", "occupation": "Patriarch"},
	}

	w := csvdict.NewDictWriter(os.Stdout, headers)

	err := w.WriteAll(records)
	if err != nil {
		panic(err)
	}
	// Output:
	// name,age,occupation
	// Mark Smith,33,Jack of all trades
	// Douglas Adams,42,Writer
	// Methuselah,969,Patriarch
}
