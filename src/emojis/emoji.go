package emojis

type Emoji struct {
	Emoji          string   `json:"emoji,omitempty"`
	Description    string   `json:"description,omitempty"`
	Category       string   `json:"category,omitempty"`
	Aliases        []string `json:"aliases,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	UnicodeVersion string   `json:"unicode_version,omitempty"`
	IosVersion     string   `json:"ios_version,omitempty"`
	SkinTones      bool     `json:"skin_tones,omitempty"`
}
