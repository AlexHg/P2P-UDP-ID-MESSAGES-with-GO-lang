package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"reflect"
	"strconv"
	"unicode"
)

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func main() {
	PORT, err := strconv.Atoi(os.Args[1])

	fmt.Println(reflect.TypeOf(err))

	ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: PORT, Zone: ""})
	defer ServerConn.Close()
	buf := make([]byte, 1024)
	go client()
	var msgRcv = 0
	var lastRcv = ""
	for {
		n, addr, _ := ServerConn.ReadFromUDP(buf)

		if string(buf[0:n]) == "0" {
			fmt.Println("0 Recived from ", addr, "Sending stadistics")

			PORTres, err := strconv.Atoi(os.Args[2])
			fmt.Println(reflect.TypeOf(err))

			Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: PORTres, Zone: ""})
			defer Conn.Close()

			Conn.Write([]byte("\nProceso terminado\n"))
			Conn.Write([]byte("\nESTADISTICAS\n"))

			lastRint, err := strconv.Atoi(lastRcv)
			fmt.Println(reflect.TypeOf(err))

			fmt.Println(msgRcv, "/", lastRint)

			stadis := fmt.Sprintf("%d%s%d%s%d%s", msgRcv, "/", lastRint, "% Recibidos, ", lastRint-msgRcv, " mensajes perdidos")
			Conn.Write([]byte(stadis))
			msgRcv = 0

		} else {

			fmt.Println("\nReceived ", string(buf[0:n]), " from ", addr)
			if isInt(string(buf[0:n])) {
				lastRcv = string(buf[0:n])
				msgRcv++
			}

		}
	}
}

func client() {
	PORT, err := strconv.Atoi(os.Args[2])
	SENDER, err := strconv.Atoi(os.Args[3])
	fmt.Println(reflect.TypeOf(err))

	for {
		fmt.Print("\nPress 'Enter' to send...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: PORT, Zone: ""})
		defer Conn.Close()
		for s := 1; s <= SENDER; s++ {
			Conn.Write([]byte(strconv.Itoa(s)))
			if s == SENDER {
				Conn.Write([]byte(strconv.Itoa(0)))
			}
		}

		fmt.Print("\nAll data sended...\n")
	}
}
