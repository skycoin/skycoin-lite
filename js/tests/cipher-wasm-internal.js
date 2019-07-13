// Runs the tests from src/cipher/sec256k1-go/ and src/cipher/sec256k1-go/secp256k1-go2/.
// It needs the tests to be compiled, so it was created for being called by "make test-suite-ts-wasm"

// This is done to make wasm_exec think that it is not running inside Node.js, so that it does not ask
// for special parameters and works in a similar way as how it would do in an browser.
global.process.title = '';

// Required for wasm_exec to work correctly in Node.js.
const util = require('util');
TextEncoder = util.TextEncoder;
TextDecoder = util.TextDecoder

require('./wasm_exec');
const fs = require('fs');

// Required for wasm_exec to work correctly in Node.js.
performance = {
  now() {
    const [sec, nsec] = process.hrtime();
    return sec * 1000 + nsec / 1000000;
  },
};

// Required for wasm_exec to work correctly in Node.js.
const nodeCrypto = require("crypto");
crypto = {
  getRandomValues(b) {
    nodeCrypto.randomFillSync(b);
  },
};

// wasm_exec uses console.warn in case of error, so this code uses it to detect when a test fails.
const tmp = console.warn;
console.warn = (message, ...optionalParams) => {
  tmp(message, optionalParams);
  // Is a test fails, the process is closed with an error code.
  process.exit(1);
};

runTest1 = function() {
  const testFile = fs.readFileSync('../../vendor/github.com/skycoin/skycoin/src/cipher/secp256k1-go/test.wasm', null).buffer;
  const go = new global.Go();
  WebAssembly.instantiate(testFile, go.importObject).then(result => {
    go.run(result.instance).then(() => {
      runTest2();
    }, err => {
      console.log(err);
      process.exit(1);
    });
  });
}

runTest2 = function() {
  const testFile = fs.readFileSync('../../vendor/github.com/skycoin/skycoin/src/cipher/secp256k1-go/secp256k1-go2/test.wasm', null).buffer;
  const go = new global.Go();
  WebAssembly.instantiate(testFile, go.importObject).then(result => {
    go.run(result.instance).then(() => {

    }, err => {
      console.log(err);
      process.exit(1);
    });
  });
}

runTest1();
