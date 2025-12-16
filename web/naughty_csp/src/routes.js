const express = require("express");
const router = express.Router();
const { db, registerUser, findUser, loginUser, insertWishlist, findWishlist } = require("./utils/db");
const { requireLogin, isValidEmail, isValidLength } = require("./utils/auth");
const runSantaBot = require("./utils/santabot");

// Main page, redirect to login if user isn't signed in
router.get("/", (req, res) => {
   if (req.session.user) {
      res.redirect("wishlist");
   } else {
      res.render("index");
   }
}); 

// Register page
router.get("/register", (req, res) => res.render("register"));

// POST logic for login
router.post("/login", async (req, res) => {
   const { username, password } = req.body;
   const user = await loginUser(username, password);
   if (user) {
      req.session.user = user;
      return res.status(200).json({ redirect: "/wishlist", message: "Login successful!" });
   } else {
      return res.status(401).json({ message: "Invalid logon credentials" });
   }
});

// POST logic for register
router.post("/register", async (req, res) => {
   const { username, password } = req.body;
   
   if (!isValidEmail(username)) {
      return res.status(400).json({ message: "Invalid email address format" });
   }
   
   if (!isValidLength(username)) {
      return res.status(400).json({ message: "Email must be shorter than 40 characters" });
   }

   if (await findUser(username)) {
      return res.status(401).json({ message: "User already exists" });
   }
   
   const created_user = await registerUser(username, password);
   if (created_user.register) {
      req.session.user = username;
      const user = await findUser(username);
      req.session.wishlistId = user.wishlistId;
      return res.status(200).json({ redirect: "/wishlist", message: "Registration successful!" });
   } else {
      return res.status(500).json({ message: "Error creating user" });
   }
});

// Logic for logout function
router.get("/logout", (req, res) => {
    req.session.destroy(err => {
        if (err) {
            console.error("Logout error:", err);
            return res.status(500).send("Error logging out");
        }
        res.redirect("/");
    });
});

// Rendering wishlist platform
router.get("/wishlist", requireLogin, async (req, res) => {
   const user = req.session.user;
   const list = await findUser(user);
   if (list) {
      return res.render("wishlist", { wishlist: list.content});
   } else {
      return res.render("wishlist", { wishlist: ""});
   }
});

// Update wishlist logic
router.post("/wishlist", requireLogin, async (req, res) => {
   const content = req.body.item;
   const user = req.session.user;
   
   const wishlist = await insertWishlist(content, user);
   if (wishlist) {
      return res.status(200).json({ message: "Wishlist successfully updated"});
   } else {
      return res.status(401).json({ message: "Error updating wishlist" });
   }
});

// Viewing wishlist logic
router.get("/wishlist/:id", async (req, res) => {
   const id = req.params.id;
   const wishlist = await findWishlist(id);
   if (wishlist) {
      return res.render("user-wishlist", { content: wishlist.content, username: wishlist.username.split("@")[0] });
   } else {
      return res.status(403).json({ message: "Wishlist inaccessible" });
   }
});

// Endpoint used to return the wishlist from the session
router.get("/me", requireLogin, async (req, res) => {
   const user = await findUser(req.session.user);
   if (user) {
      return res.status(200).json({username : user.username, wishlistId: user.wishlistId});
   } else {
      return res.status(403).json({ message: "Wishlist inaccessible" });
   }
});

// Sending wishlist to santa based on user session
router.get("/send", requireLogin, async (req, res) => {
   try {
      const uid = req.session.wishlistId;
      await runSantaBot(uid);
      return res.status(200).json({ message: "Santa will view your wishlist soon..." });
   } catch (err) {
      console.error(err);
      return res.status(500).json({ message: "Error: " + err });
   }
});


module.exports = router;
