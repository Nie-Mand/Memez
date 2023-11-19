package config

import (
	"log"
	"time"
)

func RemoteLog(msg string) {
	log.Println("[*] " + msg)
	time.Sleep(400 * time.Millisecond)
}