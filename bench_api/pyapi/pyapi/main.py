from database import Session, crud, get_db, schemas
from fastapi import Depends, FastAPI, HTTPException

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/items/", response_model=list[schemas.Item])
def get_items(offset: int = 0, limit: int = 100, db: Session = Depends(get_db)):
    return crud.get_items(db, offset=offset, limit=limit)


@app.get("/items/{item_id}", response_model=schemas.Item)
def get_item(item_id: int, db: Session = Depends(get_db)):
    item = crud.get_item(db, item_id=item_id)
    if item is None:
        raise HTTPException(status_code=404, detail="Item not found")

    return item


@app.post("/items/", response_model=schemas.Item)
def create_item(item: schemas.ItemInput, db: Session = Depends(get_db)):
    return crud.create_item(db, item=item)
