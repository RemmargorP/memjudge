// Sign Up page funcs

function checkSignUpForm() {
  var login = $("#reg_login").val();
  var ok = !(login.includes(" ") || login.length < 3 || login.length > 64), total = true;
  total = total & ok;

  if (ok) {
    $("#reg_login_req").slideUp(100);
  } else {
    $("#reg_login_req").slideDown(100);
  }

  var email = $("#reg_email").val();
  ok = validateEmail(email);
  total = total & ok;

  if (ok) {
    $("#reg_email_req").slideUp(100);
  } else {
    $("#reg_email_req").slideDown(100);
  }

  var pwd = $("#reg_pwd").val();
  ok = pwd.length >= 6 && pwd.length <= 64;
  total = total & ok;

  if (ok) {
    $("#reg_pwd_req").slideUp(100);
  } else {
    $("#reg_pwd_req").slideDown(100);
  }

  return total;
}

function signup() {
  if (!checkSignUpForm()) return;
  
  var xhr = new XMLHttpRequest();
  var url = "/api/signup";

  var req = JSON.stringify({"Login": $("#reg_login").val(), "Email": $("#reg_email").val(), "Password": $("#reg_pwd").val()});

  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4) {
      $("#result").text(xhr.responseText); 
      var tmp = JSON.parse(xhr.responseText);
    }
  };

  xhr.open("POST", url, true);
  xhr.send(req)
}

// Utils

function validateEmail(email) {
    var re = /\S+@\S+\.\S+/;
    return re.test(email);
}