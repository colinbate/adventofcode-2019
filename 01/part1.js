const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

let total = 0;
rl.on('line', mass => {
  const fuel = Math.floor(mass / 3) - 2;
  total += fuel;
});

rl.on('close', () => {
  console.log(`\nTotal: ${total}`);
});