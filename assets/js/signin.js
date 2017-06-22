/**
 * Created by igor on 16.06.17.
 */
$(document).ready(function () {
    console.log("ready!");
    var t = getSessionToken();
    if (t != undefined) {
        console.log(getSessionUser());
        document.getElementById("sign-in-a").style.display = "none";
        document.getElementById("reg-a").style.display = "none";
        document.getElementById("sign-out-a").style.display = "block";
        document.getElementById("user-a").style.display = "block";

        document.getElementById("user-a").innerHTML = getSessionUser()
    }

});

function signin() {
    var login = document.getElementById('inputLogin').value
    var pass = document.getElementById('inputPassword').value
    var obj = new Object();
    obj.login = login;
    obj.pass = pass;
    $.ajax
    ({
        type: "POST",
        url: "/getjwt",
        dataType: 'json',
        async: false,

        data: {
            login: login,
            pass: pass

        },
        success: function (token) {
            $.session.set("token", token.token);
            $.session.set("user", token.user);
            window.location.replace("/")
        },

        error: function (e) {
            document.getElementById('errorDiv').innerHTML = '<span class="label label-danger "> Wrong login or password. Try again or go away</span>'
        }
    });

}

function getSessionToken() {
    return $.session.get("token");
}
function getSessionUser() {
    return $.session.get("user");
}