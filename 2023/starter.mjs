import { readFile, p } from '../common.mjs';

const lines = await readFile('example.txt');

for (const line of lines) {
    p(line);
}
