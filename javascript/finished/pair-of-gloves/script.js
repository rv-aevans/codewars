function numberOfPairs(gloves){
    var res = new Map;
    for (i=0;i<gloves.length;i++) {
        res.get(gloves[i]) == undefined ? res.set(gloves[i], 1) : res.set(gloves[i], res.get(gloves[i])+1)
    }
    return [...res.values()].reduce((a,b) => a + Math.floor(b/2), 0)
}

console.log(numberOfPairs(['red','red','blue','blue','blue','green']))
// Prints 2
