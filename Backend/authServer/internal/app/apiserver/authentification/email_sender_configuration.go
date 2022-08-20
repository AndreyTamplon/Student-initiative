package authentification

type EmailSenderConfiguration struct {
	APIKey       string `toml:"api_key"`
	FromName     string `toml:"from_name"`
	FromEmail    string `toml:"from_email"`
	Subject      string `toml:"subject"`
	ReplyToName  string `toml:"reply_to_name"`
	ReplyToEmail string `toml:"reply_to_email"`
}
