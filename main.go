package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/haroflow/protobuf-example/protos/addressbook"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("ProtoBuffer example")

	p := &addressbook.Person{
		Id:    1,
		Name:  "John Doe",
		Email: "john.doe@gmail.com",
		Phones: []*addressbook.Person_PhoneNumber{
			{
				Number: "1111-2222",
				Type:   addressbook.Person_HOME,
			},
		},
	}
	fmt.Println("Person 1:", p)

	p2 := &addressbook.Person{
		Id:    2,
		Name:  "Mary Jane",
		Email: "mary.jane@aol.com",
		Phones: []*addressbook.Person_PhoneNumber{
			{
				Number: "3333-4444",
				Type:   addressbook.Person_HOME,
			},
			{
				Number: "95555-6666",
				Type:   addressbook.Person_MOBILE,
			},
		},
	}
	fmt.Println("Person 2:", p2)

	book := &addressbook.AddressBook{
		People: []*addressbook.Person{
			p,
			p2,
		},
	}

	fmt.Println()
	fmt.Println("Writing address book to file output.bin:")
	fmt.Println(book)

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("output.bin", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	fmt.Println()
	fmt.Println("Reading address book from file output.bin")

	loadedBook := &addressbook.AddressBook{}
	bytes, err := ioutil.ReadFile("output.bin")
	if err != nil {
		log.Fatalln("Failed to read from file:", err)
	}
	if err := proto.Unmarshal(bytes, loadedBook); err != nil {
		log.Fatalln("Failed to decode address book:", err)
	}
	fmt.Println(loadedBook)
}
