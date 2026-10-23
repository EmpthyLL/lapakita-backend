package wallet

import "github.com/google/wire"

var WalletFeatureSet = wire.NewSet(PayoutMethodFeatureSet)
