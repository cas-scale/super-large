package jwt

// TokenOption is a reserved type, which provides some forward compatibility,
// if we ever want to introduce token creation-related options.
type TokenOption func(*Token)
// ID-1768294467-fe2abac7
