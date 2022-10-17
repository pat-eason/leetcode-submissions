const matchesInStringArray = (matches, values) => {
  const output = [];

  const sanitizedMatches = matches.map((x) => x.toLowerCase());
  const sanitizedValues = values.map((x) => x.toLowerCase());

  for (let i = 0; i < sanitizedMatches.length; i++) {
    const currentMatch = sanitizedMatches[i];
    const filteredValues = sanitizedValues.filter((x) => x === currentMatch);
    if (filteredValues.length > 0) {
      output.push(...filteredValues);
    }
  }
  return output;
};

console.log(
  matchesInStringArray(["abc", "xyz", "789456"], ["xyz", "kjhgkjhgkjhg"]),
  ["xyz"]
);
