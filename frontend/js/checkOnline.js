const onlineText = "Your browser is online. Please disconnect from your network to use s4 in a safer way";

const onOffElem = document.getElementById("onOffline");

function setInfo(offline){
    if (offline){
        //Offline
    onOffElem.innerHTML = "";
    }else{
        // Online
    onOffElem.innerHTML = onlineText;
    }
}

if(!navigator.onLine){
    // Offline
    setInfo(true);
}

function updateIndicator() {
    setInfo(!navigator.onLine);
}

// Update the online status icon based on connectivity
window.addEventListener('online', updateIndicator);
window.addEventListener('offline',  updateIndicator);
updateIndicator();


//
// Check for local file
if (window.location.protocol == 'file:'){
    const onFileElem = document.getElementById("onFile");
    onFileElem.innerHTML = "";
}