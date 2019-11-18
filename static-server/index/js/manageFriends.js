document.addEventListener("DOMContentLoaded", function () {

	var cForm = document.getElementById("makeFriendGroupForm");
	if (cForm.addEventListener) {
		cForm.addEventListener("submit", function (evt) {
			evt.preventDefault();
			makeFriendGroup();
		}, true);
	}
	else {
		cForm.attachEvent('onsubmit', function (evt) {
			evt.preventDefault();
			makeFriendGroup();
		});
	}

	var jForm = document.getElementById("joinFriendGroupForm");
	if (jForm.addEventListener) {
		jForm.addEventListener("submit", function (evt) {
			evt.preventDefault();
			joinFriendGroup();
		}, true);
	}
	else {
		jForm.attachEvent('onsubmit', function (evt) {
			evt.preventDefault();
			joinFriendGroup();
		});
	}

	var rForm = document.getElementById("joinFriendGroupForm");
	if (rForm.addEventListener) {
		rForm.addEventListener("submit", function (evt) {
			evt.preventDefault();
			rmFriendGroup();
		}, true);
	}
	else {
		rForm.attachEvent('onsubmit', function (evt) {
			evt.preventDefault();
			rmFriendGroup();
		});
	}

	data = JSON.stringify({
		Intent: "memof",
		Token: getCookiePart("token"),
		User: getCookiePart("user"),
	});

	var xhr = new XMLHttpRequest();
	xhr.onload = function () {
		result = JSON.parse(this.responseText);
		for (var i = 0; i < result.Groups.length; i++) {
			var tmp = result.Groups[i];
			putUpGroup(tmp.GroupName, tmp.Username, tmp.Description);
		}
	};
	xhr.open("POST", "http://localhost:3000/group", true);
	xhr.send(data);

});

// dont do this
function goToSettings() {
	location.replace("settings.html");
	// return false;
}

function makeFriendGroup() {

	data = JSON.stringify({
		Creator: getCookiePart("user"),
		GroupName: document.getElementById("makeFGname").value,
		Description: document.getElementById("FGdesc").value,
		Intent: "mk",
		Token: getCookiePart("token"),
		User: getCookiePart("user"),
	});

	var xhr = new XMLHttpRequest();
	xhr.onload = function () {
		result = JSON.parse(this.responseText);
		if (result.Successful) {
			location.reload();
		}
	};
	xhr.open("POST", "http://localhost:3000/group", true);
	// xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
	xhr.send(data);

}

function joinFriendGroup() {

	data = JSON.stringify({
		Creator: getCookiePart("user"),
		GroupName: document.getElementById("joinFGname").value,
		Intent: "join",
		Token: getCookiePart("token"),
		User: getCookiePart("user"),
	});

	var xhr = new XMLHttpRequest();
	xhr.onload = function () {
		console.log("joined friend group successful");
	};
	xhr.open("POST", "http://localhost:3000/group", true);
	// xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
	xhr.send(data);

}

function joinFriendGroup() {
	data = JSON.stringify({
		Creator: getCookiePart("user"),
		GroupName: document.getElementById("rmFGname").value,
		Intent: "rm",
		Token: getCookiePart("token"),
		User: getCookiePart("user"),
	});

	var xhr = new XMLHttpRequest();
	xhr.onload = function () {
		console.log("removed friend group successful");
	};
	xhr.open("POST", "http://localhost:3000/group", true);
	// xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
	xhr.send(data);

}

function putUpGroup(a, c, d) {
	var tID = genID();
	var ptsg = document.getElementById("placeToShowGroups");
	var tmp = document.createElement("div");
	var aGroupHTML = `
	<div>
		<h2 class="` + tID + `"> ` + a + ` </h2>
		<br>
		<i class="` + tID + `">` + c + ` </i>
		<br>
		<p>` + d + `</p>
		<br>
		<button class="` + tID + `"> Leave FriendGroup </button>
	</div>
	`;
	tmp.innerHTML = aGroupHTML;
	ptsg.appendChild(tmp);
	var curEl = document.getElementsByClassName(tID);
	curEl[curEl.length - 1].addEventListener('click', leaveGroup)
}

function leaveGroup(e) {
	var curID = e.target.className;

	var curEl = document.getElementsByClassName(tID);

	data = JSON.stringify({
		Creator: curEl[1].innerHTML,
		GroupName: curEl[0].innerHTML,
		Intent: "leave",
		Token: getCookiePart("token"),
		User: getCookiePart("user"),
	});

	var xhr = new XMLHttpRequest();
	xhr.onload = function () {
		console.log("leave friend group successful");
	};
	xhr.open("POST", "http://localhost:3000/group", true);
	// xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
	xhr.send(data);
}

function genID() {
	return Math.floor((1 + Math.random()) * 0x10000);
}