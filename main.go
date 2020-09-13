package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"

	"github.com/golang/protobuf/proto"

	simplepb "github.com/shubhamjain2908/protobuf-example-go/src/simple"
)

func main() {
	sm := doSimple()

	readAndWriteDeme(sm)
	smAsString := toJSON(sm)

	fmt.Println("Json string =>", smAsString)
}

// returning reference to SimpleMessage (pass by reference)
func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple Message",
		SampleList: []int32{1, 2, 7, 8},
	}

	sm.Name = "I renamed you"

	fmt.Println("Id => ", sm.GetId())
	fmt.Println("Name => ", sm.GetName())

	return &sm
}

func readAndWriteDeme(sm proto.Message) {
	writeToFile("Simple.bin", sm)

	// creating a simple struct
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("Simple.bin", sm2)
	fmt.Println("Read the contents", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Coudn't put the bytes into the protocol buffer struct", err)
		return err2
	}

	return nil
}

func toJSON(pb proto.Message) string {
	marshler := jsonpb.Marshaler{}
	out, err := marshler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}
