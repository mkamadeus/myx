from typing import List
from fastapi import FastAPI, UploadFile
from keras.models import load_model
from pydantic import BaseModel
import numpy as np
import io
from PIL import Image

app = FastAPI()
model = load_model("mnist.h5")


class Output(BaseModel):
	prediction : List[float]
	


@app.post('/')
async def root(file : UploadFile):
    contents = await file.read()
    stream = io.BytesIO(contents)
    img = Image.open(stream)
    img = img.resize((32,32))

    result = np.array(img).reshape((1,32,32,1))

    prediction = model.predict(x=result).tolist()[0]

    return [Output(prediction=prediction)]
