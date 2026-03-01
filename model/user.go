package models

// User represents a social profile user and their associated information. i can add more fields as needed, such as social media links, location, etc. but for now, this should be sufficient for basic profile information.
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	Provider string `json:"provider"` // google or microsoft
	Bio      string `json:"bio"`
}
