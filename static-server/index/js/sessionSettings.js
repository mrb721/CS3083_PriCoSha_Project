document.addEventListener("DOMContentLoaded", function () {
    data = JSON.stringify({
        Username: getCookiePart("user"),
        ColorPalette: "000000",
        Intent: 'gcp',
        Token: getCookiePart("token"),
        User: getCookiePart("user")
    });
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        result = JSON.parse(this.responseText);
        document.getElementById("header").style.backgroundColor = "#" + result.ColorPalette;
    };
    xhr.open("POST", "http://localhost:3000/settings", true);
    // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
    xhr.send(data);
});

function getCookiePart(cookiename) {
    // Get name followed by anything except a semicolon
    var cookiestring = RegExp("" + cookiename + "[^;]+").exec(document.cookie);
    // Return everything after the equal sign, or an empty string if the cookie name not found
    return decodeURIComponent(!!cookiestring ? cookiestring.toString().replace(/^[^=]+./, "") : "");
}

function clearCookie() {
    var cookies = document.cookie.split(";");

    for (var i = 0; i < cookies.length; i++) {
        var cookie = cookies[i];
        var eqPos = cookie.indexOf("=");
        var name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
        document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT";
    }
}