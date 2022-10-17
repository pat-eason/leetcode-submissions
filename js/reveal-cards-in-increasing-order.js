/**
 * You are given an integer array deck. There is a deck of cards
 * where every card has a unique integer. The integer on the ith
 * card is deck[i].
 *
 * You can order the deck in any order you want. Initially, all
 * the cards start face down (unrevealed) in one deck.
 *
 * You will do the following steps repeatedly until all cards are revealed:
 *
 * - Take the top card of the deck, reveal it, and take it out of the deck.
 * - If there are still cards in the deck then put the next top card of the
 *    deck at the bottom of the deck.
 * - If there are still unrevealed cards, go back to step 1. Otherwise, stop.
 * - Return an ordering of the deck that would reveal the cards in increasing order.
 *
 * Note that the first entry in the answer is considered to be the top of the deck.
 */

/**
 * @param {number[]} deck
 * @return {number[]}
 */
const deckRevealedIncreasing = (deck) => {
  if (!deck || deck.length === 0) {
    return [];
  }

  // sort the deck in order
  deck.sort((a, b) => a - b);

  // set result to the current max
  const result = [deck[deck.length - 1]];

  // loop through the deck in reverse
  for (let i = deck.length - 2; i >= 0; i--) {
    result.unshift(result.pop()); // remove last element from the array and place it at the beginning
    result.unshift(deck[i]); // place the current element at the beginning
  }

  return result;
};

console.log(
  deckRevealedIncreasing([17, 13, 11, 2, 3, 5, 7]),
  [2, 13, 3, 11, 5, 17, 7]
);
console.log(deckRevealedIncreasing([1, 1000]), [1, 1000]);
