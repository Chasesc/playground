from sqlalchemy.orm import Session  # noqa F401

from . import _crud as crud  # noqa F401
from . import _models as models
from . import _schemas as schemas  # noqa F401
from ._engine import SessionLocal, engine

models.Base.metadata.create_all(engine)


def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
