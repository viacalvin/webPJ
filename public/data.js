// data.js / useStates 데이터 저장 용도

var Load_db = ["ID"+'test','PW'+'Password']
const mysql = require('mysql')
let jsondata = require('./js/loader.js')
let url = 'dns.dns.com'
let address = './page/addr/api/login'
let usercode = "c1982992"
// usercode - 테스트 코드, 추후 개선 예정
let enckeyword = "e1992"
// 서버 private Key Encrypt

const timer = new(Date())
let token_public_uuid = '56591bbf-38c0-47b3-a957-adc4d7042802'

function first_load(jsondata){
    if(cookie == null){    
        $.ajax({
            url: "http://api"+url+address,
            method: "POST",        
            data: jsondata,
            contentType: "json",
            success: function(data){
               let token = token_generator(add)(
                token_name = usercode,
                token_encrypt = enckeyword,
                token_time = "time.relay = " + 1200 +'s',
                token_validation = token_public_uuid
               )
            },
            error: function(data) {
                //your error function
            }
        });
    }
}

function app(main){
    let aps = this.main
    console.log("host.reference, querry " + aps)
    let Load_db = dbs()
}

function sql_connect(){
    module.exports = function () {
        return {
          init: function () {
            return mysql.createConnection({
              host: 'localhost',
              port: '3306',
              user: 'root',
              password: 'root',
              database: 'DBS'
            })
          },
          
          db_open: function (con) {
            con.connect(function (err) {
              if (err) {
                console.error('mysql connection error :' + err);
              } else {
                console.info('mysql is connected successfully.');
              }
            })
          }
        }
      };
}

function server_return(cookie) {
    COOKIE = NEW()
}

function uuid_list(uuid){
    // 테스트용 UUID 저장소
    const keyset = require('./js/keyset.js')
    let keylist = keyset
    alert("uuid keys = " + keylist)
}