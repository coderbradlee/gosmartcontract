pragma solidity ^0.4.24;
contract calletclottery {
    
    etclotteryInterface constant private etclottery = etclotteryInterface(0xd0c696767a2053d2f4dDF89bA894973D2b026834);
    function()
        public
        payable
    {
       
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