import { readFile, p } from '../common.mjs';

const lines = await readFile('input.txt');

let answer = 0

const datasetRecursive = (dataset) => {
    const latest = dataset[dataset.length - 1]
    const nextDataset = []
    for ( let i = 0; i < latest.length - 1; i++ ) {
        nextDataset.push(latest[i+1] - latest[i])
    }
    dataset.push(nextDataset)

    if (nextDataset.every(v => v === 0)) {
        return dataset
    }
    return datasetRecursive(dataset)
}

for (const line of lines) {
    const dataset = line.split(' ').map(n => parseInt(n));

    const sequence = datasetRecursive([dataset])

    let sum = 0
    for (let i = sequence.length - 1; i >= 0; i--) {
        const last = sequence[i][sequence[i].length - 1]
        sum = sum + last
    }
    answer = answer + sum
}
p(answer)
