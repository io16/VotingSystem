/**
 * Created by igor on 15.06.17.
 */
var countQuestions = 0;

function insertQuestions() {
    countQuestions = document.getElementById('countQuestions').value;
    document.getElementById('questions').innerHTML = '';
    var result = '';
    for (var i = 1; i <= countQuestions; i++) {

        result += '<div class="form-group"> <label for="question-' + i + '-countAnswer">Count Answers to question №' + i + ':</label> <input type="text" class="form-control" id="question-' + i + '-countAnswer" required> <label>Type Answer</label> <label class="radio-inline"><input type="radio" name="optradio' + i + '" checked>radio </label> <label class="radio-inline"><input type="radio" name="optradio' + i + '">checkbox </label> </div>';
    }

    document.getElementById('questions').innerHTML = result
    document.getElementById("createTestButton").style.display = "block";
}
$('#createTestTemplate').submit(function () {

    return false;
});

function insertAnswers() {
    document.getElementById('questionsWithFields').innerHTML = '';
    var result = '';

    var countAnswer;
    for (var i = 1; i <= countQuestions; i++) {
        countAnswer = document.getElementById('question-' + i + '-countAnswer').value;
        result += '<div class="form-group"> <label for="question-' + i + '-text"> Question № ' + i + ' </label> <input type="text" class="form-control" id="question-' + i + '-text" required> '
        for (var j = 1; j <= countAnswer; j++) {
            result += '<label for="question-' + i + '-answer-' + j + '">Answer № ' + j + ' </label> <input type="text" class="form-control" id="question-' + i + '-answer-' + j + '" required> </div> ';
        }
    }
    document.getElementById('questionsWithFields').innerHTML = result;
    document.getElementById("sendTestButton").style.display = "block";
}
$('#createTestAnswers').submit(function () {

    return false;
});