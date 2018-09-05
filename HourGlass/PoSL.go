// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PoSL

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PoSLABI is the input ABI used to generate the binding from.
const PoSLABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_customerAddress\",\"type\":\"address\"}],\"name\":\"dividendsOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ethereumToSpend\",\"type\":\"uint256\"}],\"name\":\"calculateTokensReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokensToSell\",\"type\":\"uint256\"}],\"name\":\"calculateEthereumReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"onlyAmbassadors\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"administrators\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sellPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingRequirement\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_includeReferralBonus\",\"type\":\"bool\"}],\"name\":\"myDividends\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalEthereumBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_customerAddress\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountOfTokens\",\"type\":\"uint256\"}],\"name\":\"setStakingRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_identifier\",\"type\":\"bytes32\"},{\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setAdministrator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"myTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableInitialStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_toAddress\",\"type\":\"address\"},{\"name\":\"_amountOfTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"setSymbol\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountOfTokens\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referredBy\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reinvest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"customerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"incomingEthereum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensMinted\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"referredBy\",\"type\":\"address\"}],\"name\":\"onTokenPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"customerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokensBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ethereumEarned\",\"type\":\"uint256\"}],\"name\":\"onTokenSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"customerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ethereumReinvested\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensMinted\",\"type\":\"uint256\"}],\"name\":\"onReinvestment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"customerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ethereumWithdrawn\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PoSLBin is the compiled bytecode used for deploying new contracts.
const PoSLBin = `0x60c0604052600460808190527f506f534c0000000000000000000000000000000000000000000000000000000060a0908152620000409160009190620004e2565b506040805180820190915260038082527f50534c000000000000000000000000000000000000000000000000000000000060209092019182526200008791600191620004e2565b5068056bc75e2d631000006002556000600855600b805460ff19166001179055348015620000b457600080fd5b507f3aa5889d466061262b782fcf0b3e070190660f564317c1fcea2de1999c51e3a78054600160ff19918216811790925560036020527f9a43adbc254c3858aa68afb074fc852f308291dcfb732e14eef35f7952015eee80548216831790557f40486e7ee99d4b525852c9258448d7970264bcb791b680d689a56ecd89e4f3f580548216831790557f1d26cfadf376195ec238a9183182ba897dfd7c2846f0d77c3617eae210acded380548216831790557f40fbe866fba9b68c383bad90baeb9540ff0a7e2a061e408f785db01d8aa0063380548216831790557f05d8ba87f71dcf6ce339dc0e2d1b8545ae44753be9425f1e35ac91503164542d80548216831790557f3842078e3288834a23d73cdba345a898974e4afe60001666a11d029381cc4aa680548216831790557fab5ada39a35b216a4651f3400c2ffb1dead41e8cad7ddfdd12ee139a848d46f880548216831790557f4e0228cc24f8bd8b93c53b6725da40ed8a091ec0a59ff0cd2cc2dd7778c3470980548216831790557f9ed6a5d8cc287453aadafcf658af742379186f8e6f5e03b8ca16a5be81404ca580548216831790557fc77675b631b7343617ce2d0aac9fb0603349427129b6822c64174e883c713ec380548216831790557f28a34fa998cdfbed31d410bb706d7d00cd224708cf6b62cd2ad29e494d307f2680548216831790557f9ca87be9caf86022c7460822fa35d57396e5fc5f88a4d487a1158764af9bbe8b80548216831790557f8d9a98e7f948c089886165e0cb3b0b2e592053374a86a8e697c0b5c756409d1c80548216831790557f1c55d299ca2cde437210518900c54997bb7b2642cb5488e5cd15cf7f3660560080548216831790557f969b3597aeedd6e315a29719eb2ea089eae406f0b7ef5704d0f688e13f2a6b5a80548216831790557fa3c13ec9d3466dbfd902e5e6289198e4bc6de487f207c7e51b21150979a7766f80548216831790557f437d246255ea612b875afc15522f93bd0e951164babc9d396c141e604d5d3f1380548216831790557f78947d42ed56e402e4b6e5be5a40a6eeda83977ef24ff45849d8b8e9a772947980548216831790557f5bb8a42aed3532514e1733604f04913ae88d64643c083ef8533a096afd1cb79680548216831790557f18d054f4c5212cbe4695eadf94d50305224d852ff40d892b918d1670ae3f8a2e80548216831790557f97a1315f54ba10a93e42946b7da8f2bbece48758368f1ce4b02195db5ca7aba180548216831790557fdf40cf4404ab725e1594827a78236ee9328bf6e709e675f5260e8ff410bd0b5380548216831790557f63e9866ed73313fca5a8559571f55f812bba24c2620498d54c3e9a5f484e95e1805482168317905573147fc6b04c95bce47d013c8d7a200ee4343236696000527f68539091a7b4a2d10f3191729c402a2a4d8ea1d3e9c986e602d60fa822414de18054909116909117905562000587565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200052557805160ff191683800117855562000555565b8280016001018555821562000555579182015b828111156200055557825182559160200191906001019062000538565b506200056392915062000567565b5090565b6200058491905b808211156200056357600081556001016200056e565b90565b61166380620005976000396000f30060806040526004361061015d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166265318b811461016b57806306fdde031461019e57806310d0ffdd1461022857806318160ddd14610240578063226093731461025557806327defa1f1461026d578063313ce56714610296578063392efb52146102c15780633ccfd60b146102d95780634b750334146102f057806356d399e814610305578063688abbf71461031a5780636b2f46321461033457806370a08231146103495780638328b6101461036a5780638620410b1461038257806389135ae914610397578063949e8acd146103b457806395d89b41146103c9578063a8e04f34146103de578063a9059cbb146103f3578063b84c824614610417578063c47f002714610470578063e4849b32146104c9578063e9fad8ee146104e1578063f088d547146104f6578063fdb5a03e1461050a575b61016834600061051f565b50005b34801561017757600080fd5b5061018c600160a060020a0360043516610af2565b60408051918252519081900360200190f35b3480156101aa57600080fd5b506101b3610b2d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ed5781810151838201526020016101d5565b50505050905090810190601f16801561021a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561023457600080fd5b5061018c600435610bbb565b34801561024c57600080fd5b5061018c610beb565b34801561026157600080fd5b5061018c600435610bf2565b34801561027957600080fd5b50610282610c2b565b604080519115158252519081900360200190f35b3480156102a257600080fd5b506102ab610c34565b6040805160ff9092168252519081900360200190f35b3480156102cd57600080fd5b50610282600435610c39565b3480156102e557600080fd5b506102ee610c4e565b005b3480156102fc57600080fd5b5061018c610d21565b34801561031157600080fd5b5061018c610d75565b34801561032657600080fd5b5061018c6004351515610d7b565b34801561034057600080fd5b5061018c610dbe565b34801561035557600080fd5b5061018c600160a060020a0360043516610dc3565b34801561037657600080fd5b506102ee600435610dde565b34801561038e57600080fd5b5061018c610e25565b3480156103a357600080fd5b506102ee6004356024351515610e6d565b3480156103c057600080fd5b5061018c610ecf565b3480156103d557600080fd5b506101b3610ee2565b3480156103ea57600080fd5b506102ee610f3c565b3480156103ff57600080fd5b50610282600160a060020a0360043516602435610f8a565b34801561042357600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102ee9436949293602493928401919081908401838280828437509497506111449650505050505050565b34801561047c57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102ee94369492936024939284019190819084018382808284375094975061119d9650505050505050565b3480156104d557600080fd5b506102ee6004356111f1565b3480156104ed57600080fd5b506102ee611342565b61018c600160a060020a036004351661136f565b34801561051657600080fd5b506102ee61137b565b60008060008060008060008060008a6000339050600b60009054906101000a900460ff16801561055f5750662386f26fc100008261055b610dbe565b0311155b1561086b57600160a060020a03811660009081526003602052604090205460ff16151560011480156105b35750600160a060020a038116600090815260076020526040902054662386f26fc1000090830111155b15156105be57600080fd5b600160a060020a0381166000908152600760205260409020546105e19083611431565b600160a060020a0382166000908152600760205260409020553399506106088d600a611447565b9850610615896003611447565b9750610621898961145e565b965061062d8d8a61145e565b955061063886611470565b9450680100000000000000008702935060008511801561066257506008546106608682611431565b115b151561066d57600080fd5b600160a060020a038c1615801590610697575089600160a060020a03168c600160a060020a031614155b80156106bd5750600254600160a060020a038d1660009081526004602052604090205410155b1561070357600160a060020a038c166000908152600560205260409020546106e59089611431565b600160a060020a038d1660009081526005602052604090205561071e565b61070d8789611431565b965068010000000000000000870293505b600060085411156107825761073560085486611431565b600881905568010000000000000000880281151561074f57fe5b6009805492909104909101905560085468010000000000000000880281151561077457fe5b048502840384039350610788565b60088590555b600160a060020a038a166000908152600460205260409020546107ab9086611431565b600460008c600160a060020a0316600160a060020a031681526020019081526020016000208190555083856009540203925082600660008c600160a060020a0316600160a060020a03168152602001908152602001600020600082825401925050819055508b600160a060020a03168a600160a060020a03167f022c0d992e4d873a3748436d960d5140c1f9721cf73f7ca5ec679d3d9f4fe2d58f88604051808381526020018281526020019250505060405180910390a3849a50610ae2565b600b805460ff191690553399506108838d600a611447565b9850610890896003611447565b975061089c898961145e565b96506108a88d8a61145e565b95506108b386611470565b945068010000000000000000870293506000851180156108dd57506008546108db8682611431565b115b15156108e857600080fd5b600160a060020a038c1615801590610912575089600160a060020a03168c600160a060020a031614155b80156109385750600254600160a060020a038d1660009081526004602052604090205410155b1561097e57600160a060020a038c166000908152600560205260409020546109609089611431565b600160a060020a038d16600090815260056020526040902055610999565b6109888789611431565b965068010000000000000000870293505b600060085411156109fd576109b060085486611431565b60088190556801000000000000000088028115156109ca57fe5b600980549290910490910190556008546801000000000000000088028115156109ef57fe5b048502840384039350610a03565b60088590555b600160a060020a038a16600090815260046020526040902054610a269086611431565b600460008c600160a060020a0316600160a060020a031681526020019081526020016000208190555083856009540203925082600660008c600160a060020a0316600160a060020a03168152602001908152602001600020600082825401925050819055508b600160a060020a03168a600160a060020a03167f022c0d992e4d873a3748436d960d5140c1f9721cf73f7ca5ec679d3d9f4fe2d58f88604051808381526020018281526020019250505060405180910390a3849a505b5050505050505050505092915050565b600160a060020a0316600090815260066020908152604080832054600490925290912054600954680100000000000000009102919091030490565b6000805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610bb35780601f10610b8857610100808354040283529160200191610bb3565b820191906000526020600020905b815481529060010190602001808311610b9657829003601f168201915b505050505081565b6000808080610bcb85600a611447565b9250610bd7858461145e565b9150610be282611470565b95945050505050565b6008545b90565b6000806000806008548511151515610c0957600080fd5b610c1285611508565b9250610c1f83600a611447565b9150610be2838361145e565b600b5460ff1681565b601281565b600a6020526000908152604090205460ff1681565b6000806000610c5d6001610d7b565b11610c6757600080fd5b339150610c746000610d7b565b600160a060020a038316600081815260066020908152604080832080546801000000000000000087020190556005909152808220805490839055905193019350909183156108fc0291849190818181858888f19350505050158015610cdd573d6000803e3d6000fd5b50604080518281529051600160a060020a038416917fccad973dcd043c7d680389db4378bd6b9775db7124092e9e0422c9e46d7985dc919081900360200190a25050565b60008060008060085460001415610d3f576414f46b04009350610d6f565b610d50670de0b6b3a7640000611508565b9250610d5d83600a611447565b9150610d69838361145e565b90508093505b50505090565b60025481565b60003382610d9157610d8c81610af2565b610db5565b600160a060020a038116600090815260056020526040902054610db382610af2565b015b91505b50919050565b303190565b600160a060020a031660009081526004602052604090205490565b604080516c010000000000000000000000003390810282528251918290036014019091206000908152600a602052919091205460ff161515610e1f57600080fd5b50600255565b60008060008060085460001415610e435764199c82cc009350610d6f565b610e54670de0b6b3a7640000611508565b9250610e6183600a611447565b9150610d698383611431565b604080516c010000000000000000000000003390810282528251918290036014019091206000908152600a602052919091205460ff161515610eae57600080fd5b506000918252600a6020526040909120805460ff1916911515919091179055565b600033610edb81610dc3565b91505b5090565b60018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610bb35780601f10610b8857610100808354040283529160200191610bb3565b604080516c010000000000000000000000003390810282528251918290036014019091206000908152600a602052919091205460ff161515610f7d57600080fd5b50600b805460ff19169055565b600080600080600080610f9b610ecf565b11610fa557600080fd5b600b5433945060ff16158015610fd35750600160a060020a0384166000908152600460205260409020548611155b1515610fde57600080fd5b6000610fea6001610d7b565b1115610ff857610ff8610c4e565b61100386600a611447565b925061100f868461145e565b915061101a83611508565b90506110286008548461145e565b600855600160a060020a03841660009081526004602052604090205461104e908761145e565b600160a060020a03808616600090815260046020526040808220939093559089168152205461107d9083611431565b600160a060020a0388811660008181526004602090815260408083209590955560098054948a16835260069091528482208054948c029094039093558254918152929092208054928502909201909155546008546110f191906801000000000000000084028115156110eb57fe5b04611431565b600955604080518381529051600160a060020a03808a1692908716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35060019695505050505050565b604080516c010000000000000000000000003390810282528251918290036014019091206000908152600a602052919091205460ff16151561118557600080fd5b81516111989060019060208501906115a9565b505050565b604080516c010000000000000000000000003390810282528251918290036014019091206000908152600a602052919091205460ff1615156111de57600080fd5b81516111989060009060208501906115a9565b6000806000806000806000611204610ecf565b1161120e57600080fd5b3360008181526004602052604090205490965087111561122d57600080fd5b86945061123985611508565b935061124684600a611447565b9250611252848461145e565b91506112606008548661145e565b600855600160a060020a038616600090815260046020526040902054611286908661145e565b600160a060020a038716600090815260046020908152604080832093909355600954600690915291812080549288026801000000000000000086020192839003905560085491925010156112f6576112f26009546008546801000000000000000086028115156110eb57fe5b6009555b60408051868152602081018490528151600160a060020a038916927fc4823739c5787d2ca17e404aa47d5569ae71dfb49cbf21b3f6152ed238a31139928290030190a250505050505050565b336000818152600460205260408120549081111561136357611363816111f1565b61136b610c4e565b5050565b6000610db8348361051f565b60008060008061138b6001610d7b565b1161139557600080fd5b61139f6000610d7b565b336000818152600660209081526040808320805468010000000000000000870201905560059091528120805490829055909201945092506113e190849061051f565b905081600160a060020a03167fbe339fc14b041c2b0e0f3dd2cd325d0c3668b78378001e53160eab36153264588483604051808381526020018281526020019250505060405180910390a2505050565b60008282018381101561144057fe5b9392505050565b600080828481151561145557fe5b04949350505050565b60008282111561146a57fe5b50900390565b6008546000906c01431e0fae6d7217caa00000009082906402540be4006114f56114ef730380d4bd8a8678c1bb542c80deb4800000000000880268056bc75e2d631000006002860a02017005e0a1fd2712875988becaad0000000000850201780197d4df19d605767337e9f14d3eec8920e40000000000000001611574565b8561145e565b8115156114fe57fe5b0403949350505050565b600854600090670de0b6b3a76400008381019181019083906115616414f46b04008285046402540be40002018702600283670de0b6b3a763ffff1982890a8b900301046402540be4000281151561155b57fe5b0461145e565b81151561156a57fe5b0495945050505050565b80600260018201045b81811015610db857809150600281828581151561159657fe5b04018115156115a157fe5b04905061157d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115ea57805160ff1916838001178555611617565b82800160010185558215611617579182015b828111156116175782518255916020019190600101906115fc565b50610ede92610bef9250905b80821115610ede57600081556001016116235600a165627a7a72305820eaff89560df403a670d2e1cdd20ccb9d7ca32638858384078ac2ef243093f7660029`

