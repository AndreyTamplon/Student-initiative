package authentification

type EmailSender struct {
	configuration *EmailSenderConfiguration
}

func NewEmailSender(configuration *EmailSenderConfiguration) *EmailSender {
	es := &EmailSender{
		configuration: configuration,
	}
	return es
}

func (es *EmailSender) Send(from *Address, to []*Address, message string, subject string, replyTo *Address) error {
	emailMessage := Message{
		Sender:      from,
		To:          to,
		TextContent: message,
		Subject:     subject,
		ReplyTo:     replyTo,
	}
	err := emailMessage.Send(es.configuration.APIKey)
	if err != nil {
		return err
	}
	return nil
}

func (es *EmailSender) DefaultSend(message string, to []*Address) error {
	return es.Send(&Address{
		Name:  es.configuration.FromName,
		Email: es.configuration.FromEmail,
	}, to, message, es.configuration.Subject, &Address{
		Name:  es.configuration.ReplyToName,
		Email: es.configuration.ReplyToEmail,
	})
}
