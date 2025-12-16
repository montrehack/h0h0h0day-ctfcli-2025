const puppeteer = require("puppeteer-core");

module.exports = async function runSantaBot(userId) {
   console.log("[SANTA BOT] Launching… user:", userId);

   const browser = await puppeteer.launch({
       headless: "new",
       executablePath: process.env.PUPPETEER_EXECUTABLE_PATH,
       args: ["--no-sandbox", "--disable-setuid-sandbox"]
   });

   const page = await browser.newPage();
   await page.setRequestInterception(true);
   page.on('request', req => req.continue());

   const target = `http://localhost:3000/wishlist/${userId}`;
   console.log("[SANTA BOT] Visiting:", target);

   await page.setCookie({
      name: "santa_secret",
      value: "flag{y0u_pwn3d_s4nt4_h0_h0_h0}",
      url: "http://localhost:3000",
      path: "/",
      httpOnly: false,
      secure: false
   });

   await page.goto(target, { waitUntil: "networkidle2" });
   await new Promise(r => setTimeout(r, 2000));

   await browser.close();
   console.log("[SANTA BOT] Done.");
};
