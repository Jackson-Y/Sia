package sia

import (
	"math/big"
	"testing"

	"github.com/NebulousLabs/Sia/consensus"
	"github.com/NebulousLabs/Sia/network"
)

// establishTestingEnvrionment sets all of the testEnv variables.
func establishTestingEnvironment(t *testing.T) (c *Core) {
	// Alter the constants to create a system more friendly to testing.
	//
	// TODO: Perhaps also have these constants as a build flag, then they don't
	// need to be variables.
	consensus.BlockFrequency = consensus.Timestamp(1)
	consensus.TargetWindow = consensus.BlockHeight(1000)
	network.BootstrapPeers = []network.Address{"localhost:9988", "localhost:9989"}
	consensus.RootTarget[0] = 255
	consensus.MaxAdjustmentUp = big.NewRat(1005, 1000)
	consensus.MaxAdjustmentDown = big.NewRat(995, 1000)

	coreConfig := Config{
		HostDir:     "hostdir",
		WalletFile:  "test.wallet",
		ServerAddr:  ":9988",
		Nobootstrap: true,
	}

	c, err := CreateCore(coreConfig)
	if err != nil {
		t.Fatal(err)
	}

	return
}

// I'm not sure how to test asynchronous code, so at this point I don't try, I
// only test the synchronous parts.
func TestEverything(t *testing.T) {
	c := establishTestingEnvironment(t)
	testEmptyBlock(t, c)
	testTransactionBlock(t, c)
	testSendToSelf(t, c)
	testWalletInfo(t, c)
}
