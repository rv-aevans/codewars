const numbers = [1, 2, 34, 2, 1, 5, 3, 5, 7, 234, 2, 1]

function threeAmigos(numbers) {
    let res = []
    for (let i = 0; i < numbers.length-2; i++) {
        if (numbers[i] === numbers[i+2]) {
            if (Math.abs(numbers[i] - numbers[i+1]) < 3) {
                res.push(numbers[i], numbers[i+1], numbers[i+2])
                return res
            }
            if (numbers[i] - numbers[i+1] - numbers)
        }
    }
    return res
}

console.log(threeAmigos(numbers))