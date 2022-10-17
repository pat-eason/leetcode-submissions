/**
 * @param {string} s
 * @return {number}
 */
const lengthOfLongestSubstring = (s) => {
  const characterArray = s.split("");

  let result = 0;
  let latestCharacterArray = [];
  let latestResult = result;
  let latestIndex = 0;

  /**
   * iterate through each character.
   *  - if character does not exist in current array then increment
   *    latest result, push character to local array, and continue
   *
   *  - if character does exist in current array, then update latestIndex
   *    to go back to the next character, reset latest result, set the
   *    output result to the max of latest and current output
   */

  for (let i = latestIndex; i < characterArray.length; i++) {
    const currentCharacter = characterArray[i];

    if (!latestCharacterArray.includes(currentCharacter)) {
      latestResult++;
      latestCharacterArray.push(currentCharacter);
    } else {
      i = latestIndex;
      latestIndex = latestIndex + 1;
      latestResult = 0;
      latestCharacterArray = [];
    }
    result = Math.max(result, latestResult);
  }

  return result;
};

console.log(lengthOfLongestSubstring("abcabcbb"), 3);
console.log(lengthOfLongestSubstring("bbbbb"), 1);
console.log(lengthOfLongestSubstring("pwwkew"), 3);
console.log(lengthOfLongestSubstring(" "), 1);
console.log(lengthOfLongestSubstring(""), 0);
console.log(lengthOfLongestSubstring("aab"), 2);
console.log(lengthOfLongestSubstring("dvdf"), 3);
