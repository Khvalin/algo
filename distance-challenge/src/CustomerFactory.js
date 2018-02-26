const LatLon = require('./LatLon');
const Customer = require('./Customer');

class CustomerFactory {
  static fromObject(customerData = {}) {
    return new Customer({
      id: customerData.user_id,
      lat: customerData.latitude,
      lon: customerData.longitude,
      name: customerData.name
    });
  }
}

module.exports = CustomerFactory;
