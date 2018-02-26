const LatLon = require('./LatLon');

/**
 * represents Customer
 */
class Customer {
  get valid() {
    return !!(this.id && this.gps.valid);
  }

  /**
  * @param customerInfo Information about the customer
  * @param customerInfo.id Customer id
  * @param customerInfo.lat latitude value of customer position
  * @param customerInfo.lon  longitude value of customer position
  * @param customerInfo.name The name of the customer
  */
  constructor(customerInfo) {
    this.id = customerInfo.id || null;
    this.gps = new LatLon(parseFloat(customerInfo.lat), parseFloat(customerInfo.lon));
    this.name = customerInfo.name || null;
  }
}

module.exports = Customer;
