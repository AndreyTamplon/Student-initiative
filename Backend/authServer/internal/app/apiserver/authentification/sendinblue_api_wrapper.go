package authentification

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Address struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Attachment struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

func InlineAttachment(name string, in io.Reader) (ret Attachment, err error) {
	ret.Name = name

	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)

	_, err = io.Copy(enc, in)
	if err == nil {
		ret.Content = buf.String()
	}
	return
}

type Message struct {
	Sender      *Address          `json:"sender"`
	To          []*Address        `json:"to,omitempty"`
	Bcc         []*Address        `json:"bcc,omitempty"`
	Cc          []*Address        `json:"cc,omitempty"`
	HTMLContent string            `json:"htmlContent,omitempty"`
	TextContent string            `json:"textContent,omitempty"`
	Subject     string            `json:"subject,omitempty"`
	ReplyTo     *Address          `json:"replyTo,omitempty"`
	Attachments []*Attachment     `json:"attachment,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	TemplateID  int64             `json:"templateId,omitempty"`
	Params      map[string]string `json:"params,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
}

func (m *Message) Send(apiKey string) error {
	url := "https://api.sendinblue.com/v3/smtp/email"

	data, err := json.Marshal(m)
	if err != nil {
		return errors.New("failed to encode message: " + err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("api-key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return errors.New("failed to transmit message: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		log.Println(resp)
		return fmt.Errorf("send failed: %d %s", resp.StatusCode, resp.Status)
	}
	return nil
}
