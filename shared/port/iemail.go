package port

type IEmail interface {
	SetAttachment(attachments interface{}) IEmail
	SetFrom(from string, fromName ...string) IEmail
	SetTo(to string, toName ...string) IEmail
	SetBody(body string) IEmail
	SetSubject(subject string) IEmail
	Send() error
}
