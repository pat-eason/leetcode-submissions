/**
 * On day 1, one person discovers a secret.
 *
 * You are given an integer delay, which means that each
 * person will share the secret with a new person every day,
 * starting from delay days after discovering the secret.
 * You are also given an integer forget, which means that each
 * person will forget the secret forget days after discovering
 * it. A person cannot share the secret on the same day they
 * forgot it, or on any day afterwards.
 *
 * Given an integer n, return the number of people who know the
 * secret at the end of day n. Since the answer may be very large,
 * return it modulo 109 + 7.
 */

/**
 * @param {number} n
 * @param {number} delay
 * @param {number} forget
 * @return {number}
 */
const peopleAwareOfSecret = (n, delay, forget) => {
  let result = 0;

  const peopleState = new Array(n + forget);
  for (let i = 0; i < n; i++) {
    peopleState[i] = new Array(3);
  }
  const modulus = 1e9 + 7;

  peopleState[0][0] = 1;
  peopleState[delay][1] = 1;
  peopleState[forget][2] = 1;

  console.log(peopleState);

  // skip first entry
  for (let i = 1; i < n; i++) {
    console.log("iterate", i);
    console.log(
      peopleState[i][1],
      peopleState[i + forget][2],
      peopleState[i + delay][1]
    );

    peopleState[i][1] =
      (peopleState[i][1] +
        peopleState[i - 1][1] -
        peopleState[i][2] +
        modulus) %
      modulus;

    peopleState[i + forget][2] = peopleState[i][1];

    peopleState[i + delay][1] = peopleState[i][1];

    peopleState[i][0] =
      (peopleState[i - 1][0] +
        peopleState[i][1] -
        peopleState[i][2] +
        modulus) %
      modulus;
  }

  console.log(peopleState);
  return peopleState[n - 1][0];
};

console.log(peopleAwareOfSecret(6, 2, 4), 5);
// console.log(peopleAwareOfSecret(4, 1, 3), 6);
