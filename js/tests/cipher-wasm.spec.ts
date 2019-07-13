import { readJSON } from 'karma-read-json';

import { Address, testCases, convertAsciiToHexa } from './utils'

declare var Go: any;

describe('CipherProvider Lib', () => {
  const fixturesPath = 'tests/test-fixtures/';
  const addressesFileName = 'many-addresses.golden';
  const inputHashesFileName = 'input-hashes.golden';

  const seedSignaturesFiles = [
    'seed-0000.golden', 'seed-0001.golden', 'seed-0002.golden',
    'seed-0003.golden', 'seed-0004.golden', 'seed-0005.golden',
    'seed-0006.golden', 'seed-0007.golden', 'seed-0008.golden',
    'seed-0009.golden', 'seed-0010.golden'
  ];

  const testSettings = { addressCount: 1000, seedFilesCount: 11 };

  describe('initialize', () => {
    it('should initialize', done => {
      const go = new Go();

      fetch('base/skycoin-lite.wasm').then(response => {
        response.arrayBuffer().then(ab => {
          const go = new Go();
          window['WebAssembly'].instantiate(ab, go.importObject).then(result => {
            go.run(result.instance);

            done();
          });
        });
      });
    });
  });

  describe('generate address', () => {
    const addressFixtureFile = readJSON(fixturesPath + addressesFileName);
    const expectedAddresses = addressFixtureFile.keys.slice(0, testSettings.addressCount);
    let seed = convertAsciiToHexa(atob(addressFixtureFile.seed));
    let generatedAddress;

    testCases(expectedAddresses, (address: any) => {
      it('should generate many address correctly', done => {
        generatedAddress = generateAddress(seed);
        seed = generatedAddress.next_seed;

        const convertedAddress = {
          address: generatedAddress.address,
          public: generatedAddress.public_key,
          secret: generatedAddress.secret_key
        };

        expect(convertedAddress).toEqual(address);
        done();
      });

      it('should pass the verification', done => {
        verifyAddress(generatedAddress);
        done();
      });
    });
  });

  describe('seed signatures', () => {
    const inputHashes = readJSON(fixturesPath + inputHashesFileName).hashes;

    testCases(seedSignaturesFiles.slice(0, testSettings.seedFilesCount), (fileName: string) => {
      describe(`should pass the verification for ${fileName}`, () => {
        let seedKeys;
        let actualAddresses;
        let testData: { signature: string, public_key: string, hash: string, secret_key: string, address: string }[] = [];

        beforeAll(() => {
          const signaturesFixtureFile = readJSON(fixturesPath + fileName);
          const seed = convertAsciiToHexa(atob(signaturesFixtureFile.seed));
          seedKeys = signaturesFixtureFile.keys;

          actualAddresses = generateAddresses(seed, seedKeys);
          testData = getSeedTestData(inputHashes, seedKeys, actualAddresses);
        });

        it('should check number of signatures and hashes', done => {
          const result = seedKeys.some(key => key.signatures.length !== inputHashes.length);

          expect(result).toEqual(false);
          done();
        });

        it('should generate many address correctly', done => {
          actualAddresses.forEach((address, index) => {
            expect(address.address).toEqual(seedKeys[index].address);
            expect(address.public_key).toEqual(seedKeys[index].public);
            expect(address.secret_key).toEqual(seedKeys[index].secret);
          });

          done();
        });

        it('address should pass the verification', done => {
          verifyAddresses(actualAddresses);
          done();
        });

        it(`should verify signature correctly`, done => {
          testData.forEach(data => {
            const result = window['SkycoinCipherExtras'].verifyPubKeySignedHash(data.public_key, data.signature, data.hash);
            expect(result).toBeNull();
            done();
          });
        });

        it(`should check signature correctly`, done => {
          testData.forEach(data => {
            const result = window['SkycoinCipherExtras'].verifyAddressSignedHash(data.address, data.signature, data.hash);
            expect(result).toBeNull();
            done();
          });
        });

        it(`should verify signed hash correctly`, done => {
          testData.forEach(data => {
            const result = window['SkycoinCipherExtras'].verifySignatureRecoverPubKey(data.signature, data.hash);
            expect(result).toBeNull();
            done();
          });
        });

        it(`should generate public key correctly`, done => {
          testData.forEach(data => {
            const pubKey = window['SkycoinCipherExtras'].pubKeyFromSig(data.signature, data.hash);
            expect(pubKey).toBeTruthy();
            expect(pubKey === data.public_key).toBeTruthy();
            done();
          });
        });

        it(`sign hash should be created`, done => {
          testData.forEach(data => {
            const sig = window['SkycoinCipherExtras'].signHash(data.hash, data.secret_key);
            expect(sig).toBeTruthy();
            done();
          });
        });
      });
    });
  });
});

function getSeedTestData(inputHashes, seedKeys, actualAddresses) {
  const data = [];

  for (let seedIndex = 0; seedIndex < seedKeys.length; seedIndex++) {
    for (let hashIndex = 0; hashIndex < inputHashes.length; hashIndex++) {
      data.push({
        signature: seedKeys[seedIndex].signatures[hashIndex],
        public_key: actualAddresses[seedIndex].public_key,
        secret_key: actualAddresses[seedIndex].secret_key,
        address: actualAddresses[seedIndex].address,
        hash: inputHashes[hashIndex]
      });
    }
  }

  return data;
}

function generateAddresses(seed: string, keys: any[]): Address[] {
  return keys.map(() => {
    const generatedAddress = generateAddress(seed);
    seed = generatedAddress.next_seed;

    return generatedAddress;
  });
}

function generateAddress(seed: string): Address {
  const address = window['SkycoinCipher'].generateAddress(seed);
  return {
    address: address.address,
    public_key: address.public,
    secret_key: address.secret,
    next_seed: address.nextSeed
  };
}

function verifyAddress(address) {
  const addressFromPubKey = window['SkycoinCipherExtras'].addressFromPubKey(address.public_key);
  const addressFromSecKey = window['SkycoinCipherExtras'].addressFromSecKey(address.secret_key);

  expect(addressFromPubKey && addressFromSecKey && addressFromPubKey === addressFromSecKey).toBe(true);

  expect(window['SkycoinCipherExtras'].verifySeckey(address.secret_key)).toBe(null);
  expect(window['SkycoinCipherExtras'].verifyPubkey(address.public_key)).toBe(null);
}

function verifyAddresses(addresses) {
  addresses.forEach(address => verifyAddress(address));
}
