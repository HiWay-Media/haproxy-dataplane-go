package haproxy


import (
    "testing"
    "log"
    )

func TestNewHaproxyClient(t *testing.T){

    got := NewHaproxyClient()
    log.Println("got haproxyClient", got)
}