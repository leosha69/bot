package cmd

func Cmd() {
	switch update.Message.Command() {
	case "help":
		msg.Text = "I understand"
	case "sayhi":
		msg.Text = "Hi :)"
	case "status":
		msg.Text = "I'm ok."
	default:
		msg.Text = "I don't know"
	}
}
