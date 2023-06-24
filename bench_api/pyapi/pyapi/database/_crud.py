from typing import Optional

from sqlalchemy.orm import Session

from . import _models as models
from . import _schemas as schemas


def get_items(db: Session, *, offset: int = 0, limit: int = 100) -> list[schemas.Item]:
    db_items = db.query(models.Item).offset(offset).limit(limit).all()
    return [schemas.Item.from_orm(item) for item in db_items]


def get_item(db: Session, *, item_id: int) -> Optional[schemas.Item]:
    item = db.query(models.Item).filter(models.Item.id == item_id).first()
    return schemas.Item.from_orm(item) if item is not None else None


def create_item(db: Session, *, item: schemas.ItemInput) -> schemas.Item:
    db_item = models.Item(**item.dict())
    db.add(db_item)
    db.commit()
    db.refresh(db_item)
    return schemas.Item.from_orm(db_item)
