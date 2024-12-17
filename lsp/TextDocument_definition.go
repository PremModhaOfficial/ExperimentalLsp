package lsp

type DefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

type DefinitionParams struct {
	TextDocumentPositionParams
}

type DefinitionResponce struct {
	Response
	Result Location `json:"result"`
}
