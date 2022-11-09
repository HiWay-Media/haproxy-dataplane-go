package haproxy

import (
	"log"
	"testing"
)

func TestNewHaproxyClient(t *testing.T) {

	got, err := NewHaproxyClient("localhost-haproxy", "admin", "password", true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("got haproxyClient", got)
}
