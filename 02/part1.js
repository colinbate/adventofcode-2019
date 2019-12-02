const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin
});

let codes;
rl.on('line', input => {
  codes = input.split(',').map(x => parseInt(x, 10));
  console.log(`\nRead ${codes.length} codes`);
});

const ADD = 1;
const MULT = 2;
const HALT = 99;
function run() {
  console.log('Running...');
  codes[1] = 12;
  codes[2] = 2;
  let ptr = 0;
  while (codes[ptr] !== HALT && ptr < 3000) {
    const op = codes[ptr];
    const arg1 = codes[codes[ptr+1]] || 0;
    const arg2 = codes[codes[ptr+2]] || 0;
    const outptr = codes[ptr+3];
    const out = op === ADD ? arg1 + arg2 : arg1 * arg2;
    // console.log(ptr, op, arg1, arg2, outptr, out);
    codes[outptr] = out;
    ptr += 4;
  }
  console.log(`\nOut: ${codes[0]}`);
}

rl.on('close', run);