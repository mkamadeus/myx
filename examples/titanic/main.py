from fastapi import FastAPI
from pydantic import BaseModel
import joblib
import onnxruntime as rt
import numpy as np

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
    t0 =  np.array([[body.pclass]], dtype=np.int_)
    t1 =  np.array([[body.sibsp]], dtype=np.int_)
    t2 =  np.array([[body.parch]], dtype=np.int_)
    scaler = joblib.load('age.scaler')
    t3 =  scaler.transform(np.array([[body.age]], dtype=np.float32))
    t4 =  np.array([[1 if body.gender == "M" else 0]], dtype=np.int_)
    t5 =  np.array([[1 if body.gender == "F" else 0]], dtype=np.int_)
    t6 =  np.array([[1 if body.embarked == "C" else 0]], dtype=np.int_)
    t7 =  np.array([[1 if body.embarked == "Q" else 0]], dtype=np.int_)
    t8 =  np.array([[1 if body.embarked == "S" else 0]], dtype=np.int_)

    prediction = np.concatenate((t0, t1, t2, t3, t4, t5, t6, t7, t8, ), axis=1).astype(np.float32)

    input_name = model.get_inputs()[0].name
    label_name = model.get_outputs()[0].name
    prediction = model.run([label_name], {input_name: prediction})[0]

    return Output(prediction=prediction)
