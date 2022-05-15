package chesgomailing

type Config struct {
	From     string
	Password string
	Host     string
	Port     string
}

type DataToSend struct {
	Format  string
	Subject string
	Body    string
	To      []string
}
