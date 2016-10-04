// Created from Watching Todd Mcleod's Videos online

package main

import(
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}
