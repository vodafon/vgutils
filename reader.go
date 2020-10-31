package vgutils

import (
	"bufio"
	"io"
)

func LinesToChan(r io.Reader, c chan<- string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		c <- scanner.Text()
	}
	close(c)
}
