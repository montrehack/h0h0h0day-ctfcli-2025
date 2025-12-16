function startGame(maze) {
  // Reset player position and direction
  playerPosition = { x: 5, y: 9 };
  playerDirection = "up";
  mazeExploration = true;
  drawPlayer();

  // Remove only the cells, not the player
  [...grid.querySelectorAll(".cell")].forEach((c) => c.remove());
  gridData = [];

  // Draw grid cells
  for (let y = 0; y < gridSize; y++) {
    for (let x = 0; x < gridSize; x++) {
      const cell = document.createElement("div");
      cell.className = "cell";
      cell.style.width = `${cellSize}px`;
      cell.style.height = `${cellSize}px`;
      cell.style.position = "absolute";
      cell.style.left = `${x * cellSize}px`;
      cell.style.top = `${y * cellSize}px`;
      cell.style.backgroundImage = "url('/images/snow.png')";
      cell.style.backgroundSize = "cover";
      cell.style.backgroundRepeat = "no-repeat";
      cell.style.backgroundPosition = "center";
      grid.appendChild(cell);
      gridData.push({ x: x, y: y, content: "empty", temperature: "cold" });
    }
  }

  for (let y = 0; y < gridSize; y++) {
    for (let x = 0; x < gridSize; x++) {
      let cellContent = maze.maze.Grid[x][y];
      gridData[y * gridSize + x].content = cellContent;
      const cells = grid.querySelectorAll(".cell");
      const cell = cells[y * gridSize + x];
      if (cellContent === "tree") {
        cell.style.backgroundImage = `url('/images/tree-1__.png')`;
      } else if (cellContent === "hat") {
        cell.style.backgroundImage = `url('/images/hat__.png')`;
      } else if (cellContent === "gift") {
        cell.style.backgroundImage = `url('/images/gift__.png')`;
      } else if (cellContent === "scroll") {
        cell.style.backgroundImage = `url('/images/scroll__.png')`;
      }
    }
  }
}

function updateHeatBar() {
  const heatBar = document.getElementById("heat-bar");
  heatBar.innerHTML = "";
  for (let i = 0; i < heatLevel; i++) {
    const pip = document.createElement("span");
    pip.className = "pip";
    pip.textContent = "🔥";
    heatBar.appendChild(pip);
  }
  for (let i = heatLevel; i < MAX_HEAT_LEVEL; i++) {
    const pip = document.createElement("span");
    pip.className = "pip empty";
    pip.textContent = "❄️";
    heatBar.appendChild(pip);
  }
}

function consumeHeat() {
  if (heatLevel === 0) {
    rip();
  }
  if (heatLevel > 0 && temperatureAtPlayer() === "cold") {
    heatLevel--;
    updateHeatBar();
  }
}

async function rip() {
  isDead = true;
  hasHat = false;
  drawPlayer();
  await interactAPI("killPlayer");
  flashMessage(`${playerName} died of cold! Maybe his son will do better`);
  exit = await waitingKeypress();
  resetGame();
}

function randomTree() {
  for (let y = 0; y < gridSize; y++) {
    for (let x = 0; x < gridSize; x++) {
      if (x === 5 && y === 9) {
        continue;
      }

      if (Math.random() < 0.75) {
        continue;
      }
      const cells = grid.querySelectorAll(".cell");
      const cell = cells[y * gridSize + x];
      cell.style.backgroundImage = `url('/images/tree-1__.png')`;
      cell.style.backgroundSize = "cover";
      cell.style.backgroundRepeat = "no-repeat";
      cell.style.backgroundPosition = "center";
      gridData[y * gridSize + x].content = "tree";
    }
  }
}

function randomHat() {
  let y = 7;
  let x = 4;
  const cells = grid.querySelectorAll(".cell");
  const cell = cells[y * gridSize + x];
  cell.style.backgroundImage = `url('/images/hat__.png')`;
  cell.style.backgroundSize = "cover";
  cell.style.backgroundRepeat = "no-repeat";
  cell.style.backgroundPosition = "center";
  gridData[y * gridSize + x].content = "hat";
}

