package domain

type Topic struct {
	// deprecated
	Name            string   `json:"name,omitempty"`
	PrincipalTopic  string   `json:"principalTopic,omitempty"`
	SecondaryTopics []string `json:"secondaryTopics,omitempty"`
}
