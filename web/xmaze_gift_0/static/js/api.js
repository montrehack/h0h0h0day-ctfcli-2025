const API_URL = "/api";
const MOVE_URL = "/move";

let meteoValues = [
  // "sunny",
  "foggy",
];


// this API has a random sleep on the server to prevent brute force
// DO NOT BRUTE FORCE THIS API, seriously, it won't help 
async function interactAPI(action) {
  if (!playerName || playerName == "") {
    return;
  }
  const data = {
    name: playerName,
    meteo: meteoValues[0],
  };

  selectedMeteo = meteoValues[0];

  let url = `${API_URL}/${action}`;

  let method = "GET";
  let body = undefined;
  if (
    action === "newPlayer" ||
    action === "killPlayer" ||
    action === "explore" ||
    action === "gift" ||
    action === "hat" ||
    action === "scroll"
  ) {
    method = "POST";
    body = JSON.stringify(data);
  } else if (action === "gift" || action === "hat" || action === "scroll") {
  } else {
    url += "?name=" + playerName;
  }

  return await fetch(url, {
    method: method,
    headers: {
      "Content-Type": "application/json",
    },
    body,
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .then((data) => {
      return data;
    })
    .catch((error) => console.error("New Player error:", error));
}

async function sendMove(move) {
  if (!playerName || playerName == "") {
    return;
  }

  if (!registeredName) {
    return;
  }

  const data = {
    name: playerName,
    move: move,
  };

  return await fetch(`${API_URL}${MOVE_URL}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    })
    .then(async (data) => {
      if (data.message != "") {
        flashMessage(data.message);
        await waitingKeypress();
      }
    })
    .catch((error) => console.error("Move error:", error));
}
