from fastapi import FastAPI
from pydantic import BaseModel
import joblib
import onnxruntime as rt
import numpy as np

app = FastAPI()
model = rt.InferenceSession("titanic.onnx")

class Input(BaseModel):
	credit_score : float
	geography : str
	gender : str
	tenure : float
	balance : float
	num_of_products : int
	has_credit_card : int
	is_active_member : int
	estimated_salary : float
	

class Output(BaseModel):
	prediction : float
	


@app.post('/')
async def root(body: Input):
    t4 =  np.array([[1 if body.gender == "M" else 0]], dtype=np.int_)
    t5 =  np.array([[1 if body.gender == "F" else 0]], dtype=np.int_)

    prediction = np.concatenate((t0, t1, ), axis=1).astype(np.float32)

    input_name = model.get_inputs()[0].name
    label_name = model.get_outputs()[0].name
    prediction = model.run([label_name], {input_name: prediction})[0]

    return Output(prediction=prediction)
