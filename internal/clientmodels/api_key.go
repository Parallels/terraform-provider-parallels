package clientmodels

type APIKeyRequest struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type APIKeyResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Revoked   bool   `json:"revoked"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	RevokedAt string `json:"revoked_at"`
}