// DeployPoSL deploys a new Ethereum contract, binding an instance of PoSL to it.
func DeployPoSL(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PoSL, error) {
	parsed, err := abi.JSON(strings.NewReader(PoSLABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PoSLBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoSL{PoSLCaller: PoSLCaller{contract: contract}, PoSLTransactor: PoSLTransactor{contract: contract}, PoSLFilterer: PoSLFilterer{contract: contract}}, nil
}

// PoSL is an auto generated Go binding around an Ethereum contract.
type PoSL struct {
	PoSLCaller     // Read-only binding to the contract
	PoSLTransactor // Write-only binding to the contract
	PoSLFilterer   // Log filterer for contract events
}

// PoSLCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoSLCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoSLTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoSLTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoSLFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoSLFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoSLSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoSLSession struct {
	Contract     *PoSL             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoSLCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoSLCallerSession struct {
	Contract *PoSLCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoSLTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoSLTransactorSession struct {
	Contract     *PoSLTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoSLRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoSLRaw struct {
	Contract *PoSL // Generic contract binding to access the raw methods on
}

// PoSLCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoSLCallerRaw struct {
	Contract *PoSLCaller // Generic read-only contract binding to access the raw methods on
}

// PoSLTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoSLTransactorRaw struct {
	Contract *PoSLTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoSL creates a new instance of PoSL, bound to a specific deployed contract.
func NewPoSL(address common.Address, backend bind.ContractBackend) (*PoSL, error) {
	contract, err := bindPoSL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoSL{PoSLCaller: PoSLCaller{contract: contract}, PoSLTransactor: PoSLTransactor{contract: contract}, PoSLFilterer: PoSLFilterer{contract: contract}}, nil
}

// NewPoSLCaller creates a new read-only instance of PoSL, bound to a specific deployed contract.
func NewPoSLCaller(address common.Address, caller bind.ContractCaller) (*PoSLCaller, error) {
	contract, err := bindPoSL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoSLCaller{contract: contract}, nil
}

// NewPoSLTransactor creates a new write-only instance of PoSL, bound to a specific deployed contract.
func NewPoSLTransactor(address common.Address, transactor bind.ContractTransactor) (*PoSLTransactor, error) {
	contract, err := bindPoSL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoSLTransactor{contract: contract}, nil
}

// NewPoSLFilterer creates a new log filterer instance of PoSL, bound to a specific deployed contract.
func NewPoSLFilterer(address common.Address, filterer bind.ContractFilterer) (*PoSLFilterer, error) {
	contract, err := bindPoSL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoSLFilterer{contract: contract}, nil
}

// bindPoSL binds a generic wrapper to an already deployed contract.
func bindPoSL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoSLABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoSL *PoSLRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PoSL.Contract.PoSLCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoSL *PoSLRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSL.Contract.PoSLTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoSL *PoSLRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoSL.Contract.PoSLTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoSL *PoSLCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PoSL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoSL *PoSLTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoSL *PoSLTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoSL.Contract.contract.Transact(opts, method, params...)
}

// Administrators is a free data retrieval call binding the contract method 0x392efb52.
//
// Solidity: function administrators( bytes32) constant returns(bool)
func (_PoSL *PoSLCaller) Administrators(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "administrators", arg0)
	return *ret0, err
}

// Administrators is a free data retrieval call binding the contract method 0x392efb52.
//
// Solidity: function administrators( bytes32) constant returns(bool)
func (_PoSL *PoSLSession) Administrators(arg0 [32]byte) (bool, error) {
	return _PoSL.Contract.Administrators(&_PoSL.CallOpts, arg0)
}

// Administrators is a free data retrieval call binding the contract method 0x392efb52.
//
// Solidity: function administrators( bytes32) constant returns(bool)
func (_PoSL *PoSLCallerSession) Administrators(arg0 [32]byte) (bool, error) {
	return _PoSL.Contract.Administrators(&_PoSL.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_customerAddress address) constant returns(uint256)
func (_PoSL *PoSLCaller) BalanceOf(opts *bind.CallOpts, _customerAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "balanceOf", _customerAddress)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_customerAddress address) constant returns(uint256)
func (_PoSL *PoSLSession) BalanceOf(_customerAddress common.Address) (*big.Int, error) {
	return _PoSL.Contract.BalanceOf(&_PoSL.CallOpts, _customerAddress)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_customerAddress address) constant returns(uint256)
func (_PoSL *PoSLCallerSession) BalanceOf(_customerAddress common.Address) (*big.Int, error) {
	return _PoSL.Contract.BalanceOf(&_PoSL.CallOpts, _customerAddress)
}

// BuyPrice is a free data retrieval call binding the contract method 0x8620410b.
//
// Solidity: function buyPrice() constant returns(uint256)
func (_PoSL *PoSLCaller) BuyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "buyPrice")
	return *ret0, err
}

