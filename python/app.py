from fastapi import FastAPI, Query
from fastapi.responses import PlainTextResponse
import asyncio

app = FastAPI()

@app.get("/", response_class=PlainTextResponse)
async def root(sleep: int = Query(10, description="Time to sleep in seconds")):
    await asyncio.sleep(sleep)
    return f"{sleep} seconds have passed"
