package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// body of url body(interact with frontend)
type ShortUrlBody struct {
	LongUrl string `json:"long_url"`
}

// structure of url stored in the Database
type URLDB struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UrlCode   string             `bson:"url_code"`
	LongUrl   string             `bson:"long_url"`
	ShortUrl  string             `bson:"short_url"`
	CreatedAt int64              `bson:"created_at"`
	ExpiredAt int64              `bson:"expired_at"`
}