// BuyPrice is a free data retrieval call binding the contract method 0x8620410b.
//
// Solidity: function buyPrice() constant returns(uint256)
func (_PoSL *PoSLSession) BuyPrice() (*big.Int, error) {
	return _PoSL.Contract.BuyPrice(&_PoSL.CallOpts)
}

// BuyPrice is a free data retrieval call binding the contract method 0x8620410b.
//
// Solidity: function buyPrice() constant returns(uint256)
func (_PoSL *PoSLCallerSession) BuyPrice() (*big.Int, error) {
	return _PoSL.Contract.BuyPrice(&_PoSL.CallOpts)
}

// CalculateEthereumReceived is a free data retrieval call binding the contract method 0x22609373.
//
// Solidity: function calculateEthereumReceived(_tokensToSell uint256) constant returns(uint256)
func (_PoSL *PoSLCaller) CalculateEthereumReceived(opts *bind.CallOpts, _tokensToSell *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "calculateEthereumReceived", _tokensToSell)
	return *ret0, err
}

// CalculateEthereumReceived is a free data retrieval call binding the contract method 0x22609373.
//
// Solidity: function calculateEthereumReceived(_tokensToSell uint256) constant returns(uint256)
func (_PoSL *PoSLSession) CalculateEthereumReceived(_tokensToSell *big.Int) (*big.Int, error) {
	return _PoSL.Contract.CalculateEthereumReceived(&_PoSL.CallOpts, _tokensToSell)
}

