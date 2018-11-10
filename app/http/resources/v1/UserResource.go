package resources

type UserResource struct {
	Id               string `json:"id,omitempty"`
	Email            string `json:"email,omitempty"`
	FullName         string `json:"full_name,omitempty"`
	Timestamp        string `json:"created_at,omitempty"`
	UpdatedTimestamp string `json:"updated_at,omitempty"`
}

type UserCollection []UserResource
