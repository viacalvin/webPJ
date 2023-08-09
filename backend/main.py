from fastapi import FastAPI,Depends,status
from fastapi.responses import RedirectResponse,HTMLResponse
from fastapi.security import OAuth2PasswordRequestForm
from fastapi_login import LoginManager 
from fastapi_login.exceptions import InvalidCredentialsException
from fastapi.templating import Jinja2Templates
import path
import uuid
import uvicorn
from fastapi import Request
import hashlib

# sha256 암호화와 간단한 수준의 로그인 API

# result = hashlib.sha256(put.encode())
# passwd += result.hexdigest()
# result = hashlib.sha256(passwd.encode())
app= FastAPI()

SECRET = "secret-key"
# pth = path.dirname(__file__)
templates = Jinja2Templates(directory="templates")

# manager = LoginManager(SECRET,tokenUrl="/auth/login",use_cookie=True)
# manager.cookie_name = "cookie"

db_password = "key"
result = hashlib.sha256(db_password.encode())
en_db_password = result.hexdigest()
# en_db_password = uuid.uuid4(db_password)
DB = {"username":{"password":f"{en_db_password}"}} 

# @manager.user_loader
# def load_user(username:str):
#     user = DB.get(username)
#     return user

@app.post("/auth/login")
def login(data: OAuth2PasswordRequestForm = Depends()):
    username = data.username
    password = data.password
    res = hashlib.sha256(password.encode())
    password = res.hexdigest()
    user = load_user(username)
    if not user:
        raise InvalidCredentialsException
    elif password != user['password']:
        raise InvalidCredentialsException
    access_token = manager.create_access_token(
        data={"sub":username}
    )
    resp = RedirectResponse(url="/private",status_code=status.HTTP_302_FOUND)
    # manager.set_cookie(resp,access_token)
    return resp

@app.get("/private")
def getPrivateendpoint():
    return "Success"

# @app.get("/",response_class=HTMLResponse)
# def loginwithCreds(request:Request):
#     with open(path.join(pth, "templates/login.html")) as f:
#         return HTMLResponse(content=f.read())

@app.get("/")
async def home(request: Request):
    return templates.TemplateResponse("index.html",{"request":request})

if __name__ == "__main__":
    uvicorn.run(app="main:app",host="127.0.0.1", port=8080)