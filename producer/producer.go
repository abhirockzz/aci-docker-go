package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-redis/redis"
)

func main() {

	host := os.Getenv("REDIS_HOST")
	if host == "" {
		log.Fatal("please set REDIS_HOST env variable")
	}
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		log.Fatal("please set REDIS_PASSWORD env variable")
	}
	numMsgsStr := os.Getenv("NUM_MESSAGES")
	if numMsgsStr == "" {
		numMsgsStr = "1000"
	}
	numMsgs, err := strconv.Atoi(numMsgsStr)

	if err != nil {
		log.Fatal("please specify valid value for number of messages you want to send")
	}
	listName := os.Getenv("REDIS_LIST")
	if listName == "" {
		log.Fatal("please set REDIS_LIST env variable")
	}
	fmt.Println("connecting to Redis", host)

	client := redis.NewClient(&redis.Options{
		Addr:      host,
		Password:  password,
		TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
	})

	_, err = client.Ping().Result()

	if err != nil {
		fmt.Println("failed to connect to Redis", err)
		os.Exit(1)
	}
	fmt.Println("successfully connected to Redis", host)
	defer func() {
		err := client.Close()
		if err != nil {
			fmt.Println("failed to close client conn ", err)
			return
		}
		fmt.Println("closed client connection")

	}()
	fmt.Println("Sending " + numMsgsStr + " messages to Redis list " + listName)

	go func() {
		for i := 1; i <= numMsgs; i++ {
			msg := "message-" + strconv.Itoa(i)
			err := client.LPush(listName, msg).Err()
			if err != nil {
				log.Fatal("unable to send data to Redis list", err)
			}
			fmt.Println("sent message", msg)
			time.Sleep(1 * time.Second)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("press ctrl+c to exit...")
	<-exit
	fmt.Println("program exited")
}
