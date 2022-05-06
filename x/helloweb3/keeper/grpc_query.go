package keeper

import (
	"github.com/tejiriaustin/hello-web3/x/helloweb3/types"
)

var _ types.QueryServer = Keeper{}
