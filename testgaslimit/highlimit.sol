pragma solidity ^0.4.24;


contract highlimit {
    
    event onDel
    (
        uint a1,
        uint a2,
        uint a3,
        uint a4,
        uint a5,
        uint a6,
        uint a17
    );
    function run()
        public
    {
        assert(false);
        // for(uint256 i=0;i<1980000;i++){
        //     i+i;
        // }
    }
    function del()
        public
        returns(uint,uint,uint,uint,uint,uint,uint)
    {
        uint[] arr;
        uint i;
        for(i=0;i<10;i++){
            arr.push(i+2);
        }
        uint len=arr.length-5;
        for(i=0;i<len;i++){
            delete arr[i];
        }
        emit onDel
        (
            arr[4],
            arr[5],
            arr[6],
            arr[7],
            arr[8],
            arr[9],
            arr.length
        );
    }
}
