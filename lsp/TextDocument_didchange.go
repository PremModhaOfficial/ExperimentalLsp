package lsp

type TextDocumentDidChageNotification struct {
	Notification
	Params DidChageTextDocumentParams `json:"params"`
}

type DidChageTextDocumentParams struct {
	TextDocument   VersionTextDocumentIdentifier    `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

/*
Event decribing a chage to a text document.
if only provided text it is considered the full document
*/
type TextDocumentContentChangeEvent struct {
	// the new text of whole document.
	Text string `json:"text"`
}
