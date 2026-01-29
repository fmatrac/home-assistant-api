import logging
from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from fastapi.responses import FileResponse
from app.config import settings
from app.routes import api_router

# Configure logging
logging.basicConfig(
    level=getattr(logging, settings.log_level.upper()),
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
)
logger = logging.getLogger(__name__)

app = FastAPI(title="Home Assistant API", version="1.0.0")

# Include API routes
app.include_router(api_router)


# Health check endpoint
@app.get("/health")
def health_check():
    return {"status": "OK"}


# Mount static files for frontend
app.mount("/css", StaticFiles(directory="web/css"), name="css")
app.mount("/js", StaticFiles(directory="web/js"), name="js")


# Serve index.html for root
@app.get("/")
def serve_frontend():
    return FileResponse("web/index.html")


if __name__ == "__main__":
    import uvicorn

    logger.info(f"Starting server on port {settings.server_port}")
    uvicorn.run(app, host="0.0.0.0", port=settings.server_port)
