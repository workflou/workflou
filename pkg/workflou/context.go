package workflou

type ContextKey string

var (
	SessionKey   ContextKey = "session"
	UserKey      ContextKey = "user"
	TeamKey      ContextKey = "team"
	TeamsKey     ContextKey = "teams"
	RequestIdKey ContextKey = "requestID"
)
