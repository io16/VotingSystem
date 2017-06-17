/**
 * Created by igor on 15.06.17.
 */
$(document).ready(function () {
    console.log("ready!");
    var t = getSessionToken();
    if ( t.toString().length>20 )
    {
        console.log(getSessionUser())
        document.getElementById("sign-in-a").style.display = "none";
        document.getElementById("reg-a").style.display = "none";
        document.getElementById("sign-out-a").style.display = "block";
        document.getElementById("user-a").style.display = "block";
        document.getElementById("user-a").innerHTML = getSessionUser()
    }



});

function getSessionToken() {
    return $.session.get("token");
}function getSessionUser() {
    return $.session.get("user");
}

var countQuestions = 0;

function insertQuestions() {
    countQuestions = document.getElementById('countQuestions').value;
    document.getElementById('questions').innerHTML = '';
    var result = '';
    for (var i = 1; i <= countQuestions; i++) {

        result += '<div class="form-group"> <label for="question-' + i + '-countAnswer">Count Answers to question №' + i + ':</label> <input type="text" class="form-control" id="question-' + i + '-countAnswer" required> <label>Type Answer</label> <label class="radio-inline"><input type="radio" name="optradio' + i + '" checked>radio </label> <label class="radio-inline"><input type="radio" name="optradio' + i + '">checkbox </label> </div>';
    }


    document.getElementById('questions').innerHTML = result;
    document.getElementById("createTestButton").style.display = "block";
}
$('#createTestTemplate').submit(function () {

    return false;
});

function insertAnswers() {
    document.getElementById('questionsWithFields').innerHTML = '';
    var result = '';
    var countAnswer;
    var status = true;

    for (var i = 1; i <= countQuestions; i++) {
        countAnswer = document.getElementById('question-' + i + '-countAnswer').value;
        if (countAnswer == '') {
            status = false;
            document.getElementById('question-' + i + '-countAnswer').focus();
            break;
        }
    }
    if (status) {

        for (var i = 1; i <= countQuestions; i++) {
            countAnswer = document.getElementById('question-' + i + '-countAnswer').value;
            result += '<div class="form-group"> <label for="question-' + i + '-text"> Question № ' + i + ' </label> <input type="text" class="form-control" id="question-' + i + '-text" required> ';
            for (var j = 1; j <= countAnswer; j++) {
                result += '<label for="question-' + i + '-answer-' + j + '">Answer № ' + j + ' </label> <input type="text" class="form-control" id="question-' + i + '-answer-' + j + '" required> </div> ';
            }
        }
        document.getElementById('questionsWithFields').innerHTML = result;
        document.getElementById("sendTestButton").style.display = "block";
    }


}
$('#createTestAnswers').submit(function () {

    return false;
});

function saveTest() {
    var countAnswer;
    var status = true;
    for (var i = 1; i <= countQuestions; i++) {
        countAnswer = document.getElementById('question-' + i + '-countAnswer').value;
        var question = document.getElementById('question-' + i + '-text');
        if (question.value == '') {
            status = false;
            question.focus();
            break;
        }
        for (var j = 1; j <= countAnswer; j++) {
            var answer = document.getElementById('question-' + i + '-answer-' + j);
            if (answer.value == '') {
                status = false;
                answer.focus();
                break;
            }
        }

    }

    if (status) {
        var answToQuest = new Object();
        answToQuest.answerToquestion = [];
        var tempArr = [];
        var questions = new Object();
        questions.question = [];
        questions.type = [];
        for (var i = 1; i <= countQuestions; i++) {
            countAnswer = document.getElementById('question-' + i + '-countAnswer').value;

            questions.question.push(document.getElementById('question-' + i + '-text').value);
            if (document.getElementsByName('optradio' + i)[0].checked) {
                questions.type.push("radio");
            }
            else {
                questions.type.push("checkbox");
            }
            tempArr = [];
            for (var j = 1; j <= countAnswer; j++) {
                tempArr.push(document.getElementById('question-' + i + '-answer-' + j).value)
            }
            answToQuest.answerToquestion.push(tempArr);

        }



        var obj = new Object();
        obj.name = $('#testName').val();
        obj.category = $('#testCategory').val();
        obj.questions = questions;

        obj.answertoquestion=answToQuest.answerToquestion;
        console.dir(obj)
        $.post("/savevote", {
                data: JSON.stringify(obj)
            }

        )
    }

}