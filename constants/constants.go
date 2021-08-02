package constants

const (
	MONGO_NO_DOC = "mongo: no documents in result"
	USER_COLL    = "users"
)

type markers int

const (
	opening = iota
	closing
	max
	min

)