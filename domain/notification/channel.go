package notification

type Channel string

const (
	Email Channel = "email"
)

type EmailTask struct {
	Subject string
}

func handleEmailChannel() {

}
