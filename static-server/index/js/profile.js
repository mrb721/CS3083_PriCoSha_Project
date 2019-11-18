document.addEventListener("DOMContentLoaded", function () {
	document.getElementById("settingsButton").addEventListener("click", goToSettings);
});

function goToSettings() {
	//why ?????
	location.replace("settings.html");
	//return false;
}