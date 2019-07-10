package rest

import (
	"fmt"
	"net/http"

	"github.com/EG-easy/votumchain/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

const (
	restName = "name"
)

//REST APIをroutingする
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/coins/bananceOf", storeName), balanceOfHandler(cdc, cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/coins/transfer", storeName), transferHandler(cdc, cliCtx)).Methods("POST")
}

type TransferReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	FromAddr string       `json:"from_addr"`
	ToAddr   string       `json:"to_addr"`
	Amt      string       `json:"amt"`
}

func transferHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req TransferReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.FromAddr)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		toAddr, err := sdk.AccAddressFromBech32(req.ToAddr)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		coins, err := sdk.ParseCoins(req.Amt)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgTransferCoin(fromAddr, toAddr, coins)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)
	}
}

func balanceOfHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/balanceOf/%s", storeName, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}
