# Linking channels in Go

Demonstrating how to link two channels in Go and pass a message

## Functions

**importToChannel** <br/>
Takes a channel and a string message as arguments and passes the message into channel
```
func importToChannel(channelInput chan<- string, message string)  {
	fmt.Println("Importing message to channel")
	channelInput <- message
}
```

**passToChannel** <br/>
Takes two channels as arguments and passes information from channelInput to channelOutput
```
func passToChannel(channelInput <-chan string, channelOutput chan<-string)  {
	fmt.Println("Linking channels")
	tmp := <-channelInput
	channelOutput <- tmp
}
```

**exportFromChannel** <br/>
Takes a channel as an arguments and returns a string (the message inside the channel)
```
func exportFromChannel(channel <-chan string ) string {
	fmt.Println("Exporting message from channel")
	result := <- channel
	return result
}
```
