function login() {
  var xhr = new XMLHttpRequest();
  var url = '/api/login';

  var sessionDuration = 24*60*60;

  if ($('#form_remember').prop('checked')) {
    sessionDuration = 31 * 24*60*60;
  }
  console.log($('#form_remember').attr('checked'));
  console.log($('#form_remember'));

  var unser = {
    'Login': $('#form_login').val(),
    'Password': $('#form_pwd').val(),
    'Duration': sessionDuration
  };

  var req = JSON.stringify(unser);

  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4) {
      var res = JSON.parse(xhr.responseText)
      if (res.status == 'OK') {
        $('#result').removeClass('FAIL').addClass('OK');
        createCookie('SID', res.sid, sessionDuration);
        
        window.setTimeout(function() {
          window.location.replace(res.redirect);
        }, 300);
      } else {
        $('#result').removeClass('OK').addClass('FAIL');
      }

      $('#result').text(res.reason);
    }
  };

  xhr.open('POST', url, true);
  xhr.send(req)
}