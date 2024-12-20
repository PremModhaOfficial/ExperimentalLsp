package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}
type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
type ServerCapabilities struct {
	TextDocumentSync   int  `json:"textDocumentSync"`
	HoverProvider      bool `json:"hoverProvider"`
	DefinitionProvide  bool `json:"definitionProvider"`
	CodeActionProvider bool `json:"codeActionProvider"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:   1,
				HoverProvider:      true,
				DefinitionProvide:  true,
				CodeActionProvider: true,
			},
			ServerInfo: ServerInfo{
				Name:    "Experimental-lsp",
				Version: "0.0.0.0.beta-1.final",
			},
		},
	}
}
