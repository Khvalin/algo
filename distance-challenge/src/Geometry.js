/**
 * contains geometrical functions
 */
class Geometry {
  static toRad(v) {
    return v * Math.PI / 180;
  }

  /**
   * calculates haversine distance between two points
   * @param {LatLon} gps1
   * @param {LatLon} gps2
   */
  static haversineDistance(gps1, gps2) {
    const R = 6371; // Earth readius in km
    const dLat = Geometry.toRad(gps2.lat - gps1.lat);
    const dLon = Geometry.toRad(gps2.lon - gps1.lon);

    const a =
      Math.sin(dLat / 2) * Math.sin(dLat / 2) +
      Math.cos(Geometry.toRad(gps1.lat)) * Math.cos(Geometry.toRad(gps2.lat)) * Math.sin(dLon / 2) * Math.sin(dLon / 2);

    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));

    return R * c;
  }
}

module.exports = Geometry;
