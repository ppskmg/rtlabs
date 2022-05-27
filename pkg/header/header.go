package header

import "net/http"

type Content struct {
	Key        string
	ValuePlain string
	ValueJson  string
}
type UserAgent struct {
	Key   string
	Value string
}
type SessionID struct {
	Key string
}
type Header struct {
	Content
	UserAgent
	SessionID
}

var (
	Headers = &Header{
		Content: Content{
			Key:        "Content-Type",
			ValuePlain: "text/plain",
			ValueJson:  "application/json",
		},
		UserAgent: UserAgent{
			Key:   "User-Agent",
			Value: "testing user agent",
		},
		SessionID: SessionID{
			Key: "Session-ID",
		},
	}
)

func SetTestUserAgent(req *http.Request) {
	req.Header.Set(Headers.UserAgent.Key, Headers.UserAgent.Value)
}
func SetTestSessionID(req *http.Request) {
	req.Header.Set(Headers.SessionID.Key, "incorrect test session id")
}
