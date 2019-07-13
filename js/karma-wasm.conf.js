// Karma configuration file, see link for more information
// https://karma-runner.github.io/0.13/config/configuration-file.html

module.exports = function (config) {

  config.set({
    basePath: '',
    frameworks: ['jasmine', 'karma-typescript'],
    plugins: [
      require('karma-jasmine'),
      require('karma-chrome-launcher'),
      require('karma-jasmine-html-reporter'),
      require('karma-read-json'),
      require('karma-typescript')
    ],
    files: [
      'tests/cipher-wasm.spec.ts',
      { pattern: 'tests/test-fixtures/*.golden', included: false },
      { pattern: 'skycoin-lite.wasm', included: false },
      { pattern: 'test1.wasm', included: false },
      { pattern: 'test2.wasm', included: false },
      { pattern: 'tests/utils.ts', included: true },
      { pattern: 'tests/wasm_exec.js', included: true },
    ],
    preprocessors: {
      "**/*.ts": "karma-typescript"
    },
    client: {
      clearContext: false // leave Jasmine Spec Runner output visible in browser
    },
    reporters: ['progress', 'kjhtml', 'karma-typescript'],
    port: 9876,
    colors: true,
    logLevel: config.LOG_INFO,
    autoWatch: true,
    browsers: ['ChromeHeadless', 'Chrome'],
    singleRun: false
  });
};
