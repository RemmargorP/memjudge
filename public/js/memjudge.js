function checkSignUpForm() {
  var login = $("#reg_login").val();
  console.log(login);
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

function validateEmail(email) {
    var re = /\S+@\S+\.\S+/;
    return re.test(email);
}