function selectName() {
  playerPosition = { x: 5, y: 9 };
  playerDirection = "up";
  playerName = "";
  registeredName = false;
  drawPlayer();

  // Remove only the cells, not the player
  [...grid.querySelectorAll(".cell")].forEach((c) => c.remove());
  gridData = [];

  // Draw grid cells
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

      const totalLetters = 26;
      const cols = Math.min(gridSize, Math.ceil(Math.sqrt(totalLetters))); // width of letter block
      const rows = Math.ceil(totalLetters / cols); // height of letter block
      const startX = Math.floor((gridSize - cols) / 2);
      const startY = Math.floor((gridSize - rows) / 2);

      // center text inside cell and add side padding
      cell.style.display = "flex";
      cell.style.alignItems = "center";
      cell.style.justifyContent = "center";
      cell.style.boxSizing = "border-box";
      cell.style.padding = "6px";
      cell.style.fontSize = "24px";
      cell.style.fontWeight = "bold";
      cell.style.color = "#333";
      cell.style.textShadow = "2px 1px 2px #fff";

      cell.style.backgroundImage = "url('/images/wood.png')";

      const localX = x - startX;
      const localY = y - startY;
      const letterIndex = localY * cols + localX;
      if (
        localX >= 0 &&
        localX < cols &&
        localY >= 0 &&
        localY < rows &&
        letterIndex < totalLetters
      ) {
        const letter = String.fromCharCode(65 + letterIndex); // A-Z
        cell.textContent = letter;
        content = "letter-" + letter;
      } else if (y === 0) {
        content = "wall";
        cell.style.backgroundImage = "url('/images/wall__.png')";
        if (x === 5) {
          content = "start";
          cell.style.backgroundImage = "url('/images/door__.png')";
        }
      } else {
        cell.textContent = "";
        content = "empty";
      }
      cell.style.backgroundSize = "cover";
      cell.style.backgroundRepeat = "no-repeat";
      cell.style.backgroundPosition = "center";
      grid.appendChild(cell);
      gridData.push({ x: x, y: y, content: content, temperature: "hot" });
    }
  }

  updateHeatBar();
}

window.onload = selectName;
