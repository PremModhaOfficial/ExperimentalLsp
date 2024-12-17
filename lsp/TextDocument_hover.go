package lsp

type HoverRequest struct {
	Request
	Params HoverParams `json:"params"`
}

type HoverParams struct {
	TextDocumentPositionParams
}

type HoverResponce struct {
	Response
	Result HoverResult `json:"result"`
}
type HoverResult struct {
	Contents string `json:"contents"`
}
