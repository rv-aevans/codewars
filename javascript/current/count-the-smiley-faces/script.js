countSmileys=a=>(a.filter(e=>e.match(/[;:][-~]?[)D]/)).length)

console.log(countSmileys([':-D', ':(', ';~)', ':D']))