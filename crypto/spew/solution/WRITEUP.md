# S.P.E.W

Exploit: The hash for the zip file sort the name and content without keeping the link between them. This means we can
    switch the content of index.html and flag.html to get the flag.

```shell
curl -v -X POST http://localhost:8080/update_website -F "theme=@./solution.zip"
```
