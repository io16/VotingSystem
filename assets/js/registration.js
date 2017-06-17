/**
 * Created by igor on 16.06.17.
 */
function registration() {

    var login = document.getElementById("inputLogin").value
    var username = document.getElementById("username").value
    var password = document.getElementById("inputPassword").value
    var email = document.getElementById("inputEmail").value
    var passsword2 = document.getElementById("inputPassword2").value
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

    }
}
