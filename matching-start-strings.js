const matchingStartStrings = (starters, values) => {
  const output = [];

  const sanitizedStarters = starters.map((x) => x.toLowerCase());
  const sanitizedValues = values.map((x) => x.toLowerCase());

  // iterate through starters
  for (let i = 0; i < sanitizedStarters.length; i++) {
    const currentStarter = sanitizedStarters[i];
    const filteredResults = sanitizedValues.filter((x) =>
      x.startsWith(currentStarter)
    );

    if (filteredResults.length > 0) {
      output.push(...filteredResults);
    }
  }

  return output.map((x) => x.toUpperCase());
};
