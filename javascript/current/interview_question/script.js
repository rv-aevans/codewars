import fetch from "node-fetch";

async function getInput() {
    return (await fetch("https://dev.mtechapis.com/interview")).json()
    .then(resp => {
        postSolution(JSON.stringify(resp), JSON.stringify(resp).match(/[bcdfghjklmnpqrstvwxyzs]/gi).reverse().join().replaceAll(",", ""))
    })
}

async function postSolution(s, r) {
    await fetch("https://dev.mtechapis.com/interview?input=" + new URLSearchParams(s), {
        method: "POST",
        body: JSON.stringify({
            output: r
        }),
    })
    .then((response) => {
        console.log(response.status)
    })
}

getInput()