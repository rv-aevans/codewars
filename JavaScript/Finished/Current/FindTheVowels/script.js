
function vowelIndices(word){
    const v = ['a','e','i','o','u','y']
    const res = []
    for (i=1;i<=word.length;i++) {
        v.includes(word[i-1].toLowerCase()) ? res.push(i) : null
    }
    console.log(res)
    return res
}

vowelIndices("testE")