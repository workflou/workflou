package workflou

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	CurrentTeam  *Team
	Teams        []*Team
}
