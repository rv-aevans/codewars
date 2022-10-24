function rotate(matrix, direction) {
    var newMatrix = []
    if (direction.includes('counter')) {
        for (i = 1; i <= matrix.length; i++) {
            let newArr = []
            for (j = 0; j < matrix.length; j++) {
                newArr.push(matrix[j][matrix[0].length-i])
            }
            newMatrix.push(newArr)
        }
        return newMatrix
    } 
    for (i = 0; i < matrix.length; i++) {
        let newArr = []
        for (j = 1; j <= matrix.length; j++) {
            newArr.push(matrix[matrix.length-j][i])
        }
        newMatrix.push(newArr)
    }
    return newMatrix
}

let matrix = [
    [1,2,3],
    [4,5,6],
    [7,8,9]
];

console.log(rotate(matrix, 'counter-clockwise'))