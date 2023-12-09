import { readFile, p } from '../common.mjs';

const numbers = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}

let sum = 0;

const lines = await readFile('input.txt');

for (const line of lines) {
    p(line);

    let first;
    let second;

    for (const c of line) {
        if (Number.isInteger(parseInt(c))) {
            if (!first) {
                first = c;
            }
            second = c;
        }
    }
    sum = sum + parseInt(first + second);
    p(sum, first, second)
    first = undefined;
    second = undefined;
}