// CalculateEthereumReceived is a free data retrieval call binding the contract method 0x22609373.
//
// Solidity: function calculateEthereumReceived(_tokensToSell uint256) constant returns(uint256)
func (_PoSL *PoSLCallerSession) CalculateEthereumReceived(_tokensToSell *big.Int) (*big.Int, error) {
	return _PoSL.Contract.CalculateEthereumReceived(&_PoSL.CallOpts, _tokensToSell)
}

// CalculateTokensReceived is a free data retrieval call binding the contract method 0x10d0ffdd.
//
// Solidity: function calculateTokensReceived(_ethereumToSpend uint256) constant returns(uint256)
func (_PoSL *PoSLCaller) CalculateTokensReceived(opts *bind.CallOpts, _ethereumToSpend *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "calculateTokensReceived", _ethereumToSpend)
	return *ret0, err
}

// CalculateTokensReceived is a free data retrieval call binding the contract method 0x10d0ffdd.
//
// Solidity: function calculateTokensReceived(_ethereumToSpend uint256) constant returns(uint256)
func (_PoSL *PoSLSession) CalculateTokensReceived(_ethereumToSpend *big.Int) (*big.Int, error) {
	return _PoSL.Contract.CalculateTokensReceived(&_PoSL.CallOpts, _ethereumToSpend)
}

// CalculateTokensReceived is a free data retrieval call binding the contract method 0x10d0ffdd.
//
// Solidity: function calculateTokensReceived(_ethereumToSpend uint256) constant returns(uint256)
func (_PoSL *PoSLCallerSession) CalculateTokensReceived(_ethereumToSpend *big.Int) (*big.Int, error) {
	return _PoSL.Contract.CalculateTokensReceived(&_PoSL.CallOpts, _ethereumToSpend)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PoSL *PoSLCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PoSL *PoSLSession) Decimals() (uint8, error) {
	return _PoSL.Contract.Decimals(&_PoSL.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PoSL *PoSLCallerSession) Decimals() (uint8, error) {
	return _PoSL.Contract.Decimals(&_PoSL.CallOpts)
}

// DividendsOf is a free data retrieval call binding the contract method 0x0065318b.
//
// Solidity: function dividendsOf(_customerAddress address) constant returns(uint256)
func (_PoSL *PoSLCaller) DividendsOf(opts *bind.CallOpts, _customerAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "dividendsOf", _customerAddress)
	return *ret0, err
}

