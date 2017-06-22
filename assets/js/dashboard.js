/**
 * Created by igor on 16.06.17.
 */
var tests; // list of tests
var stats; // list of stats to test
$(document).ready(function () {
    console.log("ready!");
    var t = getSessionToken();
    if (t != undefined) {
        console.log(getSessionUser())
        document.getElementById("sign-in-a").style.display = "none";
        document.getElementById("reg-a").style.display = "none";
        document.getElementById("sign-out-a").style.display = "block";
        document.getElementById("user-a").style.display = "block";
        document.getElementById("create-test-a").style.display = "block";
        document.getElementById("user-a").innerHTML = getSessionUser()
    }
    getTest();
    // /
    // insertTest();
    insertCategories();
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

function insertCategories() {
    var result = '';
    var categoryList = [];
    var categoryListCount = [];
    document.getElementById("listNamePre").innerHTML="     List of Categories";
    for (var i in tests.Vote) {
        var item = tests.Vote[i];
        var statusCategory = true;
        for (var j = 0; j < categoryList.length; j++) {
            if (categoryList[j] == item.category) {
                categoryListCount[j] += 1;
                statusCategory = false;
                break;
            }
        }
        if (statusCategory) {
            categoryListCount.push(1);
            categoryList.push(item.category)
        }
    }
    for (var i = 0; i < categoryList.length; i++) {

        var category =    categoryList[i];
        result +='<li> <a onclick=insertTest("'+category+'")> ' + categoryList[i] + '  <span class = "badge" >' + categoryListCount[i] + ' </span > </a ></li> ';

    }
    document.getElementById('ulTest').innerHTML = result;

}

function insertTest(category) {
    document.getElementById("listNamePre").innerHTML="     List of Tests";
    var result = '<li><a  onclick="insertCategories()">Back</a></li>';
    for (var i in tests.Vote) {

        var item = tests.Vote[i];
        if (category == item.category)
            result += ' <li><a  onclick="getTestById(' + item.ID + ')" id="' + item.ID + '">' + item.name + '</a></li>'
    }
    document.getElementById('ulTest').innerHTML = result;
}

function isValidSession() {
    $.ajax
    ({
        type: "POST",
        url: "/restricted/test",
        dataType: 'json',
        async: false,
        headers: {
            "Authorization": "Bearer " + getSessionToken()
        },
        success: function (data) {
            console.log("valid")
        },
        error: function (e) {
            console.log(e)
            $.session.set("user", "");
            $.session.set("token", "");
            window.location.replace("/signin")
        }
    });
}
var test; // current test
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
            test.idTest = id
        }

    });

    document.getElementById("contentDashBoard").style.display = "block";
    document.getElementById("contentHeader").innerHTML = 'Test Name : ' + test.Name + '<br>  Test Category: ' + test.Category;
    getStatsToTest();
    setQuestionsIntoSelect();
    getChartToQuestion(0);
    document.getElementById("tbodyUserToAnswer").style.display = "none";
    document.getElementById("createTestDiv").style.display = "none";
    document.getElementById("saveTestButton").style.display = "none";
    var t = getSessionToken();
    if (t != undefined) {
        isUserCompleteTest()

        document.getElementById("showUsersToTestButton").style.display = "block";
    }
}

function isUserCompleteTest() {
    $.ajax
    ({
        type: "POST",
        url: "/isUserCompleteTest",
        dataType: 'json',
        async: false,
        data: {
            idvote: test.idTest,
            login: getSessionUser()
        },
        success: function (data) {
            var t = getSessionToken();
            if (t != undefined) {
                if (data) {

                    document.getElementById("startTestButton").style.display = "none";
                } else {
                    document.getElementById("startTestButton").style.display = "block";
                }
            }
            return data
        }

    });
}
function getStatsToTest() {
    $.ajax
    ({
        type: "POST",
        url: "/getvotestat",
        dataType: 'json',
        async: false,
        data: {
            voteid: test.idTest
        },
        success: function (data) {
            stats = data;
        }
    });

}
function getUsersToTest() {
    var users;
    document.getElementById("createTestDiv").style.display = "none";
    document.getElementById("saveTestButton").style.display = "none";

    $.ajax
    ({
        type: "POST",
        url: "/getuserstovote",
        dataType: 'json',
        async: false,
        data: {
            idvote: test.idTest
        },
        success: function (data) {

            users = data;
            console.dir(users);
            var result = '<thead><tr><th>#</th> <th>User Name</th> <th>Time</th> </tr> </thead> <tbody >';
            for (var i in users) {
                result += '<tr> <td>' + i + '</td> <td>' + users[i].Name + '</td> <td>' + users[i].Time + '</td> </tr>'
            }
            result += ' </tbody>';
            document.getElementById('tbodyUserToAnswer').innerHTML = result;
            document.getElementById("tbodyUserToAnswer").style.display = "block";
        }
    });
}

