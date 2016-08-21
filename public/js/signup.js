// Sign Up page funcs

function checkSignUpForm() {
  var login = $('#form_login').val();
  var ok = !(login.includes(' ') || login.length < 3 || login.length > 64), total = true;
  total = total & ok;

  if (ok) {
    $('#form_login_req').slideUp(100);
  } else {
    $('#form_login_req').slideDown(100);
  }

  var email = $('#form_email').val();
  ok = validateEmail(email);
  total = total & ok;

  if (ok) {
    $('#form_email_req').slideUp(100);
  } else {
    $('#form_email_req').slideDown(100);
  }

  var pwd = $('#form_pwd').val();
  ok = pwd.length >= 6 && pwd.length <= 64;
  total = total & ok;

  if (ok) {
    $('#form_pwd_req').slideUp(100);
  } else {
    $('#form_pwd_req').slideDown(100);
  }

  return total;
}

function signup() {
  if (!checkSignUpForm()) return;
  
  var xhr = new XMLHttpRequest();
  var url = '/api/signup';

  var req = JSON.stringify({'Login': $('#form_login').val(), 'Email': $('#form_email').val(), 'Password': $('#form_pwd').val()});

  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4) {
      var res = JSON.parse(xhr.responseText)
      if (res.status == 'OK') {
        $('#result').removeClass('FAIL').addClass('OK');

        // Success, now login
        login();

        window.setTimeout(function() {
          window.location.replace(res.redirect)
        }, 200);
      } else {
        $('#result').removeClass('OK').addClass('FAIL');
      }

      $('#result').text(res.reason);
    }
  };

  xhr.open('POST', url, true);
  xhr.send(req)
}

// Utils

function validateEmail(email) {
    var re = /\S+@\S+\.\S+/;
    return re.test(email);
}