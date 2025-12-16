const tileActions = {
  tree: () => flashMessage("This tree looks tasty, but you are on a diet."),
  empty: () => flashMessage(".........."),
  gift: async () => {
    data = await interactAPI("gift");
    flashMessage(`${data.message}`);
  },
  tombstone: () => flashMessage("Someone died here."),
  door: async () => {
    flashMessage("Are you sure you want go out into the cold maze?");
    exit = await waitingKeypress();
    if (exit) {
      data = await interactAPI("explore");
      startGame(data);
      if (selectedMeteo === "foggy") {
        flashMessage(
          "Where is the gift? I should come back when the weather is better."
        );
      } else if (data.message !== "") {
        flashMessage(`${data.message}`);
      }
    }
  },
  table: async () => {
    flashMessage("I wonder who was the last elf to venture into the maze..");
    see = await waitingKeypress();
    if (see) {
      data = await interactAPI("ancestors");
      displayAncestors(data);
      await waitingKeypress();
      hideAncestors();
    }
    removeMessage();
  },
  wall: () => {
    flashMessage("And that's a wall... Yeah a wall!");
  },
  chest: () =>
    flashMessage(
      "Someone stole the gift from the chest! Go find it in the maze!"
    ),
  bed: () =>
    flashMessage(
      "No time to sleep, go search for Santa Claus lost XMaze Gift."
    ),
  hat: async ({ target }) => {
    data = await interactAPI("hat");
    flashMessage("You found a christmas hat. You feel warmer.");
    heatLevel = Math.min(heatLevel + 5, MAX_HEAT_LEVEL);
    hasHat = true;
    drawPlayer();
    clearTileAt(target.x, target.y);
    updateHeatBar();
  },
  scroll: async ({ target }) => {
    data = await interactAPI("scroll");
    flashMessage(`${data.message}`);
    clearTileAt(target.x, target.y);
  },
  start: async () => {
    if (playerName === "") {
      flashMessage("What is you name?");
      return;
    }
    flashMessage(`Are you ${playerName}?`);
    exit = await waitingKeypress();
    if (exit) {
      data = await interactAPI("newPlayer");
      if (data.name !== playerName) {
        flashMessage(`You are not ${playerName}, he died earlier`);
        await waitingKeypress();
        playerName = "";
      } else {
        startHome();
      }
    }
    removeMessage();
  },
};

"ABCDEFGHIJKLMNOPQRSTUVWXYZ".split("").forEach((letter) => {
  tileActions[`letter-${letter}`] = () => {
    addLetterToName(letter.toLowerCase());
    flashMessage(`Name: ${playerName}`);
  };
});
