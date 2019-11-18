function querySearch() {
    data = JSON.stringify({
        ID: document.getElementById("searchByName"),
        GroupName: document.getElementById('searchByFriendGroup'),
        Username: document.getElementById('searchByUname'),
        Intent: 'mk',
        Token: getCookiePart('token'),
        User: getCookiePart('user')
    });
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        result = JSON.parse(this.responseText);
        var zA, zB, zC, zD, zE;
        for (var i = 0; i < result.UserResult.length; i++) {
            var cUserRes = result.UserResult[i];
            zA = cUserRes.Username;
            zB = cUserRes.Firstname + " " + cUserRes.LastName;
            var cGroupRes = result.GroupResult[i];
            zC = cGroupRes.Creator;
            zD = cGroupRes.Description;
            zE = cGroupRes.GroupName;
            addSearchResult(zA, zB, zC, zD, zE);
        }
    }
    xhr.open("POST", "http://localhost:3000/search", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
}

document.addEventListener("DOMContentLoaded", function () {
    element = document.getElementById("search");
    if (element.addEventListener) {
        element.addEventListener("submit", function (evt) {
            evt.preventDefault();
            querySearch();
        }, true);
    }
    else {
        element.attachEvent('onsubmit', function (evt) {
            evt.preventDefault();
            querySearch();
        });
    }
});


//a: username, string
//b: user's name, string
//c: creator, string
//d: content???? WTF???
//e: friendgroup, string
function addSearchResult(a, b, c, d, e) {

    var searchResultHTML = `
            <div class="resultBox">
                <div>` + a + `</div>
                <div class="personName">` + b + `</div>
                <div class="userName">` + c + `</div>
                <div class="postContent">` + d + `</div>
                <div class="FriendGroup">` + e + `</div>
            </div>`;

    var whereToPutResult = document.getElementById("plugInSearchResult");
    whereToPutResult.innerHTML += searchResultHTML;
}