package main

import (
	"os"
	"strings"

	"fmt"

	"github.com/alecthomas/kingpin"
)

var (
	app      = kingpin.New("chat", "A command-line chat application.")
	debug    = app.Flag("debug", "Enable debug mode.").Bool()
	serverIP = app.Flag("server", "Server address.").Default("127.0.0.1").IP()

	register = app.Command("register", "Register a new user.")
	//registerNick = register.Arg("nick", "Nickname for user.").Required().String()
	registerNick = register.Arg("nick", "Nickname for user.").String()
	registerName = register.Arg("name", "Name of user.").String()

	post        = app.Command("post", "Post a message to a channel.")
	postImage   = post.Flag("image", "Image to post.").File()
	postChannel = post.Arg("channel", "Channel to post to.").String()
	postText    = post.Arg("text", "Text to post.").Strings()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	case register.FullCommand():
		fmt.Println(*registerNick)
		fmt.Println(*serverIP)

	// Post message
	case post.FullCommand():
		if *postImage != nil {
		}
		text := strings.Join(*postText, " ")
		println("Post:", text)
	default:
		{
			fmt.Println(*serverIP)
			fmt.Println(*debug)
		}
	}
}
