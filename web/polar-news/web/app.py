from collections import defaultdict
from flask import Flask, redirect, render_template, session, url_for, request
from xml.etree import ElementTree
import functools as ft
import itertools as it
import operator
import io
import requests
import base64

from python_adapter import PythonAdapter

AVAILABLE_SOURCES = [
    "www.lapresse.ca/manchettes/rss",
    "www.journaldemontreal.com/accueil/rss.xml",
    "www.ledevoir.com/rss/manchettes.xml",
    "rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
    "portswigger.net/research/rss",
    "labs.watchtowr.com/rss/",
]

app = Flask(__name__)
app.secret_key = open("secret_key.bin", "rb").read()

curl = requests.Session()
curl.mount("python://", PythonAdapter())


def resolve_feed(source_id: int) -> str:
    source = source_id
    if source_id >= 0 and source_id < len(AVAILABLE_SOURCES):
        source = AVAILABLE_SOURCES[source_id]

    source = f"http://{source}"
    return source


def hydrate_source(source_id: int) -> list[any]:
    source_url = resolve_feed(source_id)

    resp = curl.get(source_url)
    if resp.status_code != 200:
        return []

    try:
        tree = ElementTree.parse(io.BytesIO(resp.content))

        articles = []
        for item in it.islice(tree.findall(".//item"), 5):
            title = item.findtext("./title")
            link = item.findtext("./link")
            desc = item.findtext("./description")
            enclosure = item.find("./enclosure")

            if not title or not link or not desc:
                continue

            article = {"title": title, "link": link, "description": desc}

            if enclosure is not None:
                mime_type, url = enclosure.get("type"), enclosure.get("url")
                if mime_type and url and mime_type.startswith("image/"):
                    data_b64 = base64.b64encode(curl.get(url).content).decode()
                    article |= {"thumbnail": f"data:{mime_type};base64,{data_b64}"}

            articles.append(article)

        return articles
    except ElementTree.ParseError:
        return []


hydrated_sources = defaultdict(list)
for source_id in range(len(AVAILABLE_SOURCES)):
    hydrated_sources |= {source_id: hydrate_source(source_id)}


@app.get("/")
def home():
    source_ids = session.get("source_ids", [])
    articles = ft.reduce(
        operator.add, map(hydrated_sources.__getitem__, source_ids), []
    )
    return render_template("home.html", articles=articles)


@app.post("/add-source")
def add_source():
    global hydrated_sources
    source_ids = session.get("source_ids", [])

    source_id = request.form.get("source_id", None, type=int)
    if source_id is None:
        return redirect(url_for("home"))

    hydrated_sources |= {source_id: hydrate_source(source_id)}

    if source_id not in source_ids:
        source_ids.append(source_id)
        session["source_ids"] = source_ids

    return redirect(url_for("home"))


@app.get("/clear-sources")
def clear_sources():
    session["source_ids"] = []
    return redirect(url_for("home"))
