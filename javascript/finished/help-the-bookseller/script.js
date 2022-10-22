function stockList(listOfArt, listOfCat){
    if (listOfArt.length == 0 || listOfCat.length == 0) {
        return ""
    }
    stockMap = new Map;
    listOfArt.forEach(b => {
        stockMap.get(b[0]) == undefined ? stockMap.set(b[0], parseInt(b.split(" ")[1], 10)) : stockMap.set(b[0], (stockMap.get(b[0]) + parseInt(b.split(" ")[1], 10)))
    });
    var res = ""
    listOfCat.forEach(k => {
        stockMap.get(k) == undefined ? res+= `(${k} : 0) - ` : res += `(${k} : ${stockMap.get(k)}) - `
    })
    return res.substring(0, res.length - 3)
}

console.log(stockList([], ["A", "B", "X"]))

// console.log(parseInt('ABAR 200'.split(" ")[1]))