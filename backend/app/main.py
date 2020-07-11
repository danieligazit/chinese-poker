import uvicorn
from fastapi import FastAPI
from starlette.middleware.cors import CORSMiddleware
from routers import healthcheck
from core.config import settings

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}
    
app.add_middleware(
    CORSMiddleware,
    allow_origins=[str(origin) for origin in settings.BACKEND_CORS_ORIGINS],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(healthcheck.router)
