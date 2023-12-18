package paper

import (
	"github.com/karlsen-network/karlsend/cmd/karlsenwallet/libkaspawallet"
	"github.com/karlsen-network/karlsend/domain/dagconfig"
	"github.com/karlsen-network/karlsen-paper/model"
)

// Make sure we implement model.PaperWallet
var _ model.PaperAPI = &api{}

type api struct {
	dagParams *dagconfig.Params
}

func NewAPI(dagParams *dagconfig.Params) model.PaperAPI {
	return &api{dagParams: dagParams}
}

func (a *api) GenerateWallet() (model.PaperWallet, error) {
	mnemonics, err := libkaspawallet.CreateMnemonic()
	if err != nil {
		return nil, err
	}
	// It's safe to use [0] because we know there's exactly 1 key, since we passed numKeys: 1 to CreateMnemonics
	return newWallet(a.dagParams, mnemonics)
}
