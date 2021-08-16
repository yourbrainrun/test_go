package data

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Area struct {
	Name     string
	Isp      string
	Ip       string
	DataChan chan string
}

type Dealer interface {
	GetData() map[string]string
}

func NewArea() *Area {
	return &Area{
		DataChan: make(chan string),
	}
}

func (area *Area) GetData(name string) {
	defer func() {
		close(area.DataChan)
	}()

	fd, err := os.Open(name)
	if err == nil {
		bf := bufio.NewReader(fd)
		for {
			byteSlice, _, c := bf.ReadLine()
			if c == io.EOF {
				break
			}
			str := string(byteSlice)
			strSlice := strings.Split(str, "\t")

			area.DataChan <- strSlice[2]
		}
	}
}
