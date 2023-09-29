var findMissing = function (list) {  
    return list[Array.from(list.keys()).find(i => list[i+1]+Math.min(list[0]-list[1], list[1]-list[2]) != list[i])+1]+Math.min(list[0]-list[1], list[1]-list[2])
}
console.log(findMissing([29, 34, 44]))
