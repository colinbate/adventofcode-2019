const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

let total = 0;
rl.on('line', input => {
  let fuel = 0;
  let mass = input * 1;
  while (mass > 0) {
    mass = Math.max(Math.floor(mass / 3) - 2, 0);
    fuel += mass;
  }
  total += fuel;
});

rl.on('close', () => {
  console.log(`\nTotal: ${total}`);
});