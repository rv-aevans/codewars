function longestConsec(stsarr, k) {
    if (stsarr.length === 0 || k > stsarr.length || k <= 0) {
        return ""
    }

    var res = "";
    var max = 0;


    for (var i = 0; i <= stsarr.length - k; i++) {
        var newMax = 0
        var newRes = ""
        for (var j = i; j < k + i; j++) {
            newMax += stsarr[j].length
            newRes += stsarr[j]
            if (newMax > max) {
                max = newMax
                res = newRes
            }
        }
    }
    return res
}

console.log(longestConsec(["itvayloxrp","wkppqsztdkmvcuwvereiupccauycnjutlv","vweqilsfytihvrzlaodfixoyxvyuyvgpck"], 2))