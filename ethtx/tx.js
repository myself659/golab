var Web3 = require('web3');
var web3 = new Web3(new Web3.providers.HttpProvider('https://ropsten.infura.io/'));
var util = require('ethereumjs-util');
var tx = require('ethereumjs-tx');
var publicKey = util.bufferToHex(util.privateToPublic(privateKey));
console.log(publicKey)