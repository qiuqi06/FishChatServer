package main

import (
	"encoding/binary"
	"net"
)

func main() {
	println("sss")

	conn, _ := net.Dial("tcp", "localhost:17000")
	var bytes [4]byte
	str:="{\"CmdName\":\"REQ_MSG_SERVER\",\"Args\":[\"ssss\"]}"
	pkgLen := uint32(len(str))
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	conn.Write(bytes[:])
	i, err := conn.Write([]byte(str))

	println(i)
	if err != nil {
		println("error")
	}

	buf:=make([]byte, 1024)

	n, _ := conn.Read(buf)

	line:=string(buf[4:n])
	println(line)

	conn.Close()

}
