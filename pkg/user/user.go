package user

// Provides info about current user
type CurrentUserInfoProvider interface {
	// Returns current user's token
	GetToken() (string, error)
}
