pragma solidity ^0.4.24;
contract calletclottery {
    
    etclotteryInterface private etclottery;
    function()
        public
        payable
    {
       
    }
    function setup(address addr)
        external
    {
        etclottery = etclotteryInterface(addr);
    }
    function callByFun(address addr)returns (bool){
        bytes4 methodId = bytes4(keccak256("buy(uint8)"));
        return addr.call(methodId, 1);
    }
    function delegatecallByFun(address addr)returns (bool){
        bytes4 methodId = bytes4(keccak256("buy(uint8)"));
        return addr.delegatecall(methodId, 1);
    }
    function testbuy()
        public
    {
        etclottery.buy(1);

    }
    function getEndowmentBalance() constant public returns (uint)
    {
    	return address(this).balance;
    }
}
interface etclotteryInterface {
    function buy(uint8 _team)
        public
        payable;
    function getFee()
        public
        view
        returns(uint256);
    function getBlock()
        public;
    function withdrawFee(uint256 amount)
        public;
    function playerWithdraw(uint256 amount)
        public;
    function reinvest(uint256 amount, uint8 _team)
        public; 
}