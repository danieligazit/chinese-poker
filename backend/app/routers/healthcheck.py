from fastapi import APIRouter

router = APIRouter()


@router.get("/health", tags=["healthcheck"])
async def health():
    return {'status': 'available'}
