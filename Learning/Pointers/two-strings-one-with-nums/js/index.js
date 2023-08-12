const isSameStr = (str1, str2) => {
  for (let i = 0, j = 0; j < str2.length; i++, j++) {
    if (isNaN(parseInt(str2[j]))) {
      if (str1[i] !== str2[j]) return false;
    } else {
      i += +str2[j] - 1;
    }
  }

  return true;
};

let str1 = "hi this is dog";
let str2 = "h4is is 2g";

console.log(isSameStr(str1, str2));

str1 = "hi this is dog";
str2 = "h4is i2dog";

console.log(isSameStr(str1, str2));
