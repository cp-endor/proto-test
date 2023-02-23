package main

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	fmt.Println("hello")
	jsonObj := `{"published":"1988-10-01T04:00:00.000"}`

	strct := &Record{}

	if err := protojson.Unmarshal([]byte(jsonObj), strct); err != nil {
		fmt.Println(err)
	}

}
