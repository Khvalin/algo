# distance-challenge

A simple nodejs program which that reads the full list of customers and outputs the names and user ids of matching customers (within 100km), sorted by User ID (ascending).

## Install the dependencies

```bash
  npm install
```

## Usage

```bash
  npm run default
```
this runs the program with default input parameters and is equalent to:

```bash
  node index --lat 53.339428 --lon -6.257664 --max-distance 100 < ./data/customers.txt
```
so the office coordinates, max customer distance, and customers data are input params 

```bash
  npm test
```
runs unit tests


## License

[MIT](http://vjpr.mit-license.org)
