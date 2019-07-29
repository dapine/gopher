package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/dapine/gopher/item"
)

func main() {
	ln, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	line, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		conn.Close()
	}

	// XXX
	items := []item.Item{
		item.Item{ItemType: '0', Name: "About internet Gopher", Selector: "Stuff:About us", Hostname: "rawBits.micro.umn.edu", Port: 70},
		item.Item{ItemType: '1', Name: "Around University of Minnesota", Selector: "Z,5692,AUM", Hostname: "underdog.micro.umn.edu", Port: 70},
		item.Item{ItemType: '1', Name: "Courses, Schedules, Calendars", Selector: "", Hostname: "events.ais.umn.edu", Port: 9120},
		item.Item{ItemType: '1', Name: "Student-Staff Directories", Selector: "", Hostname: "uinfo.ais.umn.edu", Port: 70},
		item.Item{ItemType: '1', Name: "Departmental Publications", Selector: "Stuff:DP:", Hostname: "rawBits.micro.umn.edu", Port: 70},
	}

	its := item.Select(items, string(line))
	fmt.Println(item.Format(its))

	conn.Close()
}
