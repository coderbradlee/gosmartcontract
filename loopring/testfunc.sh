abigen --sol=LoopringProtocolImpl.sol --pkg=loopring --out=LoopringProtocolImpl.go
abigen --sol=TokenFactoryImpl.sol --pkg=loopring --out=TokenFactoryImpl.go
abigen --sol=TokenRegistryImpl.sol --pkg=loopring --out=TokenRegistryImpl.go
abigen --sol=TokenTransferDelegateImpl.sol --pkg=loopring --out=TokenTransferDelegateImpl.go
abigen --sol=TransferableMultsigImpl.sol --pkg=loopring --out=TransferableMultsigImpl.go
#go test -v -test.run TestDeploy