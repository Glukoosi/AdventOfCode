import { readFile, p } from '../common.mjs';

const possibleMax = {
    "red": 12,
    "green": 13,
    "blue": 14
}

const lines = await readFile('input.txt');

let answerSum = 0;

let index = 1;
for (const line of lines) {
    let gameGoodFlag = 1;
    const games = line.split(': ')[1].split('; ');

    for (const game of games) {
        let gameParsed = {}
        game.split(', ').map(result => {
            let number, color
            [number, color] = result.split(' ')
            gameParsed[color] = parseInt(number)
        })

        for (const key in gameParsed) {
            if (possibleMax[key] < gameParsed[key]) {
                gameGoodFlag = 0
            }
        }
    }

    if (gameGoodFlag === 1) {

        answerSum = answerSum + index
    }
    index++
}

p(answerSum)
