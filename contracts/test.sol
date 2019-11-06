pragma solidity ^0.5.1;

contract test {

    uint public myPublicVariable;

    event ValueSet(uint newValue, uint date);

    function setMyVariable(uint newValue) public {
        myPublicVariable = newValue;
        emit ValueSet(newValue,now);
    }
}