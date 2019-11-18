//postType 0 is text
//postType 1 is photo

var postType = 0;
var modal;

document.addEventListener("DOMContentLoaded", function () {
    var gcForm = document.getElementById("showComment");
    if (gcForm.addEventListener) {
        gcForm.addEventListener("submit", function (evt) {
            evt.preventDefault();
            createComment();
        }, true);
    }
    else {
        gcForm.attachEvent('onsubmit', function (evt) {
            evt.preventDefault();
            createComment();
        });
    }

    var shForm = document.getElementById("showCommentActual");
    if (gcForm.addEventListener) {
        gcForm.addEventListener("submit", function (evt) {
            evt.preventDefault();
        }, true);
    }
    else {
        gcForm.attachEvent('onsubmit', function (evt) {
            evt.preventDefault();
        });
    }

    document.getElementById("postTypeToggle").addEventListener("click", function (e) {
        if (!postType) {
            document.getElementById("photoField").style.display = "block";
            document.getElementById("textField").style.display = "none";
            document.getElementById("postTypeToggle").innerHTML = "Click here to post text";
            postType = 1;
        }
        else {
            document.getElementById("photoField").style.display = "none";
            document.getElementById("textField").style.display = "block";
            document.getElementById("postTypeToggle").innerHTML = "Click here to post photo";
            postType = 0;
        }
    })

    // Get the modal
    modal = document.getElementById('postModal');
    // Get the button that opens the modal
    var btn = document.getElementById("createPost");
    // Get the <span> element that closes the modal
    var span = document.getElementsByClassName("close")[0];
    // When the user clicks the button, open the modal 
    btn.onclick = function () {
        modal.style.display = "block";
    }
    // When the user clicks on <span> (x), close the modal
    span.onclick = function () {
        modal.style.display = "none";
    }
    // Get the modal
    var commentModal = document.getElementById('commentsModal');
    // Get the button that opens the modal
    var commentbtn = document.getElementById("activateComments");
    // Get the <span> element that closes the modal
    var commentSpan = document.getElementsByClassName("close")[1];
    var tagModal = document.getElementById('tagsModal');
    // Get the button that opens the modal
    var tagbtn = document.getElementById("activateTags");
    // Get the <span> element that closes the modal
    var tagSpan = document.getElementsByClassName("close")[2];

    window.onclick = function (event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
        if (event.target == commentModal) {
            commentModal.style.display = "none";
        }
        if (event.target == tagModal) {
            tagModal.style.display = "none";
        }
    }


    var cpf = document.getElementById("createPostForm");
    if (cpf.addEventListener) {
        cpf.addEventListener("submit", function (evt) {
            evt.preventDefault();
        }, true);
    }
    else {
        cpf.attachEvent('onsubmit', function (evt) {
            evt.preventDefault();
        });
    }

    document.getElementById("uploadPost").addEventListener("click", function (e) {
        createPost("t");
        modal.style.display = "none";

    });
    document.getElementById("uploadPhoto").addEventListener("click", function (e) {
        createPost("p");
        modal.style.display = "none";
    });

    document.getElementById("editPostBtn").addEventListener("click", function (e) {
        editPost();
    });


    //show feed
    data = JSON.stringify({
        // Intent: 'get',
        Token: getCookiePart("token"),
        User: getCookiePart("user"),
        Intent: 'public'
    });
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        result = JSON.parse(this.responseText);
        for (var i = 0; i < result.RawContents.length; i++) {
            var p = result.RawContents[i];
            addWallPost(p.ID, p.ContentName, p.UsernameCreator, p.Content, p.ContentType);
        }

    };
    xhr.open("POST", "http://localhost:3000/contentByGroup", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);

    document.getElementById('picFiles').addEventListener('change', picFileByteArray);
});

function getComment(theID) {
    var minosaur = document.getElementById("placeToShowComments");
    minosaur.innerHTML = "";
    data = JSON.stringify({
        ID: theID,
        Intent: 'get',
        Token: getCookiePart("token"),
        User: getCookiePart("user")
    });
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        result = JSON.parse(this.responseText);
        console.log(result);
        for (var i = 0; i < result.Comments.length; i++) {
            var cc = result.Comments[i];
            showComHTML(cc.ID, cc.Username, cc.Timestamp, cc.CommentText);
        }
    };
    xhr.open("POST", "http://localhost:3000/comment", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
}

function showComHTML(id, user, time, text) {
    var newComHTML = `
    <div>
        <b>`+ user + `</b>
        <p>`+ text + `</p>
        <i>`+ time + `</i>
        <br>
    </div>
    `;
    document.getElementById("placeToShowComments").innerHTML += newComHTML;
}

function createComment() {
    data = JSON.stringify({
        ID: parseInt(document.getElementById("postIDtoComment").value),
        CommentText: document.getElementById("newComment").value,
        Intent: 'mk',
        Token: getCookiePart("token"),
        User: getCookiePart("user")
    });
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        result = JSON.parse(this.responseText);
        if (result.Successful) {
            console.log("created comment successful")
        }
    };
    xhr.open("POST", "http://localhost:3000/comment", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
}

//id: post id, int
//h: head title, string
//i: user name, string
//b: post body, byte array
function addWallPost(id, h, i, b, t) {
    var wall = document.getElementById("theWall");
    var tmp = document.createElement("div");

    postHTML = `
    <div id="wPost` + id + `" class="aPost">
        <div><b>` + h + `</b></div>
        <div>by <i>` + i + `</i></div>
        
		<div class="postBody">
        `;

    if (t == ".txt") {
        postHTML += "<p>" + bin2StringBetter(b) + "</p>";
    }
    else {
        postHTML += '<img id="postImg' + id + '" style="height:100%; width:100%;">';
    }

    postHTML += `
        <br>
        <div class="postButts" id="wpBtn` + id + `"> Edit Post </div>
        <div class="postButts" id="shBtn` + id + `"> Share Post </div>
        <div class="postButts" id="dpBtn` + id + `"> Delete Post </div> <br>
        <div class="postButts" id="tgBtn` + id + `"> Tag in Post </div>
        <div class="postButts" id="stBtn` + id + `"> Show tags </div> <br>
        <div class="postButts" id="acBtn` + id + `"> Create Comment </div>
        <div class="postButts" id="showCommentsFor` + id + `">Show comments</button>

		</div>
		<br>	
		</div>
	</div>`;

    tmp.innerHTML = postHTML;
    wall.appendChild(tmp);
    if (t != ".txt") {
        showByteArrayImage("postImg" + id, b);
    }
    addEventListenersToPost(id);
    if (i.toLowerCase() != getCookiePart("user")) {
        document.getElementById("dpBtn" + id).style.display = "none";
    }
}

function addEventListenersToPost(id) {
    document.getElementById("dpBtn" + id).addEventListener("click", deletePost);
    document.getElementById("wpBtn" + id).addEventListener("click", editPostRunner);
    document.getElementById("acBtn" + id).addEventListener("click", newComment);
    document.getElementById("showCommentsFor" + id).addEventListener("click", showThemComments);
    document.getElementById("shBtn" + id).addEventListener("click", sharePost);
    document.getElementById("tgBtn" + id).addEventListener("click", tagPost);
    document.getElementById("stBtn" + id).addEventListener("click", showTagz);
}

function deletePost(e) {
    data = JSON.stringify({
        ID: parseInt(e.target.id.substring(5, e.target.id.length)),
        Intent: "rm",
        Token: getCookiePart("token"),
        User: getCookiePart("user"),
    });
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        result = JSON.parse(this.responseText);
        location.reload();
    };
    xhr.open("POST", "http://localhost:3000/content", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
}

function showTagz(e) {
    document.getElementById("placeToShowTags").innerHTML = "";
    data = JSON.stringify({
        ID: parseInt(e.target.id.substring(5, e.target.id.length)),
        UsernameTagger: getCookiePart("user"),
        UsernameTaggee: document.getElementById("personToTag").value,
        Intent: "get",
        Token: getCookiePart("token"),
        User: getCookiePart("user"),
    });
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        result = JSON.parse(this.responseText);

        console.log(result);
        for (var i = 0; i < result.Tags.length; i++) {
            var cTag = result.Tags[i];
            addTags(cTag.UsernameTagger, cTag.UsernameTaggee, cTag.Timestamp, cTag.Status, cTag.ID);

        }
    };
    xhr.open("POST", "http://localhost:3000/tag", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
}


function addTags(a, b, c, d, e) {
    var p = document.getElementById("placeToShowTags");
    if (a.toLowerCase() == getCookiePart("user").toLowerCase() || b.toLowerCase() == getCookiePart("user").toLowerCase()) {
        var tHTML = `
        <div class="aTagThing">
            <div id="taggr` + e + `">` + a + `</div>
            <div>` + b + `</div>
            <div>` + c + `</div>`;
        if (d || b.toLowerCase() != getCookiePart("user").toLowerCase()) {
            tHTML += '</div>';
        }
        else {
            tHTML += '<button id="apTag' + e + '"> Approve Tag </button></div>';
        }
        p.innerHTML += tHTML;
        if (!d && b.toLowerCase() == getCookiePart("user").toLowerCase()) {
            document.getElementById("apTag" + e).addEventListener("click", function (e) {
                data = JSON.stringify({
                    ID: parseInt(e.target.id.substring(5, e.target.id.length)),
                    UsernameTagger: document.getElementById("taggr" + e.target.id.substring(5, e.target.id.length)).innerHTML,
                    UsernameTaggee: getCookiePart("user"),
                    Intent: "ap",
                    Token: getCookiePart("token"),
                    User: getCookiePart("user"),
                });
                var xhr = new XMLHttpRequest();
                xhr.onload = function () {
                    result = JSON.parse(this.responseText);
                    location.reload();
                };
                xhr.open("POST", "http://localhost:3000/tag", true);
                // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
                xhr.send(data);
            })
        }
    }
}

