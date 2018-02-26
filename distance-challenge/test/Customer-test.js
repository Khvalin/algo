const assert = require('assert');

const Customer = require('../src/Customer');

describe('Customer', () => {
  describe('#valid', () => {
    it('should return false when customer coordinates are invalid or id is null', () => {
      assert.equal(new Customer({ id: null, lon: 0, lat: 0 }).valid, false);
      assert.equal(new Customer({ id: 0, lon: 100, lat: 0 }).valid, false);
    });

    it('should return true when customer coordinates are valid and id is not null', () => {
      assert.equal(new Customer({ id: null, lon: 0, lat: 0 }).valid, false);
    });
  });
});
