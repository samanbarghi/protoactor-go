package actor

type localProcess struct {
	mailbox Mailbox
	dead    bool
}

func (ref *localProcess) SendUserMessage(pid *PID, message interface{}, sender *PID) {
	if sender != nil {
		ref.mailbox.PostUserMessage(&messageSender{Message: message, Sender: sender})
	} else {
		ref.mailbox.PostUserMessage(message)
	}
}

func (ref *localProcess) SendSystemMessage(pid *PID, message SystemMessage) {
	ref.mailbox.PostSystemMessage(message)
}

func (ref *localProcess) Stop(pid *PID) {
	ref.dead = true
	ref.SendSystemMessage(pid, stopMessage)
}
