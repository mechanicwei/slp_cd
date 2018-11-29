package model

type DeployRepo struct {
	ID            int64    `json:"id"`
	Name          string   `json:"name"`
	GithubUrl     string   `json:"github_url" db:"github_url"`
	WebhookSecret string   `json:"webhook_secret" db:"webhook_secret"`
	Openids       []string `json:"openids"`
	CreatedAt     JsonTime `json:"created_at" db:"created_at"`
}
