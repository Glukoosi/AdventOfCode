import { readFile, p } from '../common.mjs';

const lines = await readFile('input.txt');

let answerSum = 0;

let index = 1;
for (const line of lines) {
    const games = line.split(': ')[1].split('; ');

    const possibleMax = {
        "red": 0,
        "green": 0,
        "blue": 0
    }

    for (const game of games) {
        let gameParsed = {}
        game.split(', ').map(result => {
            let number, color
            [number, color] = result.split(' ')
            gameParsed[color] = parseInt(number)
        })

        for (const key in gameParsed) {
            if (possibleMax[key] < gameParsed[key]) {
                possibleMax[key] = gameParsed[key]
            }
        }

    }
    p(possibleMax)
    answerSum = answerSum + possibleMax.red * possibleMax.green * possibleMax.blue
    index++
}

p(answerSum)
