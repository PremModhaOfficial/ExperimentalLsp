package analysis

import (
	"ExperimentalLsp/lsp"
	"fmt"
	"strings"
)

type State struct {
	Documents map[string]string
}

func (s State) TextDocumentCodeAction(id int, uri string, params lsp.TextDocumentCodeActionParams) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[uri]
	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChage := map[string][]lsp.TextEdit{}
			replaceChage[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS Code with Neovim",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChage},
			})
			censor := map[string][]lsp.TextEdit{}
			censor[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "** ****",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "censor `VS Code`",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChage},
			})
		}
	}
	responce := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
	return responce
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}
func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponce {
	document := s.Documents[uri]
	responce := lsp.HoverResponce{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Chracters: %d", uri, len(document)),
		},
	}

	return responce
}
func (s State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponce {
	return lsp.DefinitionResponce{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
