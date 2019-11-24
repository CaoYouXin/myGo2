package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

type protobufexample map[string]interface{}

// func (data *protobufexample) Marshal() ([]byte, error) {
//   return json.Marshal(data)
// }

// struct {
// 	Lang string `protobuf:"bytes,1,opt,name=lang" json:"lang"`
// }

func (data *protobufexample) Reset() {}
func (data *protobufexample) String() string {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("protobuf: %v\n", err)
		return ""
	}
	return string(bytes)
}
func (data *protobufexample) ProtoMessage() {}

func main() {
	data := protobufexample(map[string]interface{}{
		"Lang": "GO语言",
		"Tag":  "<br/>",
	})
	// data := protobufexample{Lang: "GO语言"}
	bytes, _ := proto.Marshal(&data)
	fmt.Println(string(bytes), len(bytes))
}
