

function scramble(str1, str2) {
	if (str2.length > str1.length) {
		return false;
	}
	for (i = 0; i < str2.length; i++) {
		if (str1.indexOf(str2[i]) === -1) {
			return false;
		}
		arr = str1.split('');
		arr.splice(arr.indexOf(str2[i]), 1);
		str1 = arr.join('');
	}
	return true;
}

console.log(scramble("scriptjavx", "javascript"));