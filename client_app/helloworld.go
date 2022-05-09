package main

import (
	"os"
	"time"
	"sync/atomic"

	fpc "github.com/hyperledger/fabric-private-chaincode/client_sdk/go/pkg/gateway"
	"github.com/hyperledger/fabric-private-chaincode/integration/client_sdk/go/utils"
	"github.com/hyperledger/fabric/common/flogging"
)

var logger = flogging.MustGetLogger("helloworld")
var counter int32

func main() {

	start_time := time.Now()

			ccID := os.Getenv("CC_ID")
			//logger.Infof("Use Chaincode ID: %v", ccID)

			channelID := os.Getenv("CHAN_ID")
			//logger.Infof("Use channel: %v", channelID)

			// get network
			network, err := utils.SetupNetwork(channelID)
			_ = err

			// Get FPC Contract
			contract := fpc.GetContract(network, ccID)

	for {
		if (time.Now().Sub(start_time).Seconds()) <= 30 {
			// Invoke FPC Chaincode storeAsset
			//logger.Infof("--> Invoke FPC Chaincode: storeAsset")
			go func () {
				result, err := contract.SubmitTransaction("storeAsset", "asset1", "100")
				if err != nil {
					//logger.Fatalf("Failed to Submit transaction: %v", err)
					//logger.Infof("Failed to Submit transaction: %v", err)
				} else {
					atomic.AddInt32(&counter, 1)
				}
			//logger.Infof("--> Result: %s", string(result))
				_ = result
			}()

			time.Sleep(14 * time.Millisecond)
			//time.Sleep(100 * time.Millisecond)
			//atomic.AddInt32(&counter, 1)
			// Evaluate FPC Chaincode retrieveAsset
			//logger.Infof("--> Evaluate FPC Chaincode: retrieveAsset")
			//result, err = contract.EvaluateTransaction("retrieveAsset", "asset1")
			//if err != nil {
			//	logger.Fatalf("Failed to Evaluate transaction: %v", err)
			//}
			//logger.Infof("--> Result: %s", string(result))
		} else {
			//logger.Infof("How many times: %d", counter)
			break
		}

	}
	logger.Infof("How many times: %d", counter)
}
