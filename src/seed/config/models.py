from pydantic import BaseModel
from typing import Any
class GeneratorConfig(BaseModel):
 type:str
 config:dict[str,Any]={}
