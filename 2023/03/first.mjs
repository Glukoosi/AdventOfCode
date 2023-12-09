import { readFile, p } from '../common.mjs';

const lines = await readFile('input.txt');

const isTherePart = (x, y, len, lines) => {

    for (let j = y - 1; j <= y + 1; j++){
        for (let i = x - len; i <= x + 1; i++){
            try {
                if (lines[j][i] !== '.' && isNaN(lines[j][i]) && lines[j][i]){
                    return true
                }
            } catch {
                continue
            }
        }
    }

    return false
}

let answer = 0
for (let y = 0; y < lines.length; y++) {
    const line = lines[y]

    let number = ''
    for (let x = 0; x < line.length; x++) {
        const ch = line[x]
        if (!isNaN(ch)) {
            number = number + ch
        }
        if ((!isNaN(line[x - 1]) && isNaN(ch)) || (x === line.length - 1 && !isNaN(ch))) {
            if (isTherePart(x - 1, y, number.length, lines)) {
                answer = answer + parseInt(number)
            }
            number = ''
        }

    }
}
p(answer)
