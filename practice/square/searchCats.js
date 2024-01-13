const searchCats = (cats, searchCriteria, searchValue, symbol) => {
  const predicates = {
    "=": (a, b) => (a === b),
    ">": (a, b) => (a > b),
    "<": (a, b) => (a < b),
  };

  const searchFields = {
    "WEIGHT": "weight",
    "HEIGHT": "height",
  };

  if (!predicates[symbol] || !searchFields[searchCriteria]) {
    return [];
  }

  const field = searchFields[searchCriteria];

  const cmp = (cat) => predicates[symbol](cat[field], searchValue);
  return cats.filter(cmp);
};

const test = () => {
  const names = ["a", "b", "c", "d", "e", "f", "g", "h"];
  const height = [31, 24, 67, 12, 45, 21, 31, 12];
  const weight = [120, 124, 160, 130, 175, 120, 124, 142];

  const len = Math.min(names.length, height.length, weight.length);
  const cats = [];

  for (let i = 0; i < len; i++) {
    cats.push({
      name: names[i],
      height: height[i],
      weight: weight[i],
    });
  }

  console.log(searchCats(cats, "HEIGHT", 43, "<"));
  console.log(searchCats(cats, "WEIGHT", 120, "="));
};

test();
