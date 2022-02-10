package greetings

type Conversation interface {
	SayHello() string
	SayGoodbye() string
}

type ErrBoring struct{}

func (e ErrBoring) Error() string {
	return "come on, that's boring"
}

type ErrEmptyName struct{}

func (e ErrEmptyName) Error() string {
	return "empty name"
}
