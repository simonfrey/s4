function setError(msg) {
  document.getElementById('error').style.display = msg ? 'block' : 'none';
  document.getElementById('errorText').innerHTML = msg;
}

function handleOnline(){
  document.getElementById("onOffline").innerHTML = navigator.onLine 
    ? "Your browser is online. Please disconnect from your network to use s4 in a safer way" 
    : "";
}

function handleFile() {
  document.getElementById("onFile").innerHTML = window.location.protocol !== 'file:' 
    ? "You are running s4 from a webserver, this is insecure. Please save the webpage (Ctrl+S) locally and open this file." 
    : "";
}

// Update the online status icon based on connectivity
window.addEventListener('online', handleOnline);
window.addEventListener('offline',  handleOnline);

// initialize
handleOnline();
handleFile();