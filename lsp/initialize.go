package lsp

type InitiailizeRequest struct {
	Request
	Params InitiailizeRequestParams `json:"params"`
}

type InitiailizeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// tons of other fields
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
