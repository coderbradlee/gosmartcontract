# abigen --sol=cryptokitties.sol --pkg=cryptokitties --out=cryptokitties.go

go test -v -test.run TestMix
# go test -v -test.run TestDeploy
# go test -v -test.run TestSet
# go test -v -test.run TestCreate
# 主合约：KittyCore
# 猫的交配拍卖逻辑合约：SiringClockAuction
# 猫的转让拍卖逻辑合约：SaleClockAuction
# 基因合成合约：GeneScienceInterface(代码未开源)
# 猫的Metadata合约：ERC721Metadata (好像并没有什么用)