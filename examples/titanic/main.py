from fastapi import FastAPI
from pydantic import BaseModel
import numpy as np

# input imports

# output imports

# model imports
import onnxruntime as rt

# pipeline imports
import joblib


app = FastAPI()
model = rt.InferenceSession("titanic.onnx")


class Input(BaseModel):
	pclass : int
	sibsp : int
	parch : int
	age : float
	gender : str
	embarked : str
	

class Output(BaseModel):
	prediction : float
	


@app.post('/')
async def root(body: Input):
    t0 =  np.array([[body.pclass]])
    t1 =  np.array([[body.sibsp]])
    t2 =  np.array([[body.parch]])
    scaler = joblib.load('age.scaler')
    scaled = scaler.transform(np.array([[body.age,]]))
    t3 =  np.array([ [scaled[0, 0]] ])
    t4 =  np.array([[1 if body.gender == "M" else 0]])
    t5 =  np.array([[1 if body.gender == "F" else 0]])
    t6 =  np.array([[1 if body.embarked == "C" else 0]])
    t7 =  np.array([[1 if body.embarked == "Q" else 0]])
    t8 =  np.array([[1 if body.embarked == "S" else 0]])

    result = np.concatenate((t0, t1, t2, t3, t4, t5, t6, t7, t8, ), axis=1).astype(np.float32)

    input_name = model.get_inputs()[0].name
    label_name = model.get_outputs()[0].name
    prediction = model.run([label_name], {input_name: result})[0]

    return Output(prediction=prediction)
