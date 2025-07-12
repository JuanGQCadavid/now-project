package domain

import "fmt"

type FileKind string

const (
	ProfileType FileKind = "profile"
	EventType   FileKind = "event"
)

type FileScope string

const (
	EventScope FileScope = "event"
	ChatScope  FileScope = "chat"
	DateScope  FileScope = "date"
)

type FileType string

const (
	ImageType FileScope = "WebP"
	VideoType FileScope = "H.264"
	VoiceType FileScope = "opus"
	PDF       FileScope = "pdf"
)

type ControlAccess struct {
	EventId string `json:"eventId,omitempty"`
	DateId  string `json:"dateId,omitempty"`
}

type FileMetadata struct {
	Kind          FileKind      `json:"kind,omitempty"`
	Scope         FileScope     `json:"scope,omitempty"`
	ControlAccess ControlAccess `json:"controlAccess"`
	ForcePublic   bool          `json:"forcePublic,omitempty"`
	Type          FileType      `json:"type,omitempty"`
}

type PresignedURL struct {
	URL     string `json:"url,omitempty"`
	Headers map[string]string
	Method  string
}

func (pre *PresignedURL) ToString() string {
	return fmt.Sprintf("Presigned: \n\t URL:%s \n\t Method:%s \n\t Headers: %v\n", pre.URL, pre.Method, pre.Headers)
}
