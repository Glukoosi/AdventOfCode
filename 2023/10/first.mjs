import { readFile, p } from '../common.mjs';


const lines = await readFile('input.txt');

const pipe = {
    '|': ['north', 'south'],
    '-': ['east', 'west'],
    'L': ['north', 'east'],
    'J': ['north', 'west'],
    '7': ['west', 'south'],
    'F': ['east', 'south'],
    '.': [],
    'S': ['north', 'south', 'east', 'west'],
}

const scanTiles = (y, x, lines) => {
    const originalTile = pipe[lines[y][x]]
    let possibleCord = []
    for (const dir of originalTile) {
        switch (dir) {
            case 'north':
                try {
                    if( pipe[lines[y - 1][x]].includes('south')){
                        possibleCord.push([y - 1, x])
                    }
                    continue
                } catch {
                    continue
                }
            case 'south':
                try {
                    if( pipe[lines[y + 1][x]].includes('north')){
                        possibleCord.push([y + 1, x])
                    }
                    continue
                } catch {
                    continue
                }
            case 'west':
                try {
                    if( pipe[lines[y][x - 1]].includes('east')){
                        possibleCord.push([y, x - 1])
                    }
                    continue
                } catch {
                    continue
                }
            case 'east':
                try {
                    if( pipe[lines[y][x + 1]].includes('west')){
                        possibleCord.push([y, x + 1])
                    }
                    continue
                } catch {
                    continue
                }
        }
    }
    return possibleCord

}

const pathLoop = (startingPoint) => {
    let possibleCords = scanTiles(startingPoint[0], startingPoint[1], lines)

    possibleCords = [possibleCords[0]]
    let lastTiles = [startingPoint.toString()]
    let i = 0
    while (true) {
        for (const possibleCord of possibleCords) {
            if (!lastTiles.includes(possibleCord.toString())) {
                lastTiles.push(possibleCord.toString())
                possibleCords = scanTiles(possibleCord[0], possibleCord[1], lines)
                break
            }
            if (possibleCord === possibleCords[possibleCord.length - 1]) {
                return i
            }
        }
        i++
    }
}

for (let y = 0; y < lines.length; y++) {
    const row  = lines[y]
    for (let x = 0; x < row.length; x++) {
        const tile = row[x]
        if (tile === 'S') {
            const startingPoint = [y, x]
            const count = pathLoop(startingPoint);
            const answer = Math.ceil(count / 2)
            p(answer)
        }
    }
}