// DividendsOf is a free data retrieval call binding the contract method 0x0065318b.
//
// Solidity: function dividendsOf(_customerAddress address) constant returns(uint256)
func (_PoSL *PoSLSession) DividendsOf(_customerAddress common.Address) (*big.Int, error) {
	return _PoSL.Contract.DividendsOf(&_PoSL.CallOpts, _customerAddress)
}

// DividendsOf is a free data retrieval call binding the contract method 0x0065318b.
//
// Solidity: function dividendsOf(_customerAddress address) constant returns(uint256)
func (_PoSL *PoSLCallerSession) DividendsOf(_customerAddress common.Address) (*big.Int, error) {
	return _PoSL.Contract.DividendsOf(&_PoSL.CallOpts, _customerAddress)
}

// MyDividends is a free data retrieval call binding the contract method 0x688abbf7.
//
// Solidity: function myDividends(_includeReferralBonus bool) constant returns(uint256)
func (_PoSL *PoSLCaller) MyDividends(opts *bind.CallOpts, _includeReferralBonus bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "myDividends", _includeReferralBonus)
	return *ret0, err
}

// MyDividends is a free data retrieval call binding the contract method 0x688abbf7.
//
// Solidity: function myDividends(_includeReferralBonus bool) constant returns(uint256)
func (_PoSL *PoSLSession) MyDividends(_includeReferralBonus bool) (*big.Int, error) {
	return _PoSL.Contract.MyDividends(&_PoSL.CallOpts, _includeReferralBonus)
}

// MyDividends is a free data retrieval call binding the contract method 0x688abbf7.
//
// Solidity: function myDividends(_includeReferralBonus bool) constant returns(uint256)
func (_PoSL *PoSLCallerSession) MyDividends(_includeReferralBonus bool) (*big.Int, error) {
	return _PoSL.Contract.MyDividends(&_PoSL.CallOpts, _includeReferralBonus)
}

// MyTokens is a free data retrieval call binding the contract method 0x949e8acd.
//
// Solidity: function myTokens() constant returns(uint256)
func (_PoSL *PoSLCaller) MyTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "myTokens")
	return *ret0, err
}

// MyTokens is a free data retrieval call binding the contract method 0x949e8acd.
//
// Solidity: function myTokens() constant returns(uint256)
func (_PoSL *PoSLSession) MyTokens() (*big.Int, error) {
	return _PoSL.Contract.MyTokens(&_PoSL.CallOpts)
}

// MyTokens is a free data retrieval call binding the contract method 0x949e8acd.
//
// Solidity: function myTokens() constant returns(uint256)
func (_PoSL *PoSLCallerSession) MyTokens() (*big.Int, error) {
	return _PoSL.Contract.MyTokens(&_PoSL.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PoSL *PoSLCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PoSL *PoSLSession) Name() (string, error) {
	return _PoSL.Contract.Name(&_PoSL.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PoSL *PoSLCallerSession) Name() (string, error) {
	return _PoSL.Contract.Name(&_PoSL.CallOpts)
}

// OnlyAmbassadors is a free data retrieval call binding the contract method 0x27defa1f.
//
// Solidity: function onlyAmbassadors() constant returns(bool)
func (_PoSL *PoSLCaller) OnlyAmbassadors(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "onlyAmbassadors")
	return *ret0, err
}

// OnlyAmbassadors is a free data retrieval call binding the contract method 0x27defa1f.
//
// Solidity: function onlyAmbassadors() constant returns(bool)
func (_PoSL *PoSLSession) OnlyAmbassadors() (bool, error) {
	return _PoSL.Contract.OnlyAmbassadors(&_PoSL.CallOpts)
}

// OnlyAmbassadors is a free data retrieval call binding the contract method 0x27defa1f.
//
// Solidity: function onlyAmbassadors() constant returns(bool)
func (_PoSL *PoSLCallerSession) OnlyAmbassadors() (bool, error) {
	return _PoSL.Contract.OnlyAmbassadors(&_PoSL.CallOpts)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_PoSL *PoSLCaller) SellPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "sellPrice")
	return *ret0, err
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_PoSL *PoSLSession) SellPrice() (*big.Int, error) {
	return _PoSL.Contract.SellPrice(&_PoSL.CallOpts)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_PoSL *PoSLCallerSession) SellPrice() (*big.Int, error) {
	return _PoSL.Contract.SellPrice(&_PoSL.CallOpts)
}

// StakingRequirement is a free data retrieval call binding the contract method 0x56d399e8.
//
// Solidity: function stakingRequirement() constant returns(uint256)
func (_PoSL *PoSLCaller) StakingRequirement(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "stakingRequirement")
	return *ret0, err
}

// StakingRequirement is a free data retrieval call binding the contract method 0x56d399e8.
//
// Solidity: function stakingRequirement() constant returns(uint256)
func (_PoSL *PoSLSession) StakingRequirement() (*big.Int, error) {
	return _PoSL.Contract.StakingRequirement(&_PoSL.CallOpts)
}

// StakingRequirement is a free data retrieval call binding the contract method 0x56d399e8.
//
// Solidity: function stakingRequirement() constant returns(uint256)
func (_PoSL *PoSLCallerSession) StakingRequirement() (*big.Int, error) {
	return _PoSL.Contract.StakingRequirement(&_PoSL.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PoSL *PoSLCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PoSL *PoSLSession) Symbol() (string, error) {
	return _PoSL.Contract.Symbol(&_PoSL.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PoSL *PoSLCallerSession) Symbol() (string, error) {
	return _PoSL.Contract.Symbol(&_PoSL.CallOpts)
}

// TotalEthereumBalance is a free data retrieval call binding the contract method 0x6b2f4632.
//
// Solidity: function totalEthereumBalance() constant returns(uint256)
func (_PoSL *PoSLCaller) TotalEthereumBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "totalEthereumBalance")
	return *ret0, err
}

