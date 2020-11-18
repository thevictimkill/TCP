package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"time"
)

type Persona struct {
	Nombre string
	Email  []string
}

type Proceso struct {
	Intlist []int
	Count   int
}

func (pro *Proceso) Pop() int {
	if len(pro.Intlist) == 0 {
		return 0
	}
	aux := pro.Intlist[0]
	pro.Intlist = append(pro.Intlist[:0], pro.Intlist[1:]...)
	return aux
}

func (pro *Proceso) Add(i int) {
	pro.Intlist = append(pro.Intlist, i)
}
func (pro *Proceso) Autosuma() {
	pro.Count++
}
func (pro *Proceso) Getlast() int {
	return pro.Intlist[0]
}

func servidor(pro *Proceso) {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c, pro)
	}
}

func handleClient(c net.Conn, pro *Proceso) {
	var id_r int
	id := pro.Pop()
	cadena := strconv.Itoa(id) + strconv.Itoa(pro.Count+1)
	err := gob.NewEncoder(c).Encode(cadena)
	if err != nil {
		fmt.Println(err)
	}

	err = gob.NewDecoder(c).Decode(&id_r)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		x := pro.Getlast()
		pro.Add(id_r)
		pro.Intlist = append([]int{x - 1}, pro.Intlist...)
	}

}

func proceso(pro *Proceso) {
	count := 1
	for {
		fmt.Println("------------------------")
		for _, ele := range pro.Intlist {
			fmt.Println("Proceso", ele, ":", count)

		}
		time.Sleep(time.Millisecond * 500)
		pro.Autosuma()
		count++

	}
}

func main() {
	intlist := &Proceso{
		[]int{1, 2, 3, 4, 5},
		0,
	}

	go servidor(intlist)
	go proceso(intlist)

	var input string
	fmt.Scanln(&input)

}
