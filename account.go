package hivego

import (
	//~ "fmt"
	"strings"
	"encoding/json"
)

var (
	
)

type KeyReferences struct {
	Accounts               []string     `json:"accounts"`
}

func (h *HiveRpcNode) checkAccount(wif *string) (string, error) {
	query := hrpcQuery{
		method: "account_by_key_api.get_key_references",
		params: map[string]string{"keys": [*wif]},
		Limit: limit,
	}
	res, err := h.rpcExec(h.address, query)
	if err != nil {
		return nil, err
	}
	accts := &KeyReferences{}
	if uErr := json.Unmarshal(response, &accts); uErr != nil {
		return nil, uErr
	}
	return accts, nil
}
