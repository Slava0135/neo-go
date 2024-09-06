package neotest

import (
	"testing"

	"github.com/nspcc-dev/neo-go/pkg/compiler"
	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/nspcc-dev/neo-go/pkg/vm/opcode"
	"github.com/stretchr/testify/require"
)

func TestProcessCover_OneMethodOneDocument(t *testing.T) {
	scriptHash := util.Uint160{1}
	doc := "foobar.go"
	mdi := compiler.MethodDebugInfo{
		SeqPoints: []compiler.DebugSeqPoint{
			{Opcode: 0, Document: 0, StartLine: 0, EndLine: 0},
			{Opcode: 1, Document: 0, StartLine: 1, EndLine: 1},
			{Opcode: 2, Document: 0, StartLine: 2, EndLine: 2},
		},
	}
	di := &compiler.DebugInfo{
		Documents: []string{doc},
		Methods:   []compiler.MethodDebugInfo{mdi},
	}
	contract := &Contract{Hash: scriptHash, DebugInfo: di}

	addScriptToCoverage(contract)
	coverageHook(scriptHash, 1, opcode.NOP)
	coverageHook(scriptHash, 2, opcode.NOP)
	coverageHook(scriptHash, 2, opcode.NOP)
	cover := processCover()

	require.Contains(t, cover, doc)
	documentCover := cover[doc]
	require.Equal(t, 3, len(documentCover))
	require.Contains(t, documentCover, coverBlock{startLine: 0, endLine: 0, stmts: 1, counts: 0})
	require.Contains(t, documentCover, coverBlock{startLine: 1, endLine: 1, stmts: 1, counts: 1})
	require.Contains(t, documentCover, coverBlock{startLine: 2, endLine: 2, stmts: 1, counts: 2})
}
