package main

import (
	"encoding/json"
	"github.com/oikomi/FishChatServer/libnet"
)

func main() {

	buf := make([]byte, 0, 1024)
	user := &User{
		Name: "ss",
		Age:  12,
	}
	outbuffer:=&libnet.OutBuffer{
		Data:buf,
	}

	//outbuffer.Data=outbuffer.Data[2:4]
	json.NewEncoder(outbuffer).Encode(user)
	println(len(buf))

}
type User struct {
	Name string
	Age int32
}
type Encoder func(buffer *libnet.OutBuffer) error
func Json(v interface{}) Encoder {
	return func(buffer *libnet.OutBuffer) error {
		return json.NewEncoder(buffer).Encode(v)
	}
}
