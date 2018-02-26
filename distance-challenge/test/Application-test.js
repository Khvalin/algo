const assert = require('assert');

const Application = require('../src/Application');

const InputReader = require('../src/InputReader');
const CustomerFactory = require('../src/CustomerFactory');
const Geometry = require('../src/Geometry');
const LatLon = require('../src/LatLon');

function setUp() {
  const fake_lat = 0;
  const fake_lon = 0;

  const responseMock = [
    { latitude: fake_lat, user_id: null, name: 'Null UserId', longitude: fake_lon },
    { latitude: fake_lat, user_id: 42, longitude: fake_lon },
    { latitude: fake_lat, user_id: 41, longitude: fake_lon }
  ];

  class FakeInputReader {
    readToArray(stream) {
      return new Promise((resolve) => resolve(responseMock));
    }
  }

  const officeGps = new LatLon(fake_lat, fake_lon);
  const maxDistance = 100;

  const application = new Application({ inputReader: new FakeInputReader(), CustomerFactory, Geometry });

  return application.getCustomersWithinDistance({ officeGps, maxDistance });
}

describe('Application', () => {
  describe('#getCustomersWithinDistance', () => {
    it('should ignore customers with invalid id', (done) => {
      setUp().then((list) => {
        assert.equal(list.findIndex((c) => c.name === 'Null UserId'), -1);
        done();
      });
    });

    it('should allow for empty customer names', (done) => {
      setUp().then((list) => {
        assert.notEqual(list.findIndex((c) => c.id === 42), -1);
        done();
      });
    });

    it('should sort customers by id', (done) => {
      setUp().then((list) => {
        assert.equal(list.findIndex((c) => c.id === 42), 1);
        assert.equal(list.findIndex((c) => c.id === 41), 0);
        done();
      });
    });
  });
});
