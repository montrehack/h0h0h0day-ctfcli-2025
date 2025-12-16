const sqlite3 = require("sqlite3").verbose();
const path = require("path");
const { randomBytes, randomUUID } = require("crypto");

const db = new sqlite3.Database(path.join(__dirname, "data.db"));

db.serialize(() => {
   db.run(`
       CREATE TABLE IF NOT EXISTS users (
           id INTEGER PRIMARY KEY AUTOINCREMENT,
           username TEXT UNIQUE,
           password TEXT,
           wishlistId TEXT,
           content TEXT
       )
   `);
});

function registerUser(username, password) {
   return new Promise((resolve, reject) => {
      db.run(`
       INSERT INTO users (username, password, wishlistId, content) VALUES (?, ?, ?, ?)`,
       [username, password, randomUUID(), ""], 
       (err, row) => {
           if (err) return reject(err);
           resolve({ register: true });
       });
   });
}

function findUser(username) {
   return new Promise((resolve, reject) => {
      db.get(`
       SELECT * FROM users WHERE username = ?`,
       [username], 
       (err, row) => {
           if (err) return reject(err);
           resolve(row);
       });
   });
}

function findWishlist(wishlistId) {
   return new Promise((resolve, reject) => {
      db.get(`
       SELECT * FROM users WHERE wishlistId = ?`,
       [wishlistId], 
       (err, row) => {
           if (err) return reject(err);
           resolve(row);
       });
   });
}

function loginUser(username, password) {
   return new Promise((resolve, reject) => {
      db.get(`
       SELECT * FROM users WHERE username = ? AND password = ?`,
       [username, password], 
       (err, row) => {
          if (err) return reject(err);
          resolve(row);
       });
   });
}

function insertWishlist(content, username) {
   return new Promise((resolve, reject) => {
      db.run(`
       UPDATE users SET content = ? WHERE username = ?`,
       [content, username],
       (err, row) => {
          if (err) return reject(err);
          resolve({ updated: true });
       });
   });
}


module.exports = {
   db,
   registerUser,
   loginUser,
   findUser,  
   insertWishlist,
   findWishlist
};
