'use strict';

var express = require('express');
var template = require('./DApp2_template');
var registerUser = require('./registerUser');
var bodyParser = require('body-parser');
var app = express();

const util = require('util');

app.use(bodyParser.urlencoded({extended:true}))
app.use(bodyParser.json())

let printerList = ['init'];
const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const { list_q, list_tx } = require('./DApp2_template');

const ccpPath = path.resolve(__dirname, '..', 'basic-network', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

async function query (func, name) {
    try {

        // 지갑에서 신원 선택
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // 등록된 사용자인지 확인
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // 게이트웨이에 연결
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        // 네트워크에 접속
        const network = await gateway.getNetwork('mychannel');

        // 스마트 컨트랙트 요청
        const contract = network.getContract('cl_tcc7');

        let result;
        // 트랜잭션 Submit
        result =await contract.evaluateTransaction(func, name);
                
        // 프로세스 응답
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return result;

    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    }
};


app.get('/', async function(request, response) { 
    let queryjson = await query('query', "2523648240000001ba344d8000000007ff9f800000000010a10000000000000d");
    console.log(queryjson);
    let obj = JSON.parse(queryjson);
    var list_tx = template.list_tx(obj)
    var list_ch = template.list_ch(obj)
    var list_q = template.list_q(obj)
    var list_x = template.list_x(obj)
    var list_y = template.list_y(obj)
    var list_z = template.list_z(obj)
    var list_ts = template.list_ts(obj)      
    var list = `<table class="table table-hover">

    <thead>    
    <tr>
        <th scope="col">TxID (Tracking Number)</th>            
        ${list_tx}         
    </tr>
    <tr>
        <th scope="col">channel_ID (Channel Name)</th>
        ${list_ch}
    </tr>
    <tr>
        <th scope="col">q (Public Value)</th>
        ${list_q}
    </tr>
    <tr>
        <th scope="col">X (Public Value)</th>
        ${list_x}
    </tr>
    <tr>
        <th scope="col">Y (Public Value)</th>
        ${list_y}
    </tr>
    <tr>
        <th scope="col">Z (Public Value)</th>
        ${list_z}
    </tr>
    <tr>
        <th scope="col">TimeStamp</th>
        ${list_ts}
    </tr>
    </thead>

    <tbody>`;    

      list = list + `  </tbody>
    </table>`

      var html = template.HTML(list);

      response.send(html);
});

app.get('/TxID', async function(request, response) {

  var list_tx = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "2523648240000001ba344d8000000007ff9f800000000010a10000000000000d");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_tx = list_tx + template.list_tx(obj);
  list_tx = list_tx + `  </tbody>
  </table>`

    var html = template.HTML(list_tx);
    response.send(html);
});

app.get('/chID', async function(request, response) {

  var list_ch = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "2523648240000001ba344d8000000007ff9f800000000010a10000000000000d");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_ch = list_ch + template.list_ch(obj);
  list_ch = list_ch + `  </tbody>
  </table>`

    var html = template.HTML(list_ch);
    response.send(html);
});

app.get('/q', async function(request, response) {

  var list_q = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "2523648240000001ba344d8000000007ff9f800000000010a10000000000000d");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_q = list_q + template.list_q(obj);
  list_q = list_q + `  </tbody>
  </table>`

    var html = template.HTML(list_q);
    response.send(html);
});

app.get('/X', async function(request, response) {

  var list_x = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "2523648240000001ba344d8000000007ff9f800000000010a10000000000000d");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_x = list_x + template.list_x(obj);
  list_x = list_x + `  </tbody>
  </table>`

    var html = template.HTML(list_x);
    response.send(html);
});

app.get('/Y', async function(request, response) {

  var list_y = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "2523648240000001ba344d8000000007ff9f800000000010a10000000000000d");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_y = list_y + template.list_y(obj);
  list_y = list_y + `  </tbody>
  </table>`

    var html = template.HTML(list_y);
    response.send(html);
});

app.get('/Z', async function(request, response) {

  var list_z = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "2523648240000001ba344d8000000007ff9f800000000010a10000000000000d");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_z = list_z + template.list_z(obj);
  list_z = list_z + `  </tbody>
  </table>`

    var html = template.HTML(list_z);
    response.send(html);
});

app.get('/TimeStamp', async function(request, response) {

  var list_ts = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "2523648240000001ba344d8000000007ff9f800000000010a10000000000000d");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_ts = list_ts + template.list_ts(obj);
  list_ts = list_ts + `  </tbody>
  </table>`

    var html = template.HTML(list_ts);
    response.send(html);
});

app.get('/verify', function(request, response) {
  var html = template.HTML(`
              <form action="http://34.64.221.107:3000/verify_process" method="post">
                <p><input type="text" name="Token" placeholder="q (Public Value)"></p>          
                <p>
                  <input type="submit" class="btn btn-primary"/>
                  <button type="button" class="btn" onClick="location.href='/'">취소</button>
                </p>
              </form>
            `);
    response.send(html);
});

app.get('/login', function(request, response) {
    var html = template.HTML(`
                <form action="http://34.64.221.107:3000/login_process" method="post">
                  <p><input type="text" name="ID" placeholder="ID"></p>
                  <p><input type="text" name="PW" placeholder="PW"></p>
                  <p>
                    <input type="submit" class="btn btn-primary"/>
                    <button type="button" class="btn" onClick="location.href='/'">취소</button>
                  </p>
                </form>
              `);
      response.send(html);
});


app.post('/verify_process', function(request, response) {
    var token = request.body.Token;
    printerList.push(token);
    
    query('verify', token);

    console.log(request.body)
    console.log(token)    
    response.redirect('/')
});

app.post('/login_process', function(request, response) {
    var id = request.body.ID;
    var pw = request.body.PW;
    
    console.log(request.body)
    console.log(id)
    registerUser.registerUser(id, pw);
    response.redirect('/')
});

app.get('/test', function(request, response) {
    response.send('test');
});

app.listen(3000, function() {
    console.log('Example app listening on port 3000!');
})

