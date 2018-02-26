/** 
 * encapsulates methods to work with input stream
*/
class InputReader {
  /**
   * reads stream data to a string
   * @param {stream} stream 
   */
  readToString(stream) {
    return new Promise((resolve, reject) => {
      let streamData = '';

      stream.on('error', reject);

      stream.on('data', (chunk) => {
        if (chunk !== null) {
          streamData += chunk;
        }
      });

      stream.on('end', () => {
        resolve(streamData);
      });
    });
  }

  /**
   * reads stream data and parses each line into JS object
   * @param {stream} stream 
   */
  readToArray(stream) {
    return this.readToString(stream).then((input = '') => {
      return input.split('\n').map((line) => {
        try {
          return JSON.parse(line);
        } catch (e) {
          return null;
        }
      });
    });
  }
}

module.exports = InputReader;
