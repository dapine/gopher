package item

import (
	"fmt"
	"net"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

const (
	file            = '0'
	directory       = '1'
	cso             = '2'
	err             = '3'
	binHexMac       = '4'
	binDOS          = '5'
	uuencoded       = '6'
	searchServer    = '7'
	telnet          = '8'
	binary          = '9'
	redundantServer = '+'
	telnet3270      = 'T'
	gif             = 'g'
	image           = 'I'
)

type Item struct {
	ItemType rune
	Name     string
	Selector string
	Hostname string
	Ip       net.IP
	Port     int
}

func (i Item) String() string {
	return fmt.Sprintf("%c%s\t%s\t%s\t%d\r\n", i.ItemType, i.Name, i.Selector, i.Hostname, i.Port)
}

func Select(items []Item, selector string) []Item {
	its := []Item{}
	switch selector {
	case "\r\n":
		return items
	default:
		for _, i := range items {
			if fuzzy.MatchFold(strings.TrimRight(selector, "\r\n"), i.Selector) {
				its = append(its, i)
			}
		}
	}

	return its
}

func Format(items []Item) string {
	str := ""
	for _, i := range items {

		str += i.String()

	}

	str += ".\r\n"

	return str
}
