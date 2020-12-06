module.exports = {
    HTML:function(list) {
      return `
      <!DOCTYPE html>
      <html lang="en">
      <head>
        <title>CL-Signature</title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js"></script>
      </head>
      <body>
  
      <div class="jumbotron text-center" style="margin-bottom:0">
        <h1>ZKP</h1>
      </div>
  
      <nav class="navbar navbar-expand-sm bg-secondary navbar-dark">
        <a class="navbar-brand" href="#">익명인증서</a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#collapsibleNavbar">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="collapsibleNavbar">
          <ul class="navbar-nav">
            <li class="nav-item">
              <a class="nav-link" href="/verify">검증</a>
            </li>            
          </ul>
        </div>
        <div class="collapse navbar-collapse" id="collapsibleNavbar">
          <ul class="navbar-nav">
            <li class="nav-item">
              <a class="nav-link" href="/login">로그인</a>
            </li>            
          </ul>
        </div>
      </nav>
  
      <div class="container" style="margin-top:30px">
        <div class="row">
          <div class="col-sm-4">
            <ul class="nav nav-pills flex-column">
              <li class="nav-item">
                <a class="nav-link active" href="/">CL-Signature</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/TxID">TxID (Tracking Number)</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/chID">channel_ID (Channel Name)</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/q">q (Public Value)</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/X">X (Public Value)</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/Y">Y (Public Value)</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/Z">Z (Public Value)</a>
              </li>              
              <li class="nav-item">
                <a class="nav-link" href="/TimeStamp">Timestamp</a>
              </li>
            </ul>
            <hr class="d-sm-none">
          </div>
          <div class="col-sm-8">
          ${list}
          </div>
          <script src="jquery.js"></script>
          <script>
            $(document).ready(function () {
              $('.nav-link').click(function () {
                $('.nav-item').removeClass("active");
                $(this).addClass("active");
              });
            });      

          </script> 
        </body>
      </html>
      `;
    },list:function(queryObj){
      var list = ``;
      list = list + `<tr>`;
      list = list + `<td>${queryObj.TxID}</td>`
      list = list + `<td>${queryObj.ChID}</td>`
      list = list + `<td>${queryObj.q}</td>`
      list = list + `<td>${queryObj.X}</td>`
      list = list + `<td>${queryObj.Y}</td>`
      list = list + `<td>${queryObj.Z}</td>`
      list = list + `<td>${queryObj.Timestamp}</td>`
      list = list + `</tr>`;
      return list;
    },list_tx:function(queryObj){
      var list_tx = ``;
      list_tx = list_tx + `<tr>`;
      list_tx = list_tx + `<td>${queryObj.TxID}</td>`      
      list_tx = list_tx + `</tr>`;
      return list_tx;
    },list_ch:function(queryObj){
      var list_ch = ``;
      list_ch = list_ch + `<tr>`;
      list_ch = list_ch + `<td>${queryObj.ChID}</td>`      
      list_ch = list_ch + `</tr>`;
      return list_ch;
    },list_q:function(queryObj){
      var list_q = ``;
      list_q = list_q + `<tr>`;
      list_q = list_q + `<td>${queryObj.q}</td>`      
      list_q = list_q + `</tr>`;
      return list_q;
    },list_x:function(queryObj){
      var list_x = ``;
      list_x = list_x + `<tr>`;
      list_x = list_x + `<td>${queryObj.X}</td>`      
      list_x = list_x + `</tr>`;
      return list_x;
    },list_y:function(queryObj){
      var list_y = ``;
      list_y = list_y + `<tr>`;
      list_y = list_y + `<td>${queryObj.Y}</td>`      
      list_y = list_y + `</tr>`;
      return list_y;
    },list_z:function(queryObj){
      var list_z = ``;
      list_z = list_z + `<tr>`;
      list_z = list_z + `<td>${queryObj.Z}</td>`      
      list_z = list_z + `</tr>`;
      return list_z;
    },list_ts:function(queryObj){
      var list_ts = ``;
      list_ts = list_ts + `<tr>`;
      list_ts = list_ts + `<td>${queryObj.Timestamp}</td>`      
      list_ts = list_ts + `</tr>`;
      return list_ts;
    }
  }

  