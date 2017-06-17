/**
 * Created by igor on 16.06.17.
 */
var tests;
$(document).ready(function () {
    console.log("ready!");
    var t = getSessionToken();
    if (t != undefined) {
        console.log(getSessionUser())
        document.getElementById("sign-in-a").style.display = "none";
        document.getElementById("reg-a").style.display = "none";
        document.getElementById("sign-out-a").style.display = "block";
        document.getElementById("user-a").style.display = "block";
        document.getElementById("user-a").innerHTML = getSessionUser()
    }
    getTest();
    insertTest();


});

function getSessionToken() {
    return $.session.get("token");
}
function getSessionUser() {
    return $.session.get("user");
}
function getTest() {
    $.ajax
    ({
        type: "GET",
        url: "/getvotes",
        dataType: 'json',
        async: false,

        success: function (data) {

            tests = data


        }


    });
}
function insertTest() {
    var result = '';
    for (var i in tests.Vote) {
        var item = tests.Vote[i];
        console.log(item.name)
        result += ' <li><a  onclick="getTestById(' + item.ID + ')" id="' + item.ID + '">' + item.name + '</a></li>'
    }

    document.getElementById('ulTest').innerHTML = result;
}
var test;
function getTestById(id) {

    $.ajax
    ({
        type: "POST",
        url: "/getvote",
        dataType: 'json',
        async: false,
        data: {
            idvote: id
        },

        success: function (data) {

            test = data;
        }

    });
    console.dir(test)
    document.getElementById("contentDashBoard").style.display = "block";
    var t = getSessionToken();
    if (t != undefined) {
        document.getElementById("startTestButton").style.display = "block";

    }
}

function getUsersToVote(id) {
    var users
    $.ajax
    ({
        type: "POST",
        url: "/getuserstovote",
        dataType: 'json',
        async: false,
        data: {
            idvote: id
        },

        success: function (data) {

            users = data;
        }

    });

    console.dir(users)
    var result = ''
    for (var i in users) {
        result += '<tr> <td>' + i + '</td> <td>' + users[i].Name + '</td> <td>' + users[i].Time + '</td> </tr>'

    }
    document.getElementById('tbodyUserToAnswer').innerHTML = result;
}

function createTest() {
    var result = '';
    result += '<label > Test name is ' + test.Name + ' </label><br> <label > Test category is  ' + test.Category + '</label><br>'
    for (var i = 0; i < test.Questions.Question.length; i++) {
        result += '<label > Question № ' +Number(i+1) + ': ' + test.Questions.Question[i] + ' </label> '
        var type = test.Questions.Type[i]
        var answers = test.AnswerToQuestion[i];
        for (var j = 0; j < test.AnswerToQuestion[i].length; j++) {
            result += '<label class="radio"><input type="' + type + '" name="optradio' + i + '" checked>' + answers[j] + ' </label>'

        }
    }


    document.getElementById('createTestDiv').innerHTML = result;
}