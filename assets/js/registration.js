/**
 * Created by igor on 16.06.17.
 */
function registration() {

    var login = document.getElementById("inputLogin").value;
    var username = document.getElementById("username").value;
    var password = document.getElementById("inputPassword").value;
    var email = document.getElementById("inputEmail").value;
    var passsword2 = document.getElementById("inputPassword2").value;
    var patt = /^[A-Za-z0-9_-]{3,50}$/;
    var patt2 = /^[a-zA-z_]([A-Za-z0-9_.-]{0,100})@([a-z]{2,8}[.][a-z]{2,8})$/;

    if (patt.test(login) && patt2.test(email) && patt.test(username) && patt.test(password)) {
        var obj = new Object()
        obj.name = username;
        obj.login = login;
        obj.email = email;
        obj.pass = password;
        if (password == passsword2) {
            $.post("/adduser", {
                    data: JSON.stringify(obj)
                }
            )
            document.getElementById('wrongDateDiv1').innerHTML = ''
        }
        else {
            document.getElementById('wrongDateDiv1').innerHTML = ' <span class="label label-danger "> Password does not match </span>'
        }
    } else {
        document.getElementById('wrongDateDiv1').innerHTML = ' <span class="label label-danger "> incorrect field </span>'
    }

}

