/**
 * Created by igor on 16.06.17.
 */
$(document).ready(function () {
    console.log("ready!");
    var t = getSessionToken();
    if (t!= undefined){
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
    // $.post("/getjwt", {
    //     data: JSON.stringify(obj)
    // }, function (data, status, xhr) {
    //     alert(xhr.status);
    //     alert(status);
    //     alert(data);
    // });

    $.ajax
    ({
        type: "POST",
        url: "/getjwt",
        dataType: 'json',
        async: false,
        // headers: {
        //     "Authorization": "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9uIFNub3ciLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNDkwMzYxMzg4fQ.AUp-oA5HXLojnrsnrmTHbWlZBduJs69osZEVh3ZfBfw"
        // },
        data: {
            login: login,
            pass: pass

        },
        success: function (token) {

            $(function () {
                $.session.set("token", token.token);
                $.session.set("user", token.user);
            });

                window.location.reload("http://localhost:1323/")


        }

    });

}

function getSessionToken() {
    return $.session.get("token");
}function getSessionUser() {
    return $.session.get("user");
}