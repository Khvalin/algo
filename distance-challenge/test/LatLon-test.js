const assert = require('assert');

const LatLon = require('../src/LatLon');

describe('LatLon', () => {
  describe('#valid', () => {
    it('should return true when either lat or lon value is 0', () => {
      assert.equal(new LatLon(0, 0).valid, true);
    });

    it('should return false when the lat or lon value is out of range', () => {
      assert.equal(new LatLon(-100, 0).valid, false);
      assert.equal(new LatLon(100, 0).valid, false);
      assert.equal(new LatLon(0, 4242).valid, false);
    });
  });
});
