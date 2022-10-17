/**
 * @param {string} text
 * @return {string}
 */
const arrangeWords = (text) =>
  text
    .split(" ")
    .sort((a, b) => {
      if (a.length < b.length) {
        return -1;
      }
      if (a.length > b.length) {
        return 1;
      }
      return 0;
    })
    .map((x, index) => {
      return index === 0
        ? `${x.charAt(0).toUpperCase()}${x.slice(1)}`
        : x.toLowerCase();
    })
    .join(" ");

// const testCase = "Leetcode is cool";
// const testCase = "Keep calm and code on";
const testCase = "To be or not to be";
const result = arrangeWords(testCase);
console.log("result:", `"${result}"`);
