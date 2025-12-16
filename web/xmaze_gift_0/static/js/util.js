function drawPlayer() {
  let img = document.getElementById("player");

  // Set the sprite according to current direction
  if (isDead) {
    img.src = `/images/rip.png`;
  } else if (hasHat) {
    img.src = `/images/${playerDirection}-hat__.png`;
  } else {
    img.src = `/images/${playerDirection}__.png`;
  }
  img.alt = playerDirection;

  // Ensure we keep a reference to the DOM element
  if (!player) {
    player = img;
    player.style.position = "absolute";
  }

  player.style.left = `${playerPosition.x * cellSize + (50 - 35) / 2}px`;
  player.style.top = `${playerPosition.y * cellSize + (50 - 35) / 2}px`;
}

function getContentAtTile(tileX, tileY) {
  if (tileX < 0 || tileX >= gridSize || tileY < 0 || tileY >= gridSize) {
    return "empty";
  }
  return gridData[tileY * gridSize + tileX].content;
}

function getLookingTile() {
  let lookX = playerPosition.x;
  let lookY = playerPosition.y;
  switch (playerDirection) {
    case "up":
      lookY--;
      break;
    case "down":
      lookY++;
      break;
    case "left":
      lookX--;
      break;
    case "right":
      lookX++;
      break;
  }
  return { x: lookX, y: lookY };
}

function getTileAtPosition(x, y) {
  if (x < 0 || x >= gridSize || y < 0 || y >= gridSize) {
    return null;
  }

  const cells = grid.querySelectorAll(".cell");
  return cells[y * gridSize + x];
}

function blinkLookingTile() {
  let lookingTile = getTileAtPosition(getLookingTile().x, getLookingTile().y);
  if (lookingTile) {
    lookingTile.classList.add("highlighted");
  }
  setTimeout(() => {
    if (lookingTile) {
      lookingTile.classList.remove("highlighted");
    }
  }, 800);
}

function temperatureAtPlayer() {
  return gridData[playerPosition.y * gridSize + playerPosition.x].temperature;
}

function animateButtonPress(buttonId) {
  const label = document.getElementsByClassName(buttonId)[0];
  if (label) {
    void label.offsetWidth; // force reflow to restart animation if already flashing
    label.classList.add("pressed");
    setTimeout(() => label.classList.remove("pressed"), 200);
  }
}

function flashMessage(message) {
  const messageBox = document.getElementById("message-box");
  const messageText = document.getElementById("message-text");
  messageBox.style.display = "block";
  messageText.textContent = message;
  messageText.classList.add("slow-typing");
  isMessageDisplayed = true;
}

function removeMessage() {
  const messageBox = document.getElementById("message-box");
  const messageText = document.getElementById("message-text");
  if (messageText) {
    messageText.textContent = "";
  }
  if (messageBox) {
    messageBox.style.display = "none";
  }
  isMessageDisplayed = false;
}

function checkCollision(newX, newY) {
  if (newX < 0 || newX >= gridSize || newY < 0 || newY >= gridSize) {
    return true;
  }

  if (gridData[newY * gridSize + newX].content.startsWith("letter")) {
    return false;
  }

  if (gridData[newY * gridSize + newX].content !== "empty") {
    return true;
  }

  return false;
}

function clearTileAt(tileX, tileY) {
  const cells = grid.querySelectorAll(".cell");
  const cell = cells[tileY * gridSize + tileX];
  cell.style.backgroundImage = `url('/images/snow.png')`;
  cell.style.backgroundSize = "cover";
  cell.style.backgroundRepeat = "no-repeat";
  cell.style.backgroundPosition = "center";
  gridData[tileY * gridSize + tileX].content = "empty";
}

function addLetterToName(letter) {
  playerName += letter;
}
