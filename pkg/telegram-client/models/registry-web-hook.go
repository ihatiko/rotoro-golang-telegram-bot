package models

type RegistryWebHookRequest struct {
	Url string `json:"url"`
}

type RegistryWebHookResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}
