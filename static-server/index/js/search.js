document.addEventListener("DOMContentLoaded", function () {
	document.getElementById("wallButton").addEventListener("click", goToWall);
});

function goToWall() {
	location.replace("wall.html");
	// return false;
}