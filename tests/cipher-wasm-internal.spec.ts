// Runs the tests from src/cipher/sec256k1-go/ and src/cipher/sec256k1-go/secp256k1-go2/
// after compiled to wasm

declare var Go: any;

describe('Tnternal test ', () => {

  let warningShown = false;

  const tmp = console.warn;
  console.warn = (message, ...optionalParams) => {
    warningShown = true;
    tmp(message, optionalParams);
  };

  let originalTimeout;

  beforeEach(function() {
    originalTimeout = jasmine.DEFAULT_TIMEOUT_INTERVAL;
    jasmine.DEFAULT_TIMEOUT_INTERVAL = 60000;
  });

  afterEach(function() {
    jasmine.DEFAULT_TIMEOUT_INTERVAL = originalTimeout;
  });

  it('test from src/cipher/sec256k1-go/ should pass', done => {
    warningShown = false;
    fetch('base/vendor/github.com/skycoin/skycoin/src/cipher/secp256k1-go/test.wasm').then(response => {
      response.arrayBuffer().then(ab => {
        const go = new Go();
        window['WebAssembly'].instantiate(ab, go.importObject).then(result => {
          go.run(result.instance).then(result => {
            if (warningShown == false) {
              done();
            } else {
              fail('Test failed.');
            }
          }, err => {
            fail('Test failed.');
          });
        });
      });
    });
  });

  it('test from src/cipher/sec256k1-go/secp256k1-go2/ should pass', done => {
    warningShown = false;
    fetch('base/vendor/github.com/skycoin/skycoin/src/cipher/secp256k1-go/secp256k1-go2/test.wasm').then(response => {
      response.arrayBuffer().then(ab => {
        const go = new Go();
        window['WebAssembly'].instantiate(ab, go.importObject).then(result => {
          go.run(result.instance).then(result => {
            if (warningShown == false) {
              done();
            } else {
              fail('Test failed.');
            }
          }, err => {
            fail('Test failed.');
          });
        });
      });
    });
  });
});
