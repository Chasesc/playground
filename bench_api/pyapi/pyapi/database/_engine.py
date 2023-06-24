from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

engine = create_engine("postgresql+psycopg://benchapi:benchapi@localhost:5432/items")
SessionLocal = sessionmaker(bind=engine)
