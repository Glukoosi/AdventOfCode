import { open } from 'node:fs/promises';

export const p = (...print) => console.log(...print);

export const readFile = async (fileName) => {
    const file = await open(fileName);

    const lines = []
    for await (const line of file.readLines()) {
        lines.push(line);
    }
    return lines
}
