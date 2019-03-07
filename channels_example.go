package main

import "fmt"

/**
	Importing a message to a channel
 */
func importToChannel(channelInput chan<- string, message string)  {
	fmt.Println("Importing message to channel")
	channelInput <- message
}

/**
	Pass the information holding inside a channel to a second channel
 */
func passToChannel(channelInput <-chan string, channelOutput chan<-string)  {
	fmt.Println("Linking channels")
	tmp := <-channelInput
	channelOutput <- tmp
}

/**
	Exporting a message from a channel
 */
func exportFromChannel(channel <-chan string ) string {
	fmt.Println("Exporting message from channel")
	result := <- channel
	return result
}

func main()  {
	channel_1 := make(chan string, 1)
	channel_2 := make(chan string, 1)

	importToChannel(channel_1, "Oh, we started!!")
	passToChannel(channel_1, channel_2)

	fmt.Println("\nMessage: " + exportFromChannel(channel_2))
}
