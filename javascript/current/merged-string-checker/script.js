function isMerge(s, part1, part2) {
    return s.includes(part1) && s.includes(part2)
}

console.log(isMerge('xcyc', 'yc', 'xc'))