package sfc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is SFC contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from opera-sfc 424031c81a77196f4e9d60c7d876032dd47208ce, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 200
func GetContractBin() []byte {
	return hexutil.MustDecode("0x6080604052600436106102855760003560e01c80638b0e9f3f11610153578063c65ee0e1116100cb578063d46fa5181161007f578063e08d7e6611610064578063e08d7e6614610b1f578063e6f45adf14610b9c578063f2fde38b14610bcf57610285565b8063d46fa51814610af5578063d96ed50514610b0a57610285565b8063cc8343aa116100b0578063cc8343aa14610a51578063cfd4766314610a83578063cfdbb7cd14610abc57610285565b8063c65ee0e114610a12578063c7be95de14610a3c57610285565b8063a2f6e6bc11610122578063b5d8962711610107578063b5d8962714610959578063b810e411146109c4578063c5f956af146109fd57610285565b8063a2f6e6bc146108ed578063a86a056f1461092057610285565b80638b0e9f3f1461083b5780638da5cb5b146108505780638f32d59b1461086557806396c7ee461461088e57610285565b8063592fe0c0116102015780637cacb1d6116101b5578063854873e11161019a578063854873e114610754578063860c2750146107f3578063893675c61461082657610285565b80637cacb1d61461070c578063841e45611461072157610285565b8063670322f8116101e6578063670322f8146106a9578063715018a6146106e257806376671808146106f757610285565b8063592fe0c0146105215780635fab23a81461069457610285565b80631f2701521161025857806339b80c001161023d57806339b80c001461044257806354fd4d50146104a4578063550359a0146104ee57610285565b80631f270152146103d057806328f731481461042d57610285565b80630135b1db146102ee5780630e559d821461033357806310e51e141461036457806318160ddd146103bb575b366102d7576040805162461bcd60e51b815260206004820152601560248201527f7472616e7366657273206e6f7420616c6c6f7765640000000000000000000000604482015290519081900360640190fd5b6081546102ec906001600160a01b0316610c02565b005b3480156102fa57600080fd5b506103216004803603602081101561031157600080fd5b50356001600160a01b0316610c2b565b60408051918252519081900360200190f35b34801561033f57600080fd5b50610348610c3d565b604080516001600160a01b039092168252519081900360200190f35b34801561037057600080fd5b506102ec600480360360c081101561038757600080fd5b508035906020810135906001600160a01b0360408201358116916060810135821691608082013581169160a0013516610c4c565b3480156103c757600080fd5b50610321610dd9565b3480156103dc57600080fd5b5061040f600480360360608110156103f357600080fd5b506001600160a01b038135169060208101359060400135610ddf565b60408051938452602084019290925282820152519081900360600190f35b34801561043957600080fd5b50610321610e11565b34801561044e57600080fd5b5061046c6004803603602081101561046557600080fd5b5035610e17565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b3480156104b057600080fd5b506104b9610e59565b604080517fffffff00000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b3480156104fa57600080fd5b506102ec6004803603602081101561051157600080fd5b50356001600160a01b0316610e7e565b34801561052d57600080fd5b506102ec600480360360a081101561054457600080fd5b81019060208101813564010000000081111561055f57600080fd5b82018360208201111561057157600080fd5b8035906020019184602083028401116401000000008311171561059357600080fd5b9193909290916020810190356401000000008111156105b157600080fd5b8201836020820111156105c357600080fd5b803590602001918460208302840111640100000000831117156105e557600080fd5b91939092909160208101903564010000000081111561060357600080fd5b82018360208201111561061557600080fd5b8035906020019184602083028401116401000000008311171561063757600080fd5b91939092909160208101903564010000000081111561065557600080fd5b82018360208201111561066757600080fd5b8035906020019184602083028401116401000000008311171561068957600080fd5b919350915035610f11565b3480156106a057600080fd5b506103216111cc565b3480156106b557600080fd5b50610321600480360360408110156106cc57600080fd5b506001600160a01b0381351690602001356111d2565b3480156106ee57600080fd5b506102ec611216565b34801561070357600080fd5b506103216112d1565b34801561071857600080fd5b506103216112da565b34801561072d57600080fd5b506102ec6004803603602081101561074457600080fd5b50356001600160a01b03166112e0565b34801561076057600080fd5b5061077e6004803603602081101561077757600080fd5b5035611373565b6040805160208082528351818301528351919283929083019185019080838360005b838110156107b85781810151838201526020016107a0565b50505050905090810190601f1680156107e55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156107ff57600080fd5b506102ec6004803603602081101561081657600080fd5b50356001600160a01b031661142c565b34801561083257600080fd5b506103486114bf565b34801561084757600080fd5b506103216114ce565b34801561085c57600080fd5b506103486114d4565b34801561087157600080fd5b5061087a6114e3565b604080519115158252519081900360200190f35b34801561089a57600080fd5b506108c7600480360360408110156108b157600080fd5b506001600160a01b0381351690602001356114f4565b604080519485526020850193909352838301919091526060830152519081900360800190f35b3480156108f957600080fd5b506102ec6004803603602081101561091057600080fd5b50356001600160a01b0316611526565b34801561092c57600080fd5b506103216004803603604081101561094357600080fd5b506001600160a01b0381351690602001356115b9565b34801561096557600080fd5b506109836004803603602081101561097c57600080fd5b50356115d6565b604080519788526020880196909652868601949094526060860192909252608085015260a08401526001600160a01b031660c0830152519081900360e00190f35b3480156109d057600080fd5b5061040f600480360360408110156109e757600080fd5b506001600160a01b03813516906020013561161c565b348015610a0957600080fd5b50610348611648565b348015610a1e57600080fd5b5061032160048036036020811015610a3557600080fd5b5035611657565b348015610a4857600080fd5b50610321611669565b348015610a5d57600080fd5b506102ec60048036036040811015610a7457600080fd5b5080359060200135151561166f565b348015610a8f57600080fd5b5061032160048036036040811015610aa657600080fd5b506001600160a01b03813516906020013561189e565b348015610ac857600080fd5b5061087a60048036036040811015610adf57600080fd5b506001600160a01b0381351690602001356118bb565b348015610b0157600080fd5b50610348611952565b348015610b1657600080fd5b50610321611961565b348015610b2b57600080fd5b506102ec60048036036020811015610b4257600080fd5b810190602081018135640100000000811115610b5d57600080fd5b820183602082011115610b6f57600080fd5b80359060200191846020830284011164010000000083111715610b9157600080fd5b509092509050611967565b348015610ba857600080fd5b506102ec60048036036020811015610bbf57600080fd5b50356001600160a01b0316611aab565b348015610bdb57600080fd5b506102ec60048036036020811015610bf257600080fd5b50356001600160a01b0316611b3e565b3660008037600080366000845af43d6000803e808015610c21573d6000f35b3d6000fd5b505050565b60696020526000908152604090205481565b607c546001600160a01b031681565b600054610100900460ff1680610c655750610c65611ba3565b80610c73575060005460ff16155b610cae5760405162461bcd60e51b815260040180806020018281038252602e81526020018061335a602e913960400191505060405180910390fd5b600054610100900460ff16158015610d1457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b610d1d82611ba9565b6067879055606680546001600160a01b038088167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556081805487841690831617905560828054928616929091169190911790556077869055610d84611d0b565b607f55610d8f611d14565b6000888152607860205260409020600701558015610dd057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50505050505050565b60775481565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b606d5481565b607860205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b7f33303400000000000000000000000000000000000000000000000000000000005b90565b610e866114e3565b610ed7576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b608380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610f1a33611d18565b610f555760405162461bcd60e51b81526004018080602001828103825260298152602001806133106029913960400191505060405180910390fd5b600060786000610f636112d1565b81526020019081526020016000209050606081600601805480602002602001604051908101604052809291908181526020018280548015610fc357602002820191906000526020600020905b815481526020019060010190808311610faf575b5050505050905061104a82828d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611d2f92505050565b6067546000908152607860205260409020600781015460019061106b611d14565b111561108257816007015461107e611d14565b0390505b611104818584868d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611f3692505050565b61110e8186612720565b50506111186112d1565b606755611123611d14565b6007830155608254604080517fd9a7c1f900000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169163d9a7c1f991600480820192602092909190829003018186803b15801561118657600080fd5b505afa15801561119a573d6000803e3d6000fd5b505050506040513d60208110156111b057600080fd5b5051600b83015550607754600d90910155505050505050505050565b606e5481565b60006111de83836118bb565b6111ea57506000611210565b506001600160a01b03821660009081526074602090815260408083208484529091529020545b92915050565b61121e6114e3565b61126f576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60675460010190565b60675481565b6112e86114e3565b611339576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b608080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b606a6020908152600091825260409182902080548351601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600186161502019093169290920491820184900484028101840190945280845290918301828280156114245780601f106113f957610100808354040283529160200191611424565b820191906000526020600020905b81548152906001019060200180831161140757829003601f168201915b505050505081565b6114346114e3565b611485576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b608280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6083546001600160a01b031681565b606c5481565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b607460209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b61152e6114e3565b61157f576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b607c80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b607060209081526000928352604080842090915290825290205481565b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b607560209081526000928352604080842090915290825290208054600182015460029092015490919083565b6080546001600160a01b031681565b607b6020526000908152604090205481565b606b5481565b61167882612899565b6116c9576040805162461bcd60e51b815260206004820152601760248201527f76616c696461746f7220646f65736e2774206578697374000000000000000000604482015290519081900360640190fd5b600082815260686020526040902060038101549054156116e7575060005b606654604080517fa4066fbe000000000000000000000000000000000000000000000000000000008152600481018690526024810184905290516001600160a01b039092169163a4066fbe9160448082019260009290919082900301818387803b15801561175457600080fd5b505af1158015611768573d6000803e3d6000fd5b5050505081801561177857508015155b15610c26576066546000848152606a60205260409081902081517f242a6e3f0000000000000000000000000000000000000000000000000000000081526004810187815260248201938452825460027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001831615610100020190911604604483018190526001600160a01b039095169463242a6e3f948994939091606490910190849080156118695780601f1061183e57610100808354040283529160200191611869565b820191906000526020600020905b81548152906001019060200180831161184c57829003601f168201915b50509350505050600060405180830381600087803b15801561188a57600080fd5b505af1158015610dd0573d6000803e3d6000fd5b607360209081526000928352604080842090915290825290205481565b6001600160a01b03821660009081526074602090815260408083208484529091528120600201541580159061191257506001600160a01b038316600090815260746020908152604080832085845290915290205415155b801561194b57506001600160a01b0383166000908152607460209081526040808320858452909152902060020154611948611d14565b11155b9392505050565b6082546001600160a01b031690565b607f5481565b61197033611d18565b6119ab5760405162461bcd60e51b81526004018080602001828103825260298152602001806133106029913960400191505060405180910390fd5b6000607860006119b96112d1565b8152602001908152602001600020905060008090505b82811015611a325760008484838181106119e557fe5b60209081029290920135600081815260688452604080822060030154948890529020839055600c860154909350611a2391508263ffffffff6128b016565b600c85015550506001016119cf565b50611a41600682018484613238565b50606654607f54604080517f07aaf3440000000000000000000000000000000000000000000000000000000081526004810192909252516001600160a01b03909216916307aaf3449160248082019260009290919082900301818387803b15801561188a57600080fd5b611ab36114e3565b611b04576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b608180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b611b466114e3565b611b97576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b611ba08161290a565b50565b303b1590565b600054610100900460ff1680611bc25750611bc2611ba3565b80611bd0575060005460ff16155b611c0b5760405162461bcd60e51b815260040180806020018281038252602e81526020018061335a602e913960400191505060405180910390fd5b600054610100900460ff16158015611c7157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015611d0757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050565b64174876e80090565b4290565b6066546001600160a01b038281169116145b919050565b60005b8351811015611f2f57608260009054906101000a90046001600160a01b03166001600160a01b0316635a68f01a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611d8957600080fd5b505afa158015611d9d573d6000803e3d6000fd5b505050506040513d6020811015611db357600080fd5b50518251839083908110611dc357fe5b6020026020010151118015611e655750608260009054906101000a90046001600160a01b03166001600160a01b031662cc7f836040518163ffffffff1660e01b815260040160206040518083038186803b158015611e2057600080fd5b505afa158015611e34573d6000803e3d6000fd5b505050506040513d6020811015611e4a57600080fd5b50518351849083908110611e5a57fe5b602002602001015110155b15611ea657611e88848281518110611e7957fe5b602002602001015160086129c3565b611ea6848281518110611e9757fe5b6020026020010151600061166f565b828181518110611eb257fe5b6020026020010151856004016000868481518110611ecc57fe5b6020026020010151815260200190815260200160002081905550818181518110611ef257fe5b6020026020010151856005016000868481518110611f0c57fe5b602090810291909101810151825281019190915260400160002055600101611d32565b5050505050565b611f3e61327f565b6040518060a001604052808551604051908082528060200260200182016040528015611f74578160200160208202803883390190505b508152602001600081526020018551604051908082528060200260200182016040528015611fac578160200160208202803883390190505b508152602001600081526020016000815250905060008090505b84518110156120c7576000866003016000878481518110611fe357fe5b6020026020010151815260200190815260200160002054905060008090508185848151811061200e57fe5b60200260200101511115612035578185848151811061202957fe5b60200260200101510390505b8986848151811061204257fe5b602002602001015182028161205357fe5b048460400151848151811061206457fe5b60200260200101818152505061209e8460400151848151811061208357fe5b602002602001015185606001516128b090919063ffffffff16565b606085015260808401516120b8908263ffffffff6128b016565b60808501525050600101611fc6565b5060005b845181101561219057878482815181106120e157fe5b6020026020010151898684815181106120f657fe5b60200260200101518a60000160008a878151811061211057fe5b6020026020010151815260200190815260200160002054028161212f57fe5b04028161213857fe5b048260000151828151811061214957fe5b6020026020010181815250506121838260000151828151811061216857fe5b602002602001015183602001516128b090919063ffffffff16565b60208301526001016120cb565b5060005b84518110156125cf57600061223d89608260009054906101000a90046001600160a01b03166001600160a01b031663d9a7c1f96040518163ffffffff1660e01b815260040160206040518083038186803b1580156121f157600080fd5b505afa158015612205573d6000803e3d6000fd5b505050506040513d602081101561221b57600080fd5b5051855180518690811061222b57fe5b60200260200101518660200151612aed565b905061227961226c84608001518560400151858151811061225a57fe5b60200260200101518660600151612b3c565b829063ffffffff6128b016565b9050600086838151811061228957fe5b60209081029190910181015160008181526068835260408082206006015460825482517fa778651500000000000000000000000000000000000000000000000000000000815292519496506001600160a01b0391821695939461233f948994929093169263a77865159260048082019391829003018186803b15801561230e57600080fd5b505afa158015612322573d6000803e3d6000fd5b505050506040513d602081101561233857600080fd5b5051612ca5565b6001600160a01b038316600090815260736020908152604080832087845290915290205490915080156124e65760008161237985876111d2565b84028161238257fe5b0490508083036123906132ae565b6001600160a01b03861660009081526074602090815260408083208a84529091529020600301546123c2908490612cc2565b90506123cc6132ae565b6123d7836000612cc2565b6001600160a01b0388166000908152606f602090815260408083208c8452825291829020825160608101845281548152600182015492810192909252600201549181019190915290915061242c908383612e84565b6001600160a01b0388166000818152606f602090815260408083208d84528252808320855181558583015160018083019190915595820151600291820155938352607582528083208d8452825291829020825160608101845281548152948101549185019190915290910154908201526124a7908383612e84565b6001600160a01b03881660009081526075602090815260408083208c845282529182902083518155908301516001820155910151600290910155505050505b600084815260686020526040812060030154838703918115612518578161250b612e9f565b84028161251457fe5b0490505b808e600101600089815260200190815260200160002054018f6001016000898152602001908152602001600020819055508a898151811061255557fe5b60200260200101518f6003016000898152602001908152602001600020819055508b898151811061258257fe5b60200260200101518e600201600089815260200190815260200160002054018f60020160008981526020019081526020016000208190555050505050505050508080600101915050612194565b50608081015160088701819055602082015160098801556060820151600a880155607754111561260d57600886015460778054919091039055612613565b60006077555b6080546001600160a01b031615610dd057600061262e612e9f565b608260009054906101000a90046001600160a01b03166001600160a01b03166394c3e9146040518163ffffffff1660e01b815260040160206040518083038186803b15801561267c57600080fd5b505afa158015612690573d6000803e3d6000fd5b505050506040513d60208110156126a657600080fd5b5051608084015102816126b557fe5b0490506126c181612eab565b6080546040516001600160a01b03909116908290600081818185875af1925050503d806000811461270e576040519150601f19603f3d011682016040523d82523d6000602084013e612713565b606091505b5050505050505050505050565b608254604080517f3a3ef66c00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b031691633a3ef66c916004808301926020929190829003018186803b15801561277e57600080fd5b505afa158015612792573d6000803e3d6000fd5b505050506040513d60208110156127a857600080fd5b5051830260010190506000816127bc612e9f565b8402816127c557fe5b0490506000608260009054906101000a90046001600160a01b03166001600160a01b0316632c8c36a56040518163ffffffff1660e01b815260040160206040518083038186803b15801561281857600080fd5b505afa15801561282c573d6000803e3d6000fd5b505050506040513d602081101561284257600080fd5b50519050848101612851612e9f565b8202838702018161285e57fe5b04915061286a82612f49565b91506000612876612e9f565b83607f54028161288257fe5b04905061288e81612fb7565b607f55505050505050565b600090815260686020526040902060050154151590565b60008282018381101561194b576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b03811661294f5760405162461bcd60e51b81526004018080602001828103825260268152602001806132ea6026913960400191505060405180910390fd5b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6000828152606860205260409020541580156129de57508015155b15612a0b57600082815260686020526040902060030154606d54612a079163ffffffff612fed16565b606d555b600082815260686020526040902054811115611d0757600082815260686020526040902081815560020154612ab357612a426112d1565b600083815260686020526040902060020155612a5c611d14565b6000838152606860209081526040918290206001810184905560020154825190815290810192909252805184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4792908290030190a25b60408051828152905183917fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e919081900360200190a25050565b600082612afc57506000612b34565b6000612b0e868663ffffffff61302f16565b9050612b3083612b24838763ffffffff61302f16565b9063ffffffff61308816565b9150505b949350505050565b600082612b4b5750600061194b565b6000612b6183612b24878763ffffffff61302f16565b9050612c9c612b6e612e9f565b608254604080517f94c3e9140000000000000000000000000000000000000000000000000000000081529051612b24926001600160a01b0316916394c3e914916004808301926020929190829003018186803b158015612bcd57600080fd5b505afa158015612be1573d6000803e3d6000fd5b505050506040513d6020811015612bf757600080fd5b5051608254604080517fc74dd62100000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169163c74dd62191600480820192602092909190829003018186803b158015612c5757600080fd5b505afa158015612c6b573d6000803e3d6000fd5b505050506040513d6020811015612c8157600080fd5b5051612c8b612e9f565b03038461302f90919063ffffffff16565b95945050505050565b600061194b612cb2612e9f565b612b24858563ffffffff61302f16565b612cca6132ae565b60405180606001604052806000815260200160008152602001600081525090506000608260009054906101000a90046001600160a01b03166001600160a01b0316635e2308d26040518163ffffffff1660e01b815260040160206040518083038186803b158015612d3a57600080fd5b505afa158015612d4e573d6000803e3d6000fd5b505050506040513d6020811015612d6457600080fd5b505190508215612e5c57600081612d79612e9f565b0390506000612e0b608260009054906101000a90046001600160a01b03166001600160a01b0316630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b158015612dcf57600080fd5b505afa158015612de3573d6000803e3d6000fd5b505050506040513d6020811015612df957600080fd5b5051612b24848863ffffffff61302f16565b90506000612e2c612e1a612e9f565b612b248987860163ffffffff61302f16565b9050612e49612e39612e9f565b612b24898763ffffffff61302f16565b602086018190529003845250612e7d9050565b612e77612e67612e9f565b612b24868463ffffffff61302f16565b60408301525b5092915050565b612e8c6132ae565b612b34612e9985856130ca565b836130ca565b670de0b6b3a764000090565b606654604080517f66e7ea0f0000000000000000000000000000000000000000000000000000000081523060048201526024810184905290516001600160a01b03909216916366e7ea0f9160448082019260009290919082900301818387803b158015612f1757600080fd5b505af1158015612f2b573d6000803e3d6000fd5b5050607754612f43925090508263ffffffff6128b016565b60775550565b60006064612f55612e9f565b60690281612f5f57fe5b04821115612f83576064612f71612e9f565b60690281612f7b57fe5b049050611d2a565b6064612f8d612e9f565b605f0281612f9757fe5b04821015612fb3576064612fa9612e9f565b605f0281612f7b57fe5b5090565b600066038d7ea4c68000821115612fd6575066038d7ea4c68000611d2a565b633b9aca00821015612fb35750633b9aca00611d2a565b600061194b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061313c565b60008261303e57506000611210565b8282028284828161304b57fe5b041461194b5760405162461bcd60e51b81526004018080602001828103825260218152602001806133396021913960400191505060405180910390fd5b600061194b83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506131d3565b6130d26132ae565b60408051606081019091528251845182916130f3919063ffffffff6128b016565b8152602001613113846020015186602001516128b090919063ffffffff16565b8152602001613133846040015186604001516128b090919063ffffffff16565b90529392505050565b600081848411156131cb5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613190578181015183820152602001613178565b50505050905090810190601f1680156131bd5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600081836132225760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315613190578181015183820152602001613178565b50600083858161322e57fe5b0495945050505050565b828054828255906000526020600020908101928215613273579160200282015b82811115613273578235825591602001919060010190613258565b50612fb39291506132cf565b6040518060a0016040528060608152602001600081526020016060815260200160008152602001600081525090565b60405180606001604052806000815260200160008152602001600081525090565b610e7b91905b80821115612fb357600081556001016132d556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737363616c6c6572206973206e6f7420746865204e6f64654472697665724175746820636f6e7472616374536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a265627a7a723158208380cbacd33be0cf15aac5c729f0af38bad66f8146f0e0dcb033ae3bef57098c64736f6c63430005110032")
}

// ContractAddress is the SFC contract address
var ContractAddress = common.HexToAddress("0xfc00face00000000000000000000000000000000")