function interact(event) {
  switch (event.key.toLowerCase()) {
    case "q":
      if (isMessageDisplayed || isAnscestorsDisplayed) return;
      animateButtonPress("btn-a");

      blinkLookingTile();

      const target = getLookingTile();
      const item = getContentAtTile(target.x, target.y) || "empty";
      const action = tileActions[item];

      // Allow action functions to receive context if needed
      action({ target, item });
      break;
    case "e":
      animateButtonPress("btn-b");

      if (isMessageDisplayed) {
        removeMessage();
      } else if (temperatureAtPlayer() === "cold") {
        resetGame();
      }
      break;
  }
}

function movePlayer(event) {
  if (isDead || isMessageDisplayed) return;
  switch (event.key.toLowerCase()) {
    case "w":
      animateButtonPress("up");
      playerDirection = "up";
      collision = checkCollision(playerPosition.x, playerPosition.y - 1);
      if (!collision) {
        consumeHeat();
        playerPosition.y--;
      }
      break;
    case "s":
      animateButtonPress("down");
      playerDirection = "down";
      collision = checkCollision(playerPosition.x, playerPosition.y + 1);
      if (!collision) {
        consumeHeat();
        playerPosition.y++;
      }
      break;
    case "a":
      animateButtonPress("left");
      playerDirection = "left";
      collision = checkCollision(playerPosition.x - 1, playerPosition.y);
      if (!collision) {
        consumeHeat();
        playerPosition.x--;
      }
      break;
    case "d":
      animateButtonPress("right");
      playerDirection = "right";
      collision = checkCollision(playerPosition.x + 1, playerPosition.y);
      if (!collision) {
        consumeHeat();
        playerPosition.x++;
      }
      break;
  }
  logKeyPress(event.key.toLowerCase());
  drawPlayer();
}

async function logKeyPress(key) {
  if (isDead || !mazeExploration) return;
  keyPressed.push(key);
  data = await sendMove(key);

  if (easterEgg === "redacted") return; // Disable easter egg for public version

  if (data.message !== "") {
    flashMessage(data.message);
  }
}

function displayAncestors(data) {
  isAnscestorsDisplayed = true;
  let ancestors = data;
  const ancestorsDiv = document.getElementById("ancestors");
  ancestorsDiv.innerHTML = "";
  ancestors.forEach((ancestor) => {
    const ancestorInfo = document.createElement("div");
    ancestorInfo.className = "ancestor-info";
    ancestorInfo.innerHTML = `<strong>${ancestor.name}</strong><br>
      Born in: ${ancestor.birthDate}<br>Died at: ${ancestor.deathDate}<br>`;
    ancestorsDiv.appendChild(ancestorInfo);
  });
  const pressE = document.createElement("div");
  pressE.className = "instruction";
  pressE.innerHTML = `<br>Press E to continue...</div>`;
  ancestorsDiv.appendChild(pressE);
  ancestorsDiv.style.display = "block";
}

function hideAncestors() {
  const ancestorsDiv = document.getElementById("ancestors");
  ancestorsDiv.style.display = "none";
  ancestorsDiv.innerHTML = "";
  isAnscestorsDisplayed = false;
}

async function resetGame() {
  if (!isDead) {
    flashMessage(
      `Sacrifice ${playerName} to Santa and try again! Are you sure?`
    );
    exit = await waitingKeypress();
    if (!exit) {
      removeMessage();
      return;
    }
  }

  heatLevel = 5;
  isDead = false;
  gridData = [];
  keyPressed = [];
  isMessageDisplayed = false;
  mazeExploration = false;
  hasHat = false;
  removeMessage();
  selectName();
}

function waitingKeypress() {
  return new Promise((resolve) => {
    document.addEventListener("keydown", onKeyHandler);
    function onKeyHandler(e) {
      if (e.key.toLowerCase() === "e") {
        document.removeEventListener("keydown", onKeyHandler);
        return resolve(true);
      }
      if (e.key.toLowerCase() === "q") {
        document.removeEventListener("keydown", onKeyHandler);
        return resolve(false);
      }
    }
  });
}

document.addEventListener("keydown", movePlayer);
document.addEventListener("keydown", interact);
