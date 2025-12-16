from requests.adapters import BaseAdapter
from requests import Response
import subprocess
import re


class PythonAdapter(BaseAdapter):
    def send(
        self, request, stream=False, timeout=None, verify=True, cert=None, proxies=None
    ):
        code = request.url[9:]
        code = re.sub(r"[\(\)@]", "", code)

        proc = subprocess.run(
            ["timeout", "1s", "python", "-c", code], stdout=subprocess.PIPE
        )

        response = Response()
        response.status_code = proc.returncode
        response._content = proc.stdout

        return response

    def close(self):
        pass
