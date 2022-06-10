function csvColumns(csv, indices) {
    var matrix = [[]]
    indices = [...new Set(indices)]
    i = 0
	csv.split("\n").forEach(s => {
        s.split(",").forEach(v => {
            matrix[i] ? matrix[i].push(v) : matrix.push([v])
        })
        i++
    })
    console.log(matrix)
    var res = ""
    matrix.forEach(v => {
        indices.sort().forEach(i => {
            i >= v.length ? null : res+=`${v[i]},`
        })
        res = res.slice(0, -1) + "\n"
    })
    return res.slice(0, -1)
}

console.log(csvColumns("m,EpSj,kMKFm,yI,FlTjD,Jb,khZ,9rsrV" , [5,0,10,4,9,10,9,8]))