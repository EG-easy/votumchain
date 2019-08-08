package votum

import (
	"github.com/EG-easy/votumchain/x/votum/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	// TODO: Msgをここで定める
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	MsgIssueToken = types.MsgIssueToken
)
