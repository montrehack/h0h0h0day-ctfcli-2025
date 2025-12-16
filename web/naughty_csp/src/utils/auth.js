function requireLogin(req, res, next) {
    if (!req.session.user) return res.redirect("/");
    next();
}

function isValidEmail(email) {
    const emailRegex = /^.+@.+$/;
    return (
        typeof email === "string" &&
        emailRegex.test(email)
    );
}

function isValidLength(email) {
    return (
        email.length <= 40
    );
}

module.exports = { requireLogin, isValidEmail, isValidLength };
