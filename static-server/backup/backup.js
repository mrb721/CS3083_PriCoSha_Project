userList = document.getElementById('user-list');
netConn = document.getElementById('net-connectivity');

let failureCounter = 0;
let successCounter = 0;

function registerFailedConnection(failMessage) {
  console.log(failMessage)
  failureCounter += 1;
  netVal.textContent = "Unable to contact the server.";
  netVal.style = 'color: red'
  netConn.style = 'color: red'
  netConn.textContent = "Consecutive failed pings: " + failureCounter;
  successCounter = 0;
}

function updateValue(msg) {
  userList.textContent = msg;
}

function updateUserList(response) {
  console.log("received a response: " + response)
  successCounter += 1;
  userList.textContent = response;
  netConn.textContent = "Consecutive successful pings: " + successCounter;
  netConn.style = 'color: green'
  netVal.style = 'color: black'
  failureCounter = 0;
}

function getAllUsers() {
  $.ajax({
    url: "http://localhost:3000/get/users/all",
    type: "get",
    success: (response) => {
      updateValue(response)
    },
    error: (xhr) => {
      registerFailedConnection(xhr)
    }
  });
}

let log = document.getElementById('net-log')

// function outputLog(msg) {
//   const node = document.createElement("p");
//   x = new Date();
//   node.textContent = x.toLocaleTimeString() + ' - ' + msg;
//   log.insertBefore(node, log.firstChild);
//   log.scrollTo(0,0);
// }

document.getElementById('get-users').addEventListener("click", () => {
  getAllUsers()
  outputLog('Get request sent!')
})

window.setInterval(() => {
  updateUserList()
}, 500)
