package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"time"
)

var id_a int

func cliente() int {
	var id, count int
	var cadena string
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	err = gob.NewDecoder(c).Decode(&cadena)
	if err != nil {
		fmt.Println(err)
		return 0
	} else {
		//fmt.Println(cadena)
		id_s := cadena[0]
		count_s := cadena[1:]
		id, err = strconv.Atoi(string(id_s))
		count, err = strconv.Atoi(string(count_s))
		id_a = id
		//fmt.Println(id)
		//fmt.Println(count)
		go Proceso(id, count)
	}
	c.Close()
	return id

}

func clienteend(i int) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(i)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()

}

func Proceso(id int, i int) {
	for {
		fmt.Printf("Proceso %d: %d \n", id, i)
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go cliente()

	var input string
	fmt.Scanln(&input)

	clienteend(id_a)

}
