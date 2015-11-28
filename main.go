package main

import (
	"log"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM|syscall.SOCK_CLOEXEC|syscall.SOCK_NONBLOCK, syscall.IPPROTO_TCP)
	if err != nil {
		log.Fatalln("Cannot create socket:", err)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		log.Fatalln("Cannot setsockopt:", err)
	}

	sa := &syscall.SockaddrInet4{
		Port: 12345,     // Listen on this port number
		Addr: [4]byte{}, // Listen to all IPs
	}

	err = syscall.Bind(fd, sa)
	if err != nil {
		log.Fatalln("Cannot bind to socket:", err)
	}

	err = syscall.Listen(fd, 10)
	if err != nil {
		log.Fatalln("Cannot listen to socket:", err)
	}

	// TODO select

	/*
		go func() {
			x, raddr, err := syscall.Accept(fd)
			if err != nil {
				log.Fatalln("Cannot listen to socket:", err)
			}

			log.Printf("received:")
			log.Printf("%#v", raddr)
			log.Printf("%#v", x)
		}()
	*/

}
