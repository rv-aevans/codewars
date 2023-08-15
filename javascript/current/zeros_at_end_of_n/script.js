function countZeros(n) {
    if (n < 3) {
        return 0
    }

    let total = BigInt(1)

    for (let i = n; i > 0; i -= 2) {
        total *= BigInt(i)
    }

    let count = 0
    let s = total.toString()
    for (let i = s.length-1; i > 0; i--) {
        if (s[i] === "0") {
            count++
        } else {
            return count
        }
    }
}


console.log(countZeros(290))