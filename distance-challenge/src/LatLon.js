/**
 * Represents coordinate latitude/longitude pair
 */
class LatLon {
  /**
   * creates a new LatLon instance
   * @param {float} lat -- latitude 
   * @param {float} lon -- longitude 
   */
  constructor(lat = 0, lon = 0) {
    this.lat = lat;
    this.lon = lon;
  }

  /**
   * 
   * @returns {boolean}
   */
  get valid() {
    const { lat, lon } = this;
    return (
      Number.isFinite(lat) &&
      Number.isFinite(lon) && // this also
      lat >= -90 &&
      lat <= 90 &&
      lon >= -180 &&
      lon <= 180
    );
  }

  toString() {
    return `${this.lat}, ${this.lon}`;
  }
}

module.exports = LatLon;
