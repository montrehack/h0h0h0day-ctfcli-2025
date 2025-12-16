from fastapi import FastAPI, Request
from fastapi.responses import HTMLResponse
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates
from fastapi import Depends
from elasticsearch import Elasticsearch
import os

app = FastAPI(title="HoHoHo Blackmarket")

BASE_DIR = os.path.dirname(os.path.abspath(__file__))
templates = Jinja2Templates(directory=os.path.join(BASE_DIR, "templates"))
app.mount("/static", StaticFiles(directory=os.path.join(BASE_DIR, "static")), name="static")

INDEX_NAME = "blackmarket"
FLAG = os.getenv("FLAG", "FLAG-2b42567a794646178227a6a62679274c")

def get_os_client():
    host = os.getenv("OPENSEARCH_HOST", "http://localhost:9200")
    # security disabled in docker-compose, so no auth
    client = Elasticsearch(hosts=[host])
    return client


@app.on_event("startup")
async def startup_event():
    client = get_os_client()
    # create index with hidden field
    mapping = {
        "settings": {
            "index": {
                "number_of_shards": 1,
                "number_of_replicas": 0
            }
        },
        "mappings": {
            "properties": {
                "name": {"type": "text"},
                "price": {"type": "integer"},
                "description": {"type": "text"},
                "category": {"type": "keyword"},
                "flag": {"type": "keyword"}  # hidden field
            }
        }
    }
    exists = client.indices.exists(INDEX_NAME)
    if not exists:
        client.indices.create(index=INDEX_NAME, body=mapping)
        # seed data
        items = [
            {
            "name": "Lockpick Set",
            "price": 50,
            "description": "For opening doors, not hearts.",
            "category": "tools",
            "flag": FLAG
            },
            {
            "name": "RC Drone",
            "price": 120,
            "description": "Surveillance made easy.",
            "category": "toys",
            "flag": FLAG
            },
            {
            "name": "Soldering Kit",
            "price": 35,
            "description": "Fix your circuits, or break them.",
            "category": "tools",
            "flag": FLAG
            },
            {
            "name": "Flipper zero",
            "price": 200,
            "description": "Multi-tool for pentesters and geeks.",
            "category": "tools",
            "flag": FLAG
            },
            {
            "name": "RFID Cloner",
            "price": 80,
            "description": "Copy and emulate RFID badges.",
            "category": "gadgets",
            "flag": FLAG
            },
            {
            "name": "Signal Jammer",
            "price": 150,
            "description": "Disrupt wireless communications.",
            "category": "gadgets",
            "flag": FLAG
            },
            {
            "name": "Spy Camera",
            "price": 60,
            "description": "Tiny camera for covert surveillance.",
            "category": "toys",
            "flag": FLAG
            },
            {
            "name": "USB Rubber Ducky",
            "price": 90,
            "description": "Automate keystrokes for pranks or pentests.",
            "category": "tools",
            "flag": FLAG
            },
        ]
        for i, doc in enumerate(items):
            client.index(index=INDEX_NAME, id=i+1, body=doc, refresh=True)


@app.get("/", response_class=HTMLResponse)
async def home(request: Request):
    return templates.TemplateResponse("index.html", {"request": request})


@app.get("/search", response_class=HTMLResponse)
async def search(request: Request, q: str = ""):
    try:
        client = get_os_client()
        # naive vulnerable pass-through to Lucene query via q param
        res = client.search(index=INDEX_NAME, q=q, size=10)
        hits = []
        for h in res.get("hits", {}).get("hits", []):
            src = h.get("_source", {})
            # mask hidden fields from response
            src.pop("flag", None)
            hits.append(src)
        return templates.TemplateResponse("partials/results.html", {"request": request, "hits": hits})
    except Exception as e:
        return templates.TemplateResponse("partials/results.html", {"request": request, "hits": []})
