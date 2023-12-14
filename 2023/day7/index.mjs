import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf-8")
  .split("\n")
  .map((line) => line.split(" "));

const problem1 = (input) => {
  const cardMap = [
    "A",
    "K",
    "Q",
    "J",
    "T",
    "9",
    "8",
    "7",
    "6",
    "5",
    "4",
    "3",
    "2",
  ];

  const getHandValue = (hand) => {
    const cards = hand.split("").reduce((acc, card) => {
      if (acc[card]) {
        acc[card] += 1;
      } else {
        acc[card] = 1;
      }
      return acc;
    }, {});

    const cardsObj = Object.entries(cards);

    // 5 of a kind
    if (cardsObj.length === 1) {
      return 0;
    }
    // 4 of a kind
    if (cardsObj.length === 2 && cardsObj.find((card) => card[1] === 4)) {
      return 1;
    }
    // Full house
    if (cardsObj.length === 2 && cardsObj.find((card) => card[1] === 3)) {
      return 2;
    }
    // 3 of a kind
    if (cardsObj.find((card) => card[1] === 3)) {
      return 3;
    }
    // 2 pair
    if (cardsObj.filter((card) => card[1] === 2).length === 2) {
      return 4;
    }
    // 1 pair
    if (cardsObj.find((card) => card[1] === 2)) {
      return 5;
    }
    // High card
    return 6;
  };

  const sortHands = (a, b) => {
    const [aHand, aBid] = a;
    const [bHand, bBid] = b;

    const aHandValue = getHandValue(aHand);

    const bHandValue = getHandValue(bHand);

    if (aHandValue === bHandValue) {
      for (let i = 0; i < aHand.length; i++) {
        if (aHand.charAt(i) === bHand.charAt(i)) {
          continue;
        }
        return (
          cardMap.indexOf(aHand.charAt(i)) - cardMap.indexOf(bHand.charAt(i))
        );
      }
    }
    return aHandValue - bHandValue;
  };

  const sorted = input.sort(sortHands).reverse();

  let sum = 0;
  for (let i = 1; i < sorted.length + 1; i++) {
    const [, bid] = sorted[i - 1];
    const bidValue = parseInt(bid);
    //   console.log(`${hand}: ${bidValue} * ${i}`);
    sum += bidValue * i;
  }

  console.log(sum);
};

const problem2 = (input) => {
  const cardMap = [
    "A",
    "K",
    "Q",
    "T",
    "9",
    "8",
    "7",
    "6",
    "5",
    "4",
    "3",
    "2",
    "J",
  ];

  const getHandValue = (hand) => {
    const cards = hand.split("").reduce((acc, card) => {
      if (acc[card]) {
        acc[card] += 1;
      } else {
        acc[card] = 1;
      }
      return acc;
    }, {});

    const cardsObj = Object.entries(cards).sort((a, b) => {
      return b[1] - a[1];
    });
    const joker = cardsObj.find((card) => card[0] === "J");

    const removeJoker = cardsObj.filter((card) => card[0] !== "J");
    if (joker) {
      const jokerCount = joker[1];
      if (removeJoker.length > 0) {
        removeJoker[0][1] += jokerCount;
      } else {
        removeJoker.push(joker);
      }
    }

    // 5 of a kind
    if (removeJoker.length === 1) {
      return 0;
    }
    // 4 of a kind
    if (removeJoker.length === 2 && removeJoker.find((card) => card[1] === 4)) {
      return 1;
    }
    // Full house
    if (removeJoker.length === 2 && removeJoker.find((card) => card[1] === 3)) {
      return 2;
    }
    // 3 of a kind
    if (removeJoker.find((card) => card[1] === 3)) {
      return 3;
    }
    // 2 pair
    if (removeJoker.filter((card) => card[1] === 2).length === 2) {
      return 4;
    }
    // 1 pair
    if (removeJoker.find((card) => card[1] === 2)) {
      return 5;
    }
    // High card
    return 6;
  };

  const sortHands = (a, b) => {
    const [aHand, _aBid] = a;
    const [bHand, _bBid] = b;

    const aHandValue = getHandValue(aHand);

    const bHandValue = getHandValue(bHand);

    if (aHandValue === bHandValue) {
      for (let i = 0; i < aHand.length; i++) {
        if (aHand.charAt(i) === bHand.charAt(i)) {
          continue;
        }
        return (
          cardMap.indexOf(aHand.charAt(i)) - cardMap.indexOf(bHand.charAt(i))
        );
      }
    }
    return aHandValue - bHandValue;
  };

  const sorted = input.sort(sortHands).reverse();

  let sum = 0;
  for (let i = 1; i < sorted.length + 1; i++) {
    const [hand, bid] = sorted[i - 1];
    const bidValue = parseInt(bid);
    //   console.log(`${hand}: ${bidValue} * ${i}`);
    sum += bidValue * i;
  }

  console.log(sum);
};

problem2(input);
