package hivego


import (
	"fmt"
	
	//~ "strings"
	"encoding/json"
)

var (
	
)



func (h *HiveRpcNode) CheckAccount(account string) (string, error) {
	query := hrpcQuery {
		method: "condenser_api.lookup_accounts",
		params: []any{account, 1},
	}

	response, err := h.rpcExec(h.address, query)
	if err != nil {
		return "", fmt.Errorf("response err: %s", err)
	}

	var accts []interface{}
	if uErr := json.Unmarshal(response, &accts); uErr != nil {
		return "", fmt.Errorf("Unmarshal err: %s", uErr)
	}
	if accts[0].(string) == account {
		return accts[0].(string), nil
	} else {
		return "", fmt.Errorf("Account Not Found.")
	}
}

func (h *HiveRpcNode) CheckKey(wif *string, account string) (bool, error) {
	//get key pair from private key
	keyPair, err := KeyPairFromWif(*wif)
	if err != nil {
		return false, err
	}
	//decode public key to string
	pubKey := GetPublicKeyString(keyPair.PublicKey)
	//get account name from public key
	query := hrpcQuery {
		method: "condenser_api.get_key_references",
		params: []any{[]string{*pubKey}},
	}
	response, err := h.rpcExec(h.address, query)
	if err != nil {
		return false, fmt.Errorf("response err: %s", err)
	}
	var accts []interface{}
	if uErr := json.Unmarshal(response, &accts); uErr != nil {
		return false, fmt.Errorf("Unmarshal err: %s", uErr)
	}
	//untangle response and see if it matches expected account
	holder := accts[0].([]interface{})
	if holder[0].(string) == account {
		return true, nil
	} else {
		return false, fmt.Errorf("Key does not match account.")
	}
}
