
'use strict';

var express = require('express');
var template = require('./DApp_template');
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
const { list_token, list_proof } = require('./report_template');

const ccpPath = path.resolve(__dirname, '..', 'basic-network', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

const total_tx = 1;

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
        const contract = network.getContract('tcc8');

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
    const dt = new Date(); // submit time
    console.log('1');
    console.log(dt);
    let queryjson = await query('query', "vh4s4qndc4mj7gla");
    console.log(queryjson);
    let obj = JSON.parse(queryjson);
    var list_tx = template.list_tx(obj)
    var list_ch = template.list_ch(obj)
    var list_token = template.list_token(obj)
    var list_proof = template.list_proof(obj)
    var list_pk = template.list_pk(obj)
    var list_ts = template.list_ts(obj)      
    var list = `<table class="table table-hover">

    <thead>    
    
    <tr>
        <th scope="col">Read Latency</th>`;
        console.log('3');
        
        `<th scope="col">Read Throughput</th>`;
        console.log('4');
        
    `</tr>
    </thead>

    <tbody>`;    

      list = list + `  </tbody>
    </table>
        <script>`;
            var dt2 = new Date(); // Time when response received
            console.log('2');
            console.log(dt2);
            var read_latency = dt2.getTime() - dt.getTime(); // Read Latency
            read_latency = read_latency / 1000;
            var read_through = total_tx / read_latency; // Read Throughput

            console.log('4');
            console.log(read_latency);       // s
            //document.write(read_latency);            
            console.log('5');
            console.log(read_through);       // tps
            console.log('6');
            //document.write(read_through);             
        `</script>`

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

  let queryjson = await query('query', "vh4s4qndc4mj7gla");
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

  let queryjson = await query('query', "vh4s4qndc4mj7gla");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_ch = list_ch + template.list_ch(obj);
  list_ch = list_ch + `  </tbody>
  </table>`

    var html = template.HTML(list_ch);
    response.send(html);
});

app.get('/Token', async function(request, response) {

  var list_token = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "vh4s4qndc4mj7gla");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_token = list_token + template.list_token(obj);
  list_token = list_token + `  </tbody>
  </table>`

    var html = template.HTML(list_token);
    response.send(html);
});

app.get('/Hased_Proof', async function(request, response) {

  var list_proof = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "vh4s4qndc4mj7gla");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_proof = list_proof + template.list_proof(obj);
  list_proof = list_proof + `  </tbody>
  </table>`

    var html = template.HTML(list_proof);
    response.send(html);
});

app.get('/PublicInputs', async function(request, response) {

  var list_pk = `<table class="table table-hover">
  <thead>
  <tr>
      <th scope="col">Value</th>        
  </tr>
  </thead>
  <tbody>`;

  let queryjson = await query('query', "vh4s4qndc4mj7gla");
  console.log(queryjson);
  let obj = JSON.parse(queryjson);
  list_pk = list_pk + template.list_pk(obj);
  list_pk = list_pk + `  </tbody>
  </table>`

    var html = template.HTML(list_pk);
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

  let queryjson = await query('query', "vh4s4qndc4mj7gla");
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
                <p><input type="text" name="Token" placeholder="Token"></p>          
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

