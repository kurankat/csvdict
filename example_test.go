package csvdict_test

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/kurankat/csvdict"
)

func exampleDictReader() {
	in := `name, age, occupation
"Mark Smith",33,"Jack of all trades"
"Douglas Adams",42,Writer
Methuselah,969,Patriarch
`

	r := csvdict.NewDictReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
	// Output:
	// map[name:Mark Smith  age:33  occupation:Jack of all trades]
	// map[ age:42  occupation:Writer name:Douglas Adams]
	// map[name:Methuselah  age:969  occupation:Patriarch]
}

func ExamplDictReader_ReadAll() {
	in := `name, age, occupation
"Mark Smith",33,"Jack of all trades"
"Douglas Adams",42,Writer
Methuselah,969,Patriarch
`

	r := csvdict.NewDictReader(strings.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
	// Output:
	// [map[name:Mark Smith  age:33  occupation:Jack of all trades] map[name:Douglas Adams  age:42  occupation:Writer] map[name:Methuselah  age:969  occupation:Patriarch]]
}
