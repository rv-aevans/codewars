function arrayDiff(a, b) {
    b.forEach(element => { a = a.filter(item => item !== element)})
    return a
}

console.log(arrayDiff([1,1,1,1,2],[1])) 