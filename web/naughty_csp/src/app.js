const express = require("express");
const session = require("express-session");
const path = require("path");
const crypto = require("crypto");

const app = express();

app.engine("html", require("ejs").renderFile);
app.set("view engine", "html");
app.set("views", path.join(__dirname, "views"));

app.use(express.urlencoded({ extended: true }));
app.use(express.static("public"));

app.use(
   session({
      secret: crypto.randomBytes(32).toString("hex"),
      resave: false,
      saveUninitialized: false
   })
);

app.use((req, res, next) => {
   res.locals.nonce = crypto.randomBytes(16).toString("base64url");

   res.setHeader(
   "Content-Security-Policy",
   `script-src 'self' data: 'nonce-${res.locals.nonce}'; img-src 'self' data:; style-src 'self' 'unsafe-inline'; connect-src 'self'; object-src 'none';`
   );
   
   next();
});

app.use(require("./routes"));

app.listen(3000, () =>
   console.log("Naughty CSP instance running on port 3000")
);
