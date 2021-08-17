package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"

	osmosisapp "github.com/osmosis-labs/osmosis/app"
	osmosisparams "github.com/osmosis-labs/osmosis/app/params"
)

// Codec is the application-wide Amino codec and is initialized upon package loading.
var (
	AppCodec       codec.Marshaler
	AminoCodec     *codec.LegacyAmino
	EncodingConfig osmosisparams.EncodingConfig
)

// SetCodec sets encoding config.
func SetCodec() {
	EncodingConfig = osmosisapp.MakeEncodingConfig()
	AppCodec = EncodingConfig.Marshaler
	AminoCodec = EncodingConfig.Amino
}
