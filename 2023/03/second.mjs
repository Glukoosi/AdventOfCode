import { readFile, p } from '../common.mjs';

const lines = await readFile('input.txt');

const isThereNumbers = (x, y, lines) => {
    let numbersC = []
    let flag = 0
    for (let j = y - 1; j <= y + 1; j++){
        for (let i = x - 1; i <= x + 1; i++){
            try {
                if (!isNaN(lines[j][i]) && lines[j][i]){
                    if (flag === 0) {
                        numbersC.push([j, i])
                        flag = 1
                    }
                    if (i === x + 1) {
                        flag = 0
                    }
                    continue
                }
                flag = 0
            } catch {
                continue
            }
        }
    }

    if (numbersC.length === 2) {
        return numbersC
    }
    return false
}

let numberCord = {}
let foundCord = []
for (let y = 0; y < lines.length; y++) {
    const line = lines[y]

    let number = ''
    for (let x = 0; x < line.length; x++) {
        const ch = line[x]
        if (!isNaN(ch)) {
            number = number + ch
        }
        if ((!isNaN(line[x - 1]) && isNaN(ch)) || (x === line.length - 1 && !isNaN(ch))) {
            let n = parseInt(number)
            for ( let c = x - 1; c > x - number.length - 1; c-- ) {
                numberCord[y.toString() + ' ' + c] = n
            }
            number = ''
        }
        if (ch === '*') {
            const numbersCoordinates = isThereNumbers(x, y, lines)
            if (numbersCoordinates) {
                foundCord.push(numbersCoordinates)
            }
        }
    }
}

let answer = 0
for (const cords of foundCord) {
    const result = (numberCord[cords[0][0].toString() + ' ' + cords[0][1]]) * 
        (numberCord[cords[1][0].toString() + ' ' + cords[1][1]])
    answer = answer + result
}
p(answer)
