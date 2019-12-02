const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin
});

const GOAL = 19690720;

let memory;
rl.on('line', input => {
  memory = input.split(',').map(x => parseInt(x, 10));
  console.log(`\nRead ${memory.length} codes`);
});

const ADD = 1;
const MULT = 2;
const HALT = 99;

function compute(noun, verb, codes) {
  // console.log(`Running for noun ${noun} and verb ${verb}`);
  codes[1] = noun;
  codes[2] = verb;
  let ptr = 0;
  while (codes[ptr] !== HALT && ptr < codes.length) {
    const op = codes[ptr];
    const arg1 = codes[codes[ptr+1]] || 0;
    const arg2 = codes[codes[ptr+2]] || 0;
    const outptr = codes[ptr+3];
    const out = op === ADD ? arg1 + arg2 : arg1 * arg2;
    // console.log(ptr, op, arg1, arg2, outptr, out);
    codes[outptr] = out;
    ptr += 4;
  }
  return codes[0];
}

function run() {
  for (let n = 0; n < 100; n += 1) {
    for (let v = 0; v < 100; v += 1) {
      const out = compute(n, v, memory.slice());
      if (out === GOAL) {
        console.log(`\nOut: ${100 * n + v}`);
        return;
      }
    }
  }
  console.log('Not found... :(');
}

rl.on('close', run);