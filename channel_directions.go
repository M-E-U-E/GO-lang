package main

import "fmt"

func ping(pings chan<- string, msg string) {
    pings <- msg // only sending the perameter
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1) // Create a buffered channel for sending messages.
    pongs := make(chan string, 1) // Create another buffered channel for sending messages.
    ping(pings, "passed message") // Call the `ping` function to send a message into the `pings` channel.
    pong(pings, pongs)            // Call the `pong` function to pass the message from `pings` to `pongs`.
    fmt.Println(<-pongs)          // Receive and print the message from the `pongs` channel.
}
