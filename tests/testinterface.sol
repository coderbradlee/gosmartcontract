

pragma solidity ^0.4.24;
interface DiviesInterface {
    function deposit() external payable returns(uint256);
}
contract Divies{
    function deposit()
        external
        payable
        returns(uint256)
    {
        return 11258963;
    }
}
contract FoMo3Dlong {
    DiviesInterface constant private Di = DiviesInterface(0x0);
    function what()public returns(uint256){
        return Di.deposit();
    }
}