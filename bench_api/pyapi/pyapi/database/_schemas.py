# Pydantic Models

from pydantic import BaseModel


class ItemInput(BaseModel):
    name: str
    description: str
    mass_grams: int


class Item(ItemInput):
    id: int

    class Config:
        orm_mode = True
