/** 
 * entry point for the solution -- glues modules together and performs searching
*/
class Application {
  constructor(modules) {
    this._modules = modules;
  }

  getCustomersWithinDistance({ officeGps, maxDistance, inputStream }) {
    inputStream = inputStream || process.stdin;

    const { inputReader, CustomerFactory, Geometry } = this._modules;

    return inputReader.readToArray(inputStream).then((customers = []) => {
      return (
        customers
          // filter out null values and convert the rest to "Customer" objects:
          .map((c) => c && CustomerFactory.fromObject(c))
          // filter out invalid Customer objects and Customers outside the required distance:
          .filter((c) => c && c.valid && Geometry.haversineDistance(officeGps, c.gps) <= maxDistance)
          // sort the rest by id ASC:
          .sort((c1, c2) => (c1.id < c2.id ? -1 : c1.id > c2.id ? 1 : 0))
      );
    });
  }
}

module.exports = Application;
