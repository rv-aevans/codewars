function possibilities(str) {
    let res = []
    str.split('').forEach(e => {
        if (e == '?') {
            res.push(str.replace('?', '0'))
            res.push(str.replace('?', '1'))
        }
    });
    return res
};

console.log(possibilities('101?'))