const assert = require('assert');
const streamify = require('streamify-string');

const InputReader = require('../src/InputReader');

describe('InputReader', () => {
  describe('#readToArray', () => {
    it('should tolerate invalid JSON lines', (done) => {
      const readStream = streamify('errorline\n{"user_id":42}');

      new InputReader().readToArray(readStream).then((parsedInput) => {
        assert.equal(parsedInput[0], null);

        done();
      });
    });

    it('should parse users data correctly', (done) => {
      const readStream = streamify('errorline\n{"user_id":42}');

      new InputReader().readToArray(readStream).then((parsedInput) => {
        assert.equal(parsedInput[1].user_id, 42);

        done();
      });
    });
  });
});
