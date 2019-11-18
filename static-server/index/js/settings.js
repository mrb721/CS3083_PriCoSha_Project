document.addEventListener("DOMContentLoaded", function () {
    var allColors = document.getElementsByClassName("setColor");
    for (var i = 0; i < allColors.length; i++) {
        allColors[i].addEventListener("click", function (e) {
            data = JSON.stringify({
                Username: getCookiePart("user"),
                ColorPalette: rgbToHex(grabColor(this)),
                Intent: 'ccp',
                Token: getCookiePart("token"),
                User: getCookiePart("user")
            });
            var xhr = new XMLHttpRequest();
            xhr.onload = function () {
                result = JSON.parse(this.responseText);
                if (result.Successful) {
                    location.href = "settings.html";
                    return;
                }
                location.href = "settings.html";

            };
            xhr.open("POST", "http://localhost:3000/settings", true);
            // xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
            xhr.send(data);
        });
    }
});

function grabColor(e) {
    return e.style.backgroundColor;
}

function rgbToHex(col) {
    if (col.charAt(0) == 'r') {
        col = col.replace('rgb(', '').replace(')', '').split(',');
        var r = parseInt(col[0], 10).toString(16);
        var g = parseInt(col[1], 10).toString(16);
        var b = parseInt(col[2], 10).toString(16);
        r = r.length == 1 ? '0' + r : r; g = g.length == 1 ? '0' + g : g; b = b.length == 1 ? '0' + b : b;
        var colHex = '' + r + g + b;
        return colHex;
    }
}