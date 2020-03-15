package pepo

type EventPayload struct {
	ID        string `bson:"_id" json:"id"`
	Topic     string `json:"topic"`
	CreatedAt int64  `json:"created_at"`
	WebhookID string `json:"webhook_id"`
	Version   string `json:"version"`
	Data      Data   `json:"data"`
}

type Data struct {
	Users    map[string]User `json:"users"`
	Activity Activity        `json:"activity"`
}

type User struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	ProfileImage       *string `json:"profile_image"`
	TokenholderAddress *string `json:"tokenholder_address"`
	TwitterHandle      *string `json:"twitter_handle"`
	GithubLogin        *string `json:"github_login"`
}

type Activity struct {
	Kind    string `json:"kind"`
	ActorID int64  `json:"actor_id"`
	Video   Video  `json:"video"`
}

type Video struct {
	ID                      int64    `json:"id"`
	CreatorID               int64    `json:"creator_id"`
	URL                     string   `json:"url"`
	VideoURL                string   `json:"video_url"`
	TotalContributors       int64    `json:"total_contributors"`
	TotalContributionAmount string   `json:"total_contribution_amount"`
	Description             *string  `json:"description"`
	PosterImage             *string  `json:"poster_image"`
	Status                  string   `json:"status"`
	Tags                    []string `json:"tags"`
}
