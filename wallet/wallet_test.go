package wallet_test

import (
	"os"
	"testing"

	"github.com/test-go/testify/require"

	bip39 "github.com/cosmos/go-bip39"

	"github.com/kogisin/arbitrage-bot/codec"
	"github.com/kogisin/arbitrage-bot/wallet"
)

func TestMain(m *testing.M) {
	codec.SetCodec()
	os.Exit(m.Run())
}

func TestRecoverAccAddrFromMnemonic(t *testing.T) {
	testCases := []struct {
		mnemonic   string
		password   string
		expAccAddr string
	}{
		{
			mnemonic:   "guard cream sadness conduct invite crumble clock pudding hole grit liar hotel maid produce squeeze return argue turtle know drive eight casino maze host",
			password:   "",
			expAccAddr: "osmo1zaavvzxez0elundtn32qnk9lkm8kmcsz2tlhe7",
		},
		{
			mnemonic:   "friend excite rough reopen cover wheel spoon convince island path clean monkey play snow number walnut pull lock shoot hurry dream divide concert discover",
			password:   "",
			expAccAddr: "osmo1mzgucqnfr2l8cj5apvdpllhzt4zeuh2ccv3ysw",
		},
		{
			mnemonic:   "friend excite rough reopen cover wheel spoon convince island path clean monkey play snow number walnut pull lock shoot hurry dream divide concert discover",
			password:   "",
			expAccAddr: "osmo1mzgucqnfr2l8cj5apvdpllhzt4zeuh2ccv3ysw",
		},
	}

	for _, tc := range testCases {
		accAddr, _, err := wallet.RecoverAccountFromMnemonic(tc.mnemonic, tc.password)
		require.NoError(t, err)

		require.Equal(t, tc.expAccAddr, accAddr)
	}
}

func TestNewMnemonic(t *testing.T) {
	for i := 0; i < 5; i++ {
		entropy, err := bip39.NewEntropy(256)
		require.NoError(t, err)

		mnemonic, err := bip39.NewMnemonic(entropy)
		require.NoError(t, err)

		accAddr, _, err := wallet.RecoverAccountFromMnemonic(mnemonic, "")
		require.NoError(t, err)

		t.Log(mnemonic, accAddr)
	}
}
