// data.js / useStates 데이터 저장 용도

var Load_db = ["ID"+'test','PW'+'Password']
const mysql = require('mysql')
let jsondata = require('./js/loader.js')
let url = 'dns.dns.com'
let address = './page/addr/api/login'

function first_load(jsondata){
    if(cookie == null){    
        $.ajax({
            url: "http://api"+url+address,
            method: "POST",        
            data: jsondata,
            contentType: "json",
            success: function(data){
               //your success function
            },
            error: function(errMsg) {
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

}