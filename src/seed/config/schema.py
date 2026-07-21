from pydantic import BaseModel
from typing import Any, Dict


class FieldConfig(BaseModel):
    type: str
    config: Dict[str, Any] = {}


class DatasetConfig(BaseModel):
    rows: int
    fields: Dict[str, FieldConfig]
