from fastapi import FastAPI, Query, HTTPException
from fastapi.responses import PlainTextResponse
import httpx

app = FastAPI()

@app.get("/", response_class=PlainTextResponse)
async def root(sleep: int = Query(..., description="Time to sleep in seconds")):
    url = f"http://rust:3000/?sleep={sleep}"

    try:
        async with httpx.AsyncClient() as client:
            response = await client.get(url)

        if response.status_code == 200:
            return response.text
    except httpx.RequestError as e:
        pass

    return PlainTextResponse("Failed to fetch data", status_code=500)
