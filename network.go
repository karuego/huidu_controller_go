package main

import (
	"net"
	"fmt"
	"time"
	"bytes"
	"errors"
)

func searchDeviceAsk(deviceCh chan<-string, errCh chan<-error) {
	timeout := 2 * time.Second

	server := fmt.Sprintf("255.255.255.255:%d", REMOTE_PORT)
	serverAddr, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		errCh <- errors.New(fmt.Sprintf("Error resolving address:", err))
		return
	}

	localAddr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		errCh <- errors.New(fmt.Sprintf("Error resolving address:", err))
		return
	}

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		errCh <- errors.New(fmt.Sprintf("Error creating UDP connection:", err))
		return
	}
	// fmt.Println("Listening on local address:", conn.LocalAddr())
	conn.SetReadDeadline(time.Now().Add(timeout))
	defer conn.Close()

	message := []byte{0x05, 0x00, 0x00, 0x01, 0x01, 0x10}
	buffer := make([]byte, 1024)

	stopChan := make(chan struct{})

	go func() {
		time.AfterFunc(timeout, func() {
			close(stopChan)
		})
	}()

	Outer: for {
		select {
		case <-stopChan:
			break Outer

		default:
			n, err := conn.WriteToUDP(message, serverAddr)
			if err != nil {
				fmt.Println("Error sending message:", err)
				errCh <- err
				continue
			}
			// fmt.Println("Broadcast message sent to", serverAddr)

			n, _, err = conn.ReadFromUDP(buffer)
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					fmt.Println("No response received within 5 seconds. Stopping broadcast.")
					break
				}

				fmt.Println("Error receiving response:", err)
				errCh <- err
				continue
			}

			time.Sleep(300 * time.Millisecond)

			if bytes.Equal(buffer[:n], message) {
				continue
			}

			// fmt.Printf("Received response from %s: % x\n", fromAddr.String(), buffer[:n])
			deviceCh <- string(buffer[6:6+15])
		}
	}

	// keys := make([]string, 0, len(devices))

	// for key := range devices {
		// keys = append(keys, key)
	// }

	// devicesCh <- keys

	close(errCh)
	close(deviceCh)
}
