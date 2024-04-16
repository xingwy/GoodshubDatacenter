package message_helper

type MessageHelper interface {
}

type MessageInstance struct {
}

func (i *MessageInstance) Init() *MessageInstance {
	return i
}
