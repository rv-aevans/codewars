function stringExpansion(s) {
    let res = ''
    let curMult = 1
    s = s.replace(/\s+/g, '')
    for (i = 0; i < s.length-1; i++) {
        if (isNum(s[i])) {
            curMult = s[i]
            if (!isNum(s[i+1])) {
                for (j = 0; j < curMult; j++) {
                    res += s[i+1]
                }
                continue
            }
        } 
        if (!isNum(s[i]) && i == 0) {
            res += s[i]
        }
        if (!isNum(s[i+1])) {
            for (j = 0; j < curMult; j++) {
                res += s[i+1]
            }
        } 
    }
    return res
}

function isNum(n) {
    if ((/^\d$/).test(n)) {
        return true
    } else {
        return false
    }
}

console.log(stringExpansion('3D2a5d2f'))

// DESCRIPTION:
// Given a string that includes alphanumeric characters ("3a4B2d") return the expansion of that string: The numeric values represent the occurrence of each letter preceding that numeric value. There should be no numeric characters in the final string.

// Notes
// The first occurrence of a numeric value should be the number of times each character behind it is repeated, until the next numeric value appears
// If there are multiple consecutive numeric characters, only the last one should be used (ignore the previous ones)
// Empty strings should return an empty string.
// Your code should be able to work for both lower and capital case letters.

// function stringExpansion(s) {
//     return s.replace(/\d\D*/g,m=>m.slice(1).replace(/./g,n=>n.repeat(m[0])))
// }