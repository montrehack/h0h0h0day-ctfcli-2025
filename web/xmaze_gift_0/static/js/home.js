const placedCoordinates = ["5,9"];

function getRandomInt(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

function getUniqueRandomCoordinates() {
  let x, y, coordinateKey;
  do {
    x = getRandomInt(0, gridSize - 1);
    y = getRandomInt(1, gridSize - 1);
    coordinateKey = `${x},${y}`;
  } while (placedCoordinates.includes(coordinateKey));

  placedCoordinates.push(coordinateKey);
  return { x, y };
}

async function startHome() {
  hasHat = false;
  playerPosition = { x: 5, y: 9 };
  playerDirection = "up";
  registeredName = true;
  drawPlayer();

  [...grid.querySelectorAll(".cell")].forEach((c) => c.remove());
  gridData = [];

  const bedPos = getUniqueRandomCoordinates();
  const tablePos = getUniqueRandomCoordinates();
  const chestPos = getUniqueRandomCoordinates();

  let content = "empty";
  for (let y = 0; y < gridSize; y++) {
    for (let x = 0; x < gridSize; x++) {
      const cell = document.createElement("div");
      cell.className = "cell";
      cell.style.width = `${cellSize}px`;
      cell.style.height = `${cellSize}px`;
      cell.style.position = "absolute";
      cell.style.left = `${x * cellSize}px`;
      cell.style.top = `${y * cellSize}px`;

      if (y === 0) {
        content = "wall";
        cell.style.backgroundImage = "url('/images/wall__.png')";
        if (x === 5) {
          content = "door";
          cell.style.backgroundImage = "url('/images/door__.png')";
        }
      } else if (x === tablePos.x && y === tablePos.y) {
        content = "table";
        cell.style.backgroundImage = "url('/images/table__.png')";
      } else if (x === bedPos.x && y === bedPos.y) {
        content = "bed";
        cell.style.backgroundImage = "url('/images/bed__.png')";
      } else if (x === chestPos.x && y === chestPos.y) {
        content = "chest";
        cell.style.backgroundImage = "url('/images/chest__.png')";
      } else {
        content = "empty";
        cell.style.backgroundImage = "url('/images/wood.png')";
      }
      cell.style.backgroundSize = "cover";
      cell.style.backgroundRepeat = "no-repeat";
      cell.style.backgroundPosition = "center";
      grid.appendChild(cell);
      gridData.push({ x: x, y: y, content: content, temperature: "hot" });
    }
  }

  drawPlayer();
  await sleep(500);
  flashMessage(`Welcome home ${playerName}!`);

  updateHeatBar();
}

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
