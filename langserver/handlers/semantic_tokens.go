package handlers

import (
	"context"
	"fmt"

	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	lsctx "github.com/hashicorp/terraform-ls/internal/context"
	ilsp "github.com/hashicorp/terraform-ls/internal/lsp"
	lsp "github.com/hashicorp/terraform-ls/internal/protocol"
)

func (lh *logHandler) TextDocumentSemanticTokensFull(ctx context.Context, params lsp.SemanticTokensParams) (lsp.SemanticTokens, error) {
	tks := lsp.SemanticTokens{}

	ds, err := lsctx.DocumentStorage(ctx)
	if err != nil {
		return tks, err
	}

	rmf, err := lsctx.RootModuleFinder(ctx)
	if err != nil {
		return tks, err
	}

	fh := ilsp.FileHandlerFromDocumentURI(params.TextDocument.URI)
	doc, err := ds.GetDocument(fh)
	if err != nil {
		return tks, err
	}

	rm, err := rmf.RootModuleByPath(doc.Dir())
	if err != nil {
		return tks, fmt.Errorf("finding compatible decoder failed: %w", err)
	}

	d, err := rm.Decoder()
	if err != nil {
		return tks, err
	}

	tokens, err := d.TokensInFile(doc.Filename())
	if err != nil {
		return tks, err
	}

	// TODO: encode lang.SemanticTokens to []float64
	tks.Data = EncodeSemanticTokens(tokens)

	return tks, nil
}

type SemanticTokenTypes []SemanticTokenType

type SemanticTokenType uint64



func EncodeSemanticTokens(tokens []lang.Token) []float64 {
	data := make([]float64, 0)

	for _, token := range tokens {
		var tokenType lsp.
		switch token.TokenType {
		case hclsyntax.TokenIdent:
			// TODO
		case hclsyntax.TokenComment:
			// TODO
		case hclsyntax.TokenQuotedLit:
			// TODO: (include surrounding TokenOQuote & TokenCQuote)
		case hclsyntax.TokenNumberLit:

		}

		encodedToken := [...]float64{
			line, startChar, length, tokenType,
		}
	}

	return data
}
