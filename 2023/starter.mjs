import { readFile, p } from '../common.mjs';

const lines = await readFile('input.txt');

for (const line of lines) {
    p(line);
}
