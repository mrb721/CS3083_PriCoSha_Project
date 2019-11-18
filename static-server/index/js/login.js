document.addEventListener("DOMContentLoaded", function () {
	document.getElementById("register").style.display = "none";
	document.getElementById("toggleForm").addEventListener("click", toggleForm);
	$("form").submit(function (e) {
		e.preventDefault();
	});
	$("#register").submit(function (e) {
		if (document.getElementById("signupPass").value != document.getElementById("confirmPass").value) {
			return;
		}
		data = JSON.stringify({
			Username: document.getElementById("regUname").value,
			Password: document.getElementById("signupPass").value,
			Fname: document.getElementById("regFname").value,
			Lname: document.getElementById("regLname").value
		});
		var xhr = new XMLHttpRequest();
		xhr.onreadystatechange = function () {
			if (xhr.readyState == XMLHttpRequest.DONE || xhr.status >= 200 && xhr.status < 300) {
				result = JSON.parse(this.responseText);
				if (result.Successful) {
					location.href = "index.html";
				}
			}
			else {
				location.reload();
			}
		};
		xhr.open("POST", "http://localhost:3000/register/user", true);
		// xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
		xhr.send(data);
	})
	$("#login").submit(function (e) {
		data = JSON.stringify({
			Username: document.getElementById("loginUname").value,
			Password: document.getElementById("loginPass").value
		});
		var xhr = new XMLHttpRequest();
		xhr.onload = function () {
			if (xhr.readyState == XMLHttpRequest.DONE || xhr.status >= 200 && xhr.status < 300) {
				result = JSON.parse(this.responseText);
				console.log(result.ErrMsg)
				if (!result.ErrMsg) {
					document.cookie = "token=" + result.Token;
					document.cookie = "user=" + document.getElementById("loginUname").value;
					document.cookie = "expires=" + result.ExpirationTime;
					location.href = "wall.html";
				}
			}
			else {
				location.href = "index.html";
			}
		};
		xhr.open("POST", "http://localhost:3000/login/user", true);
		// xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
		xhr.send(data);
	})
});

function toggleForm() {
	var something = document.getElementById("register").style.display;
	if (something == "none") {
		document.getElementById("register").style.display = "block";
		document.getElementById("login").style.display = "none";
		document.getElementById("toggleForm").innerHTML = "Click here to login";
	}
	else {
		document.getElementById("register").style.display = "none";
		document.getElementById("login").style.display = "block";
		document.getElementById("toggleForm").innerHTML = "Click here to register";
	}
}