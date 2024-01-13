use std::io::{self, Read, BufReader};
use std::io::{ Write, BufWriter};

fn main() -> io::Result<()> {
    // Get the standard output (stdout) handle
    let stdout = io::stdout();

    // Create a buffered writer for stdout
    let writer = stdout.lock();
    let mut buffered_writer = BufWriter::new(writer);

    // Get the standard input (stdin) handle
    let stdin = io::stdin();

    // Create a buffered reader using stdin
    let reader = stdin.lock();
    let buffered_reader = BufReader::new(reader);

    let mut num = [0;3];
    let mut line_index = 0;
    // Read bytes from the buffered reader
    for byte_result in buffered_reader.bytes() {
        let byte = byte_result?;

        if byte == b'\n' {
            num = 0;
            line_index += 1;
        }

        if byte >= b'0' && byte <= b'9' {
            num = 10 * num;
            num = num + (byte - b'0');
        }


        buffered_writer.write(&[byte])?;
    }

    // Ensure all buffered data is flushed to stdout
    buffered_writer.flush()?;

    Ok(())
}