function tagPost(e) {
    data = JSON.stringify({
        ID: parseInt(e.target.id.substring(5, e.target.id.length)),
        UsernameTagger: getCookiePart("user"),
        UsernameTaggee: document.getElementById("personToTag").value,
        Intent: "mk",
        Token: getCookiePart("token"),
        User: getCookiePart("user"),
    });
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        if (xhr.readyState == XMLHttpRequest.DONE || xhr.status >= 200 && xhr.status < 300) {
            result = JSON.parse(this.responseText);
            location.reload();
        }
        else {
            console.log("failed to tag");
        }
    };
    xhr.open("POST", "http://localhost:3000/tag", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
}

function sharePost(e) {
    data = JSON.stringify({
        ID: parseInt(e.target.id.substring(5, e.target.id.length)),
        GroupName: document.getElementById("minosaurShare"), //input for group name
        Creator: getCookiePart("user"),
        Intent: "mk",
        Token: getCookiePart("token"),
        User: getCookiePart("user"),
    });

    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        if (xhr.readyState == XMLHttpRequest.DONE || xhr.status >= 200 && xhr.status < 300) {
            result = JSON.parse(this.responseText);
            location.reload();
        }
        else {
            console.log("failed to share post");
        }
    };
    xhr.open("POST", "http://localhost:3000/share", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
}

function showThemComments(e) {
    var minosaur = document.getElementById("placeToShowComments");
    minosaur.innerHTML = "";
    var theID = e.target.id.substring(15, e.target.id.length);
    document.getElementById("postIDshowComment").value = theID;
    getComment(parseInt(theID));
}

function editPostRunner(e) {
    modal.style.display = "block";
    document.getElementById("createPostBtn").style.display = "none";
    document.getElementById("uploadPost").style.display = "none";
    document.getElementById("editPostBtn").style.display = "block";
    document.getElementById("postID").value = e.target.id.substring(5, e.target.id.length);
}

function newComment(e) {
    document.getElementById("showComment").style.display = "block";
    document.getElementById("postIDtoComment").value = e.target.id.substring(5, e.target.id.length);
}

function createPost(e) {
    var data;
    if (e == "t") {
        data = JSON.stringify({
            Username: getCookiePart("user"),
            ContentName: document.getElementById("postTitle").value,
            ContentType: "txt",
            Content: string2Bin(document.getElementById("postText").value),
            Intent: "mk",
            Token: getCookiePart("token"),
            User: getCookiePart("user")
        });
    }
    else {
        data = JSON.stringify({
            Username: getCookiePart("user"),
            ContentName: picFileDetails.name.split(".")[0],
            ContentType: picFileDetails.name.split(".")[1],
            Content: picFileToUpload,
            Intent: "mk",
            Token: getCookiePart("token"),
            User: getCookiePart("user")
        });
    }

    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        result = JSON.parse(this.responseText);
    };
    xhr.open("POST", "http://localhost:3000/content", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
}

function editPost() {

    data = JSON.stringify({
        ID: parseInt(document.getElementById("postID").value),
        Username: getCookiePart("user"),
        ContentName: document.getElementById("postTitle").value,
        ContentType: "txt",
        Content: string2Bin(document.getElementById("postText").value),
        Intent: 'ed',
        Token: getCookiePart("token"),
        User: getCookiePart("user")
    });

    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        result = JSON.parse(this.responseText);
        if (result.Successful) {
            location.reload();
        }
    };
    xhr.open("POST", "http://localhost:3000/content", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
}

function bin2String(array) {
    var result = "";
    for (var i = 0; i < array.length; i++) {
        result += String.fromCharCode(array[i]);
    }
    return result;
}

function string2Bin(str) {
    var result = [];
    for (var i = 0; i < str.length; i++) {
        result.push(str.charCodeAt(i));
    }
    return result;
}

decodeBase64 = function (s) {
    var e = {}, i, b = 0, c, x, l = 0, a, r = '', w = String.fromCharCode, L = s.length;
    var A = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
    for (i = 0; i < 64; i++) { e[A.charAt(i)] = i; }
    for (x = 0; x < L; x++) {
        c = e[s.charAt(x)]; b = (b << 6) + c; l += 6;
        while (l >= 8) { ((a = (b >>> (l -= 8)) & 0xff) || (x < (L - 2))) && (r += w(a)); }
    }
    return r;
};

function bin2StringBetter(a) {
    var result = "";
    return decodeBase64(a);
}


//use this function to set an DOM image object to byte array
function showByteArrayImage(id, bytes) {
    document.getElementById(id).src = "data:image/png;base64," + bytes;
}

var picFileToUpload;
var picFileDetails;
function picFileByteArray(evt) {
    var files = evt.target.files;
    for (var i = 0, f; f = files[i]; i++) {
        var reader = new FileReader();
        reader.onload = function () {
            var arrayB = this.result,
                arr = new Uint8Array(arrayB)
            arrRes = [];

            for (var mino = 0; mino < arr.length; mino++) {
                arrRes.push(arr[mino]);
            }

            picFileToUpload = arrRes;

        }
        reader.readAsArrayBuffer(f);
        picFileDetails = f;
    }
};