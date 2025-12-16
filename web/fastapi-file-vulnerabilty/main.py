from fastapi import FastAPI, Query, Request, UploadFile, File
from pathlib import Path
import shutil
import urllib

from starlette.responses import FileResponse

app = FastAPI(
    docs_url=None, 
    redoc_url=None,
    openapi_url=None,
)

IMAGES_DIR = Path("images")
IMAGES_DIR.mkdir(exist_ok=True)

@app.post("/upload")
async def upload_file(file: UploadFile = File(...)):
    file_path = IMAGES_DIR / file.filename

    with file_path.open("wb") as buffer:
        shutil.copyfileobj(file.file, buffer)

    return {"filename": file.filename}


@app.get("/image/")
def get_image(request: Request, name: str = Query(...)):
    image_path = "images/" + name 

    print("IMAGE PATH: ", image_path)
    print("IMAGE PATH: ", image_path)
    print("IMAGE PATH: ", image_path)
    raw_name = request.scope["query_string"]
    if ".." in str(raw_name):
        return {"error": "invalid filename"}
    return FileResponse(image_path)

@app.get("/")
def home():
    return FileResponse("index.html")

@app.get("/images")
def list_images():
    return [p.name for p in IMAGES_DIR.iterdir() if p.is_file()]
