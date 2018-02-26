const args = require('yargs').argv;

const InputReader = require('./src/InputReader');
const CustomerFactory = require('./src/CustomerFactory');
const Geometry = require('./src/Geometry');
const LatLon = require('./src/LatLon');

const Application = require('./src/Application');

const application = new Application({ inputReader: new InputReader(), CustomerFactory, Geometry });

const officeGps = new LatLon(args.lat || 53.339428, args.lon || -6.257664);
const maxDistance = args.maxDistance || 100;

application
  .getCustomersWithinDistance({
    officeGps,
    maxDistance
  })
  .then((list = []) => {
    console.log(`Customers within ${maxDistance}km distance from ${officeGps}:`);
    console.log(list.map((c) => `${c.id} ${c.name}`).join('\n'));
  })
  .catch((e) => console.error(e));
