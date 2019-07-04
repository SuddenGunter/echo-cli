package tokenstore

// TokenStore defines interface for auth token persistent storage
type TokenStore interface {

	// Save is main and only method of TokenStore that takes token string and save it somewhere
	Save(token string) error
}
