from flask import Flask , render_template , request
import requests

app = Flask(__name__)

api_url = "labapi-service"

@app.route('/')
def Index():
    return render_template("index.html")

@app.route('/getall')
def GetAll():
    res = requests.get(f"http://{api_url}/users").json()
    return res

@app.route('/addpage')
def AddPage():
    return render_template("add.html")

@app.route('/add' , methods=['POST'])
def Add():
    
    req = request.get_json()
    
    header = {
        
    }
    
    body = {
        "name":req['name'],
        "detail":req['detail'],
        "role":req['role']
    }
    
    res = requests.post(f"http://{api_url}/user/add" , headers=header , json=body).text
    return res

if __name__ == "__main__":
    app.run(host="0.0.0.0" , port=5000 , debug=True)
    
    