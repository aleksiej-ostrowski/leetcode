function romanToInt(s: string): number {
  if (s.length == 0) {
    return 0;
  }

  let res = 0;
  let old_level = 1;
  let max_level = -1;

  //for (const el of s.split("").reverse()) {
  for (let i = s.length - 1; i >= 0; i--) {    
    // console.log(el)
    switch (s[i]) {
      case "I": {
        res = max_level > 1 ? res - 1 : res + 1;
        old_level = 1;
        break;
      }
      case "V": {
        res = max_level > 2 ? res - 5 : res + 5;
        old_level = 2;
        break;
      }
      case "X": {
        res = max_level > 3 ? res - 10 : res + 10;
        old_level = 3;
        break;
      }
      case "L": {
        res = max_level > 4 ? res - 50 : res + 50;
        old_level = 4;
        break;
      }

      case "C": {
        res = max_level > 5 ? res - 100 : res + 100;
        old_level = 5;
        break;
      }

      case "D": {
        res = max_level > 6 ? res - 500 : res + 500;
        old_level = 6;
        break;
      }

      case "M": {
        res += 1000;
        old_level = 7;
        break;
      }
    }

    max_level = Math.max(max_level, old_level);
  }
  return res;
}

console.log(romanToInt("MCMXCIV"));
