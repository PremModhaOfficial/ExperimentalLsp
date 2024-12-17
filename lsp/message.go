package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Mathod string `json:"method"`

	// DONE: specify params in all the Requests types
	// params
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	// Results
}
type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
	// Results
}
