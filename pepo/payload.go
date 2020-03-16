package pepo

type EventPayload struct {
	ID        string `bson:"_id" json:"id"`
	Topic     string `bson:"topic" json:"topic"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	WebhookID string `bson:"webhook_id" json:"webhook_id"`
	Version   string `bson:"version" json:"version"`
	Data      Data   `bson:"data" json:"data"`
}

type Data struct {
	Users    map[string]User `bson:"users" json:"users"`
	Activity Activity        `bson:"activity" json:"activity"`
}

type User struct {
	ID                 string  `bson:"id" json:"id"`
	Name               string  `bson:"name" json:"name"`
	ProfileImage       *string `bson:"profile_image" json:"profile_image"`
	TokenholderAddress *string `bson:"tokenholder_address" json:"tokenholder_address"`
	TwitterHandle      *string `bson:"twitter_handle" json:"twitter_handle"`
	GithubLogin        *string `bson:"github_login" json:"github_login"`
}

type Activity struct {
	Kind    string `bson:"kind" json:"kind"`
	ActorID int64  `bson:"actor_id" json:"actor_id"`
	Video   Video  `bson:"video" json:"video"`
}

type Video struct {
	ID                      int64    `bson:"id" json:"id"`
	CreatorID               int64    `bson:"creator_id" json:"creator_id"`
	URL                     string   `bson:"url" json:"url"`
	VideoURL                string   `bson:"video_url" json:"video_url"`
	TotalContributors       int64    `bson:"total_contributors" json:"total_contributors"`
	TotalContributionAmount string   `bson:"total_contribution_amount" json:"total_contribution_amount"`
	Description             *string  `bson:"description" json:"description"`
	PosterImage             *string  `bson:"poster_image" json:"poster_image"`
	Status                  string   `bson:"status" json:"status"`
	Tags                    []string `bson:"tags" json:"tags"`
}
