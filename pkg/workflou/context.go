package workflou

type ContextKey string

var (
	SessionKey   ContextKey = "session"
	UserKey      ContextKey = "user"
	RequestIdKey ContextKey = "requestID"
)
