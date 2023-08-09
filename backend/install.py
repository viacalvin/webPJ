import os
# fastapi 설치

# main.py를 실행하기 전에 실행해주세요
package_list = ['fastapi', 'fastapi_login', 'path', 'jinja2', 'python-multipart', 'uvicorn']

for i in package_list:
    print("pacakge : ", i, "설치중")
    os.system(f"pip install {i}")