// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const DataAvailabilityChallengeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"challengeWindow\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"resolveWindow\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"bondSize\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"balances\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1009,\"contract\":\"src/L1/DataAvailabilityChallenge.sol:DataAvailabilityChallenge\",\"label\":\"challenges\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_mapping(t_uint256,t_mapping(t_bytes32,t_struct(Challenge)1010_storage))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_bytes32,t_struct(Challenge)1010_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e struct Challenge)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_struct(Challenge)1010_storage\"},\"t_mapping(t_uint256,t_mapping(t_bytes32,t_struct(Challenge)1010_storage))\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e mapping(bytes32 =\u003e struct Challenge))\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_mapping(t_bytes32,t_struct(Challenge)1010_storage)\"},\"t_struct(Challenge)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Challenge\",\"numberOfBytes\":\"128\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var DataAvailabilityChallengeStorageLayout = new(solc.StorageLayout)

var DataAvailabilityChallengeDeployedBin = "0x6080604052600436106101125760003560e01c8063715018a6116100a5578063c459f80911610074578063d0e30db011610059578063d0e30db014610398578063d7d04e54146103a0578063f2fde38b146103c057600080fd5b8063c459f809146102dd578063c4ee20d4146102fd57600080fd5b8063715018a614610250578063861a1412146102655780638da5cb5b1461027b578063b740a2db146102b057600080fd5b80633ccfd60b116100e15780633ccfd60b146101af5780634ec81af1146101c457806354fd4d50146101e45780637099c5811461023a57600080fd5b806302b2f7c7146101265780630b1a73f41461013957806321cf39ee1461015957806327e235e31461018257600080fd5b366101215761011f6103e0565b005b600080fd5b61011f610134366004610ed6565b61044e565b34801561014557600080fd5b5061011f610154366004610ef8565b610656565b34801561016557600080fd5b5061016f60665481565b6040519081526020015b60405180910390f35b34801561018e57600080fd5b5061016f61019d366004610fa1565b60686020526000908152604090205481565b3480156101bb57600080fd5b5061011f61077c565b3480156101d057600080fd5b5061011f6101df366004610fc3565b6107da565b3480156101f057600080fd5b5061022d6040518060400160405280600581526020017f302e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101799190610ffc565b34801561024657600080fd5b5061016f60675481565b34801561025c57600080fd5b5061011f61098c565b34801561027157600080fd5b5061016f60655481565b34801561028757600080fd5b5060335460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610179565b3480156102bc57600080fd5b506102d06102cb366004610ed6565b6109a0565b604051610179919061109e565b3480156102e957600080fd5b5061011f6102f8366004610ed6565b610a48565b34801561030957600080fd5b50610361610318366004610ed6565b6069602090815260009283526040808420909152908252902080546001820154600283015460039093015473ffffffffffffffffffffffffffffffffffffffff90921692909184565b6040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152608001610179565b61011f6103e0565b3480156103ac57600080fd5b5061011f6103bb3660046110df565b610b5d565b3480156103cc57600080fd5b5061011f6103db366004610fa1565b610ba0565b33600090815260686020526040812080543492906103ff908490611127565b909155505033600081815260686020908152604091829020548251938452908301527fa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05910160405180910390a1565b6104566103e0565b6067543360009081526068602052604090205410156104c75733600090815260686020526040908190205460675491517e0155b50000000000000000000000000000000000000000000000000000000081526104be9290600401918252602082015260400190565b60405180910390fd5b60006104d383836109a0565b60038111156104e4576104e461106f565b1461051b576040517f9bb6c64e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61052482610c57565b61055a576040517ff9e0d1f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606754336000908152606860205260408120805490919061057c90849061113f565b90915550506040805160808101825233815260675460208083019182524383850190815260006060850181815288825260698452868220888352909352859020935184547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178455915160018085019190915591516002840155516003909201919091559051839183917f73b78891d84bab8633915b22168a5ed8a2f0b86fbaf9733698fbacea9a2b11f89161064a9161109e565b60405180910390a35050565b8181604051610666929190611156565b604051809103902083146106c4578181604051610684929190611156565b6040519081900381207f3b7d73720000000000000000000000000000000000000000000000000000000082526004820152602481018490526044016104be565b60016106d085856109a0565b60038111156106e1576106e161106f565b14610718576040517fbeb11d3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000848152606960209081526040808320868452909152908190204360038201559051859085907f73b78891d84bab8633915b22168a5ed8a2f0b86fbaf9733698fbacea9a2b11f89061076d9060029061109e565b60405180910390a35050505050565b3360008181526068602052604081208054908290559161079d905a84610c79565b9050806107d6576040517f27fcd9d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b600054610100900460ff16158080156107fa5750600054600160ff909116105b806108145750303b158015610814575060005460ff166001145b6108a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104be565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108fe57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610906610c8f565b6065849055606683905561091982610b5d565b61092285610d2e565b801561098557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b610994610da5565b61099e6000610d2e565b565b600082815260696020908152604080832084845282528083208151608081018352815473ffffffffffffffffffffffffffffffffffffffff1680825260018301549482019490945260028201549281019290925260030154606082015290610a0c576000915050610a42565b606081015115610a20576002915050610a42565b610a2d8160400151610e26565b15610a3c576001915050610a42565b60039150505b92915050565b6003610a5483836109a0565b6003811115610a6557610a6561106f565b14610a9c576040517f151f07fe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260696020908152604080832084845282528083206001810154815473ffffffffffffffffffffffffffffffffffffffff1685526068909352908320805491939091610aed908490611127565b9091555050600060018201819055815473ffffffffffffffffffffffffffffffffffffffff1680825260686020908152604092839020548351928352908201527fa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05910160405180910390a1505050565b610b65610da5565b60678190556040518181527f4468d695a0389e5f9e8ef0c9aee6d84e74cc0d0e0a28c8413badb54697d1bbae9060200160405180910390a150565b610ba8610da5565b73ffffffffffffffffffffffffffffffffffffffff8116610c4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104be565b610c5481610d2e565b50565b60008143118015610a425750606554610c709083611127565b43111592915050565b600080600080600080868989f195945050505050565b600054610100900460ff16610d26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104be565b61099e610e36565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461099e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104be565b600060665482610c709190611127565b600054610100900460ff16610ecd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104be565b61099e33610d2e565b60008060408385031215610ee957600080fd5b50508035926020909101359150565b60008060008060608587031215610f0e57600080fd5b8435935060208501359250604085013567ffffffffffffffff80821115610f3457600080fd5b818701915087601f830112610f4857600080fd5b813581811115610f5757600080fd5b886020828501011115610f6957600080fd5b95989497505060200194505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610f9c57600080fd5b919050565b600060208284031215610fb357600080fd5b610fbc82610f78565b9392505050565b60008060008060808587031215610fd957600080fd5b610fe285610f78565b966020860135965060408601359560600135945092505050565b600060208083528351808285015260005b818110156110295785810183015185820160400152820161100d565b8181111561103b576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60208101600483106110d9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6000602082840312156110f157600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561113a5761113a6110f8565b500190565b600082821015611151576111516110f8565b500390565b818382376000910190815291905056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(DataAvailabilityChallengeStorageLayoutJSON), DataAvailabilityChallengeStorageLayout); err != nil {
		panic(err)
	}

	layouts["DataAvailabilityChallenge"] = DataAvailabilityChallengeStorageLayout
	deployedBytecodes["DataAvailabilityChallenge"] = DataAvailabilityChallengeDeployedBin
	immutableReferences["DataAvailabilityChallenge"] = false
}
