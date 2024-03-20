package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type Respose struct {
	RPC    string `json:"jsonrpc"`
	ID     *int   `json:"id,omitempty"`
	Method string `json:"method"`
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