// TotalEthereumBalance is a free data retrieval call binding the contract method 0x6b2f4632.
//
// Solidity: function totalEthereumBalance() constant returns(uint256)
func (_PoSL *PoSLSession) TotalEthereumBalance() (*big.Int, error) {
	return _PoSL.Contract.TotalEthereumBalance(&_PoSL.CallOpts)
}

// TotalEthereumBalance is a free data retrieval call binding the contract method 0x6b2f4632.
//
// Solidity: function totalEthereumBalance() constant returns(uint256)
func (_PoSL *PoSLCallerSession) TotalEthereumBalance() (*big.Int, error) {
	return _PoSL.Contract.TotalEthereumBalance(&_PoSL.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PoSL *PoSLCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PoSL.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PoSL *PoSLSession) TotalSupply() (*big.Int, error) {
	return _PoSL.Contract.TotalSupply(&_PoSL.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PoSL *PoSLCallerSession) TotalSupply() (*big.Int, error) {
	return _PoSL.Contract.TotalSupply(&_PoSL.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(_referredBy address) returns(uint256)
func (_PoSL *PoSLTransactor) Buy(opts *bind.TransactOpts, _referredBy common.Address) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "buy", _referredBy)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(_referredBy address) returns(uint256)
func (_PoSL *PoSLSession) Buy(_referredBy common.Address) (*types.Transaction, error) {
	return _PoSL.Contract.Buy(&_PoSL.TransactOpts, _referredBy)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(_referredBy address) returns(uint256)
func (_PoSL *PoSLTransactorSession) Buy(_referredBy common.Address) (*types.Transaction, error) {
	return _PoSL.Contract.Buy(&_PoSL.TransactOpts, _referredBy)
}

// DisableInitialStage is a paid mutator transaction binding the contract method 0xa8e04f34.
//
// Solidity: function disableInitialStage() returns()
func (_PoSL *PoSLTransactor) DisableInitialStage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "disableInitialStage")
}

// DisableInitialStage is a paid mutator transaction binding the contract method 0xa8e04f34.
//
// Solidity: function disableInitialStage() returns()
func (_PoSL *PoSLSession) DisableInitialStage() (*types.Transaction, error) {
	return _PoSL.Contract.DisableInitialStage(&_PoSL.TransactOpts)
}

// DisableInitialStage is a paid mutator transaction binding the contract method 0xa8e04f34.
//
// Solidity: function disableInitialStage() returns()
func (_PoSL *PoSLTransactorSession) DisableInitialStage() (*types.Transaction, error) {
	return _PoSL.Contract.DisableInitialStage(&_PoSL.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_PoSL *PoSLTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_PoSL *PoSLSession) Exit() (*types.Transaction, error) {
	return _PoSL.Contract.Exit(&_PoSL.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_PoSL *PoSLTransactorSession) Exit() (*types.Transaction, error) {
	return _PoSL.Contract.Exit(&_PoSL.TransactOpts)
}

// Reinvest is a paid mutator transaction binding the contract method 0xfdb5a03e.
//
// Solidity: function reinvest() returns()
func (_PoSL *PoSLTransactor) Reinvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "reinvest")
}

// Reinvest is a paid mutator transaction binding the contract method 0xfdb5a03e.
//
// Solidity: function reinvest() returns()
func (_PoSL *PoSLSession) Reinvest() (*types.Transaction, error) {
	return _PoSL.Contract.Reinvest(&_PoSL.TransactOpts)
}

// Reinvest is a paid mutator transaction binding the contract method 0xfdb5a03e.
//
// Solidity: function reinvest() returns()
func (_PoSL *PoSLTransactorSession) Reinvest() (*types.Transaction, error) {
	return _PoSL.Contract.Reinvest(&_PoSL.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(_amountOfTokens uint256) returns()
func (_PoSL *PoSLTransactor) Sell(opts *bind.TransactOpts, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "sell", _amountOfTokens)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(_amountOfTokens uint256) returns()
func (_PoSL *PoSLSession) Sell(_amountOfTokens *big.Int) (*types.Transaction, error) {
	return _PoSL.Contract.Sell(&_PoSL.TransactOpts, _amountOfTokens)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(_amountOfTokens uint256) returns()
func (_PoSL *PoSLTransactorSession) Sell(_amountOfTokens *big.Int) (*types.Transaction, error) {
	return _PoSL.Contract.Sell(&_PoSL.TransactOpts, _amountOfTokens)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0x89135ae9.
//
// Solidity: function setAdministrator(_identifier bytes32, _status bool) returns()
func (_PoSL *PoSLTransactor) SetAdministrator(opts *bind.TransactOpts, _identifier [32]byte, _status bool) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "setAdministrator", _identifier, _status)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0x89135ae9.
//
// Solidity: function setAdministrator(_identifier bytes32, _status bool) returns()
func (_PoSL *PoSLSession) SetAdministrator(_identifier [32]byte, _status bool) (*types.Transaction, error) {
	return _PoSL.Contract.SetAdministrator(&_PoSL.TransactOpts, _identifier, _status)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0x89135ae9.
//
// Solidity: function setAdministrator(_identifier bytes32, _status bool) returns()
func (_PoSL *PoSLTransactorSession) SetAdministrator(_identifier [32]byte, _status bool) (*types.Transaction, error) {
	return _PoSL.Contract.SetAdministrator(&_PoSL.TransactOpts, _identifier, _status)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns()
func (_PoSL *PoSLTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns()
func (_PoSL *PoSLSession) SetName(_name string) (*types.Transaction, error) {
	return _PoSL.Contract.SetName(&_PoSL.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns()
func (_PoSL *PoSLTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _PoSL.Contract.SetName(&_PoSL.TransactOpts, _name)
}

// SetStakingRequirement is a paid mutator transaction binding the contract method 0x8328b610.
//
// Solidity: function setStakingRequirement(_amountOfTokens uint256) returns()
func (_PoSL *PoSLTransactor) SetStakingRequirement(opts *bind.TransactOpts, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "setStakingRequirement", _amountOfTokens)
}

// SetStakingRequirement is a paid mutator transaction binding the contract method 0x8328b610.
//
// Solidity: function setStakingRequirement(_amountOfTokens uint256) returns()
func (_PoSL *PoSLSession) SetStakingRequirement(_amountOfTokens *big.Int) (*types.Transaction, error) {
	return _PoSL.Contract.SetStakingRequirement(&_PoSL.TransactOpts, _amountOfTokens)
}

// SetStakingRequirement is a paid mutator transaction binding the contract method 0x8328b610.
//
// Solidity: function setStakingRequirement(_amountOfTokens uint256) returns()
func (_PoSL *PoSLTransactorSession) SetStakingRequirement(_amountOfTokens *big.Int) (*types.Transaction, error) {
	return _PoSL.Contract.SetStakingRequirement(&_PoSL.TransactOpts, _amountOfTokens)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(_symbol string) returns()
func (_PoSL *PoSLTransactor) SetSymbol(opts *bind.TransactOpts, _symbol string) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "setSymbol", _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(_symbol string) returns()
func (_PoSL *PoSLSession) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _PoSL.Contract.SetSymbol(&_PoSL.TransactOpts, _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(_symbol string) returns()
func (_PoSL *PoSLTransactorSession) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _PoSL.Contract.SetSymbol(&_PoSL.TransactOpts, _symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_toAddress address, _amountOfTokens uint256) returns(bool)
func (_PoSL *PoSLTransactor) Transfer(opts *bind.TransactOpts, _toAddress common.Address, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "transfer", _toAddress, _amountOfTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_toAddress address, _amountOfTokens uint256) returns(bool)
func (_PoSL *PoSLSession) Transfer(_toAddress common.Address, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _PoSL.Contract.Transfer(&_PoSL.TransactOpts, _toAddress, _amountOfTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_toAddress address, _amountOfTokens uint256) returns(bool)
func (_PoSL *PoSLTransactorSession) Transfer(_toAddress common.Address, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _PoSL.Contract.Transfer(&_PoSL.TransactOpts, _toAddress, _amountOfTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PoSL *PoSLTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSL.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PoSL *PoSLSession) Withdraw() (*types.Transaction, error) {
	return _PoSL.Contract.Withdraw(&_PoSL.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PoSL *PoSLTransactorSession) Withdraw() (*types.Transaction, error) {
	return _PoSL.Contract.Withdraw(&_PoSL.TransactOpts)
}

// PoSLTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PoSL contract.
type PoSLTransferIterator struct {
	Event *PoSLTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoSLTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSLTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoSLTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoSLTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSLTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSLTransfer represents a Transfer event raised by the PoSL contract.
type PoSLTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, tokens uint256)
func (_PoSL *PoSLFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoSLTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoSL.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoSLTransferIterator{contract: _PoSL.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, tokens uint256)
func (_PoSL *PoSLFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoSLTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoSL.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSLTransfer)
				if err := _PoSL.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PoSLOnReinvestmentIterator is returned from FilterOnReinvestment and is used to iterate over the raw logs and unpacked data for OnReinvestment events raised by the PoSL contract.
type PoSLOnReinvestmentIterator struct {
	Event *PoSLOnReinvestment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoSLOnReinvestmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSLOnReinvestment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoSLOnReinvestment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoSLOnReinvestmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSLOnReinvestmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSLOnReinvestment represents a OnReinvestment event raised by the PoSL contract.
type PoSLOnReinvestment struct {
	CustomerAddress    common.Address
	EthereumReinvested *big.Int
	TokensMinted       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnReinvestment is a free log retrieval operation binding the contract event 0xbe339fc14b041c2b0e0f3dd2cd325d0c3668b78378001e53160eab3615326458.
//
// Solidity: event onReinvestment(customerAddress indexed address, ethereumReinvested uint256, tokensMinted uint256)
func (_PoSL *PoSLFilterer) FilterOnReinvestment(opts *bind.FilterOpts, customerAddress []common.Address) (*PoSLOnReinvestmentIterator, error) {

	var customerAddressRule []interface{}
	for _, customerAddressItem := range customerAddress {
		customerAddressRule = append(customerAddressRule, customerAddressItem)
	}

	logs, sub, err := _PoSL.contract.FilterLogs(opts, "onReinvestment", customerAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoSLOnReinvestmentIterator{contract: _PoSL.contract, event: "onReinvestment", logs: logs, sub: sub}, nil
}

// WatchOnReinvestment is a free log subscription operation binding the contract event 0xbe339fc14b041c2b0e0f3dd2cd325d0c3668b78378001e53160eab3615326458.
//
// Solidity: event onReinvestment(customerAddress indexed address, ethereumReinvested uint256, tokensMinted uint256)
func (_PoSL *PoSLFilterer) WatchOnReinvestment(opts *bind.WatchOpts, sink chan<- *PoSLOnReinvestment, customerAddress []common.Address) (event.Subscription, error) {

	var customerAddressRule []interface{}
	for _, customerAddressItem := range customerAddress {
		customerAddressRule = append(customerAddressRule, customerAddressItem)
	}

	logs, sub, err := _PoSL.contract.WatchLogs(opts, "onReinvestment", customerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSLOnReinvestment)
				if err := _PoSL.contract.UnpackLog(event, "onReinvestment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PoSLOnTokenPurchaseIterator is returned from FilterOnTokenPurchase and is used to iterate over the raw logs and unpacked data for OnTokenPurchase events raised by the PoSL contract.
type PoSLOnTokenPurchaseIterator struct {
	Event *PoSLOnTokenPurchase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoSLOnTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSLOnTokenPurchase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoSLOnTokenPurchase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoSLOnTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSLOnTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSLOnTokenPurchase represents a OnTokenPurchase event raised by the PoSL contract.
type PoSLOnTokenPurchase struct {
	CustomerAddress  common.Address
	IncomingEthereum *big.Int
	TokensMinted     *big.Int
	ReferredBy       common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnTokenPurchase is a free log retrieval operation binding the contract event 0x022c0d992e4d873a3748436d960d5140c1f9721cf73f7ca5ec679d3d9f4fe2d5.
//
// Solidity: event onTokenPurchase(customerAddress indexed address, incomingEthereum uint256, tokensMinted uint256, referredBy indexed address)
func (_PoSL *PoSLFilterer) FilterOnTokenPurchase(opts *bind.FilterOpts, customerAddress []common.Address, referredBy []common.Address) (*PoSLOnTokenPurchaseIterator, error) {

	var customerAddressRule []interface{}
	for _, customerAddressItem := range customerAddress {
		customerAddressRule = append(customerAddressRule, customerAddressItem)
	}

	var referredByRule []interface{}
	for _, referredByItem := range referredBy {
		referredByRule = append(referredByRule, referredByItem)
	}

	logs, sub, err := _PoSL.contract.FilterLogs(opts, "onTokenPurchase", customerAddressRule, referredByRule)
	if err != nil {
		return nil, err
	}
	return &PoSLOnTokenPurchaseIterator{contract: _PoSL.contract, event: "onTokenPurchase", logs: logs, sub: sub}, nil
}

// WatchOnTokenPurchase is a free log subscription operation binding the contract event 0x022c0d992e4d873a3748436d960d5140c1f9721cf73f7ca5ec679d3d9f4fe2d5.
//
// Solidity: event onTokenPurchase(customerAddress indexed address, incomingEthereum uint256, tokensMinted uint256, referredBy indexed address)
func (_PoSL *PoSLFilterer) WatchOnTokenPurchase(opts *bind.WatchOpts, sink chan<- *PoSLOnTokenPurchase, customerAddress []common.Address, referredBy []common.Address) (event.Subscription, error) {

	var customerAddressRule []interface{}
	for _, customerAddressItem := range customerAddress {
		customerAddressRule = append(customerAddressRule, customerAddressItem)
	}

	var referredByRule []interface{}
	for _, referredByItem := range referredBy {
		referredByRule = append(referredByRule, referredByItem)
	}

	logs, sub, err := _PoSL.contract.WatchLogs(opts, "onTokenPurchase", customerAddressRule, referredByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSLOnTokenPurchase)
				if err := _PoSL.contract.UnpackLog(event, "onTokenPurchase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PoSLOnTokenSellIterator is returned from FilterOnTokenSell and is used to iterate over the raw logs and unpacked data for OnTokenSell events raised by the PoSL contract.
type PoSLOnTokenSellIterator struct {
	Event *PoSLOnTokenSell // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoSLOnTokenSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSLOnTokenSell)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoSLOnTokenSell)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoSLOnTokenSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSLOnTokenSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSLOnTokenSell represents a OnTokenSell event raised by the PoSL contract.
type PoSLOnTokenSell struct {
	CustomerAddress common.Address
	TokensBurned    *big.Int
	EthereumEarned  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOnTokenSell is a free log retrieval operation binding the contract event 0xc4823739c5787d2ca17e404aa47d5569ae71dfb49cbf21b3f6152ed238a31139.
//
// Solidity: event onTokenSell(customerAddress indexed address, tokensBurned uint256, ethereumEarned uint256)
func (_PoSL *PoSLFilterer) FilterOnTokenSell(opts *bind.FilterOpts, customerAddress []common.Address) (*PoSLOnTokenSellIterator, error) {

	var customerAddressRule []interface{}
	for _, customerAddressItem := range customerAddress {
		customerAddressRule = append(customerAddressRule, customerAddressItem)
	}

	logs, sub, err := _PoSL.contract.FilterLogs(opts, "onTokenSell", customerAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoSLOnTokenSellIterator{contract: _PoSL.contract, event: "onTokenSell", logs: logs, sub: sub}, nil
}

// WatchOnTokenSell is a free log subscription operation binding the contract event 0xc4823739c5787d2ca17e404aa47d5569ae71dfb49cbf21b3f6152ed238a31139.
//
// Solidity: event onTokenSell(customerAddress indexed address, tokensBurned uint256, ethereumEarned uint256)
func (_PoSL *PoSLFilterer) WatchOnTokenSell(opts *bind.WatchOpts, sink chan<- *PoSLOnTokenSell, customerAddress []common.Address) (event.Subscription, error) {

	var customerAddressRule []interface{}
	for _, customerAddressItem := range customerAddress {
		customerAddressRule = append(customerAddressRule, customerAddressItem)
	}

	logs, sub, err := _PoSL.contract.WatchLogs(opts, "onTokenSell", customerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSLOnTokenSell)
				if err := _PoSL.contract.UnpackLog(event, "onTokenSell", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PoSLOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the PoSL contract.
type PoSLOnWithdrawIterator struct {
	Event *PoSLOnWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoSLOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSLOnWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoSLOnWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoSLOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSLOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSLOnWithdraw represents a OnWithdraw event raised by the PoSL contract.
type PoSLOnWithdraw struct {
	CustomerAddress   common.Address
	EthereumWithdrawn *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0xccad973dcd043c7d680389db4378bd6b9775db7124092e9e0422c9e46d7985dc.
//
// Solidity: event onWithdraw(customerAddress indexed address, ethereumWithdrawn uint256)
func (_PoSL *PoSLFilterer) FilterOnWithdraw(opts *bind.FilterOpts, customerAddress []common.Address) (*PoSLOnWithdrawIterator, error) {

	var customerAddressRule []interface{}
	for _, customerAddressItem := range customerAddress {
		customerAddressRule = append(customerAddressRule, customerAddressItem)
	}

	logs, sub, err := _PoSL.contract.FilterLogs(opts, "onWithdraw", customerAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoSLOnWithdrawIterator{contract: _PoSL.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0xccad973dcd043c7d680389db4378bd6b9775db7124092e9e0422c9e46d7985dc.
//
// Solidity: event onWithdraw(customerAddress indexed address, ethereumWithdrawn uint256)
func (_PoSL *PoSLFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *PoSLOnWithdraw, customerAddress []common.Address) (event.Subscription, error) {

	var customerAddressRule []interface{}
	for _, customerAddressItem := range customerAddress {
		customerAddressRule = append(customerAddressRule, customerAddressItem)
	}

	logs, sub, err := _PoSL.contract.WatchLogs(opts, "onWithdraw", customerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSLOnWithdraw)
				if err := _PoSL.contract.UnpackLog(event, "onWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820390ab5a58934471475a75f0593cceaf4a891c1cd7d08d55ec4fb2d31d7cf504a0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