function createTest() {
    document.getElementById("tbodyUserToAnswer").style.display = "none";
    document.getElementById("createTestDiv").style.display = "block";
    document.getElementById("saveTestButton").style.display = "block";
    isValidSession();
    var result = '';
    result += '<label > Test name is ' + test.Name + ' </label><br> <label > Test category is  ' + test.Category + '</label><br>'
    for (var i = 0; i < test.Questions.Question.length; i++) {
        result += '<label > Question № ' + Number(i + 1) + ': ' + test.Questions.Question[i] + ' </label> '
        var type = test.Questions.Type[i]
        var answers = test.AnswerToQuestion[i];
        for (var j = 0; j < test.AnswerToQuestion[i].length; j++) {
            result += '<label class="radio"><input type="' + type + '" name="optradio' + i + '" checked>' + answers[j] + ' </label>'
        }
    }
    document.getElementById('createTestDiv').innerHTML = result;
    document.getElementById("saveTestButton").style.display = "block";
}

function saveTest() {
    var arrayOfRadios = [];

    for (var i = 0; i < test.Questions.Question.length; i++) {
        var radioToQuestion = [];
        for (var j = 0; j < test.AnswerToQuestion[i].length; j++) {

            if (document.getElementsByName('optradio' + i)[j].checked) {
                radioToQuestion.push("1");
            } else radioToQuestion.push("0");
        }
        arrayOfRadios.push(radioToQuestion)
    }

    var obj = new Object()
    obj.vote = [];

    for (var i in arrayOfRadios) {

        var answer = new Object()
        answer.answer = []
        answer.answer = arrayOfRadios[i]

        obj.vote.push(answer)
    }
    obj.voteid = test.idTest

    $.ajax
    ({
        type: "POST",
        url: "/restricted/saveuservote",
        dataType: 'json',
        async: false,
        headers: {
            "Authorization": "Bearer " + getSessionToken()
        },
        data: {
            data: JSON.stringify(obj)
        },
        success: function (data) {
            getTestById(test.idTest)
        },
        error: function () {
            $.session.set("user", "");
            $.session.set("token", "");
            window.location.replace("/signin")
        }
    });
}

function setQuestionsIntoSelect() {

    var select = document.getElementById("sel1");
    var result = '';
    for (var i = 0; i < test.Questions.Question.length; i++) {
        result += '  <option value="' + i + '">' + test.Questions.Question[i] + '</option>';
    }
    select.innerHTML = result;

}
function onChangeSelect() {

    var e = document.getElementById("sel1");
    var questionNumber = e.options[e.selectedIndex].value;
    getChartToQuestion(questionNumber)
}
function getChartToQuestion(questionNumber) {
    var answers = test.AnswerToQuestion[questionNumber];

    var data = [[]];
    data[0][0] = "Count of answers";
    data[0][1] = "Question";
    var answer = [];
    var stat = [];
    var n = test.AnswerToQuestion[questionNumber].length;
    for (var i = 0; i < n; i++) {
        answer.push(answers[i]);
        stat.push(stats["question"][questionNumber]["Stats"][i])
    }

    var temp = [];

    for (var i = 0; i < n; i++) {
        temp = [];
        temp.push(answer[i]);
        temp.push(stat[i]);
        data.push(temp);
    }
    temp = [];
    temp.push("Your Contribution / Вклад Вашого Голосу");
    temp.push(1);
    data.push(temp);
    console.dir(data);
    var d = google.visualization.arrayToDataTable(data);
    var options = {
        title: test.Questions.Question[questionNumber],
        pieHole: 0.4,
    };
    var chart = new google.visualization.PieChart(document.getElementById('donutchart'));

    chart.draw(d, options);
}

function signOut() {
    console.log("out")

    $.session.set("token", "");
    $.session.set("user", "");
    location.reload();
}
