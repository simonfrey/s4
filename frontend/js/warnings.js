var warningState = {
  offline: false,
  file: false
}

// Update the offline status icon based on connectivity
function updateWarnings() {
  warningState.offline = !navigator.onLine;
  warningState.file = window.location.protocol == 'file:';
}

window.addEventListener('online', updateWarnings);
window.addEventListener('offline',  updateWarnings);
updateWarnings();

function renderWarnings() {
  return [
    warningState.offline ? '' : el('p', {className: 'has-text-danger'}, 'Your browser is online. Please disconnect from your network to use s4 in a safer way'),
    warningState.file ? '' : el('p', {className: 'has-text-danger'}, 'You are running s4 from a webserver, this is insecure. Please savethe webpage (Ctrl+S) locally and open this file.')
  ];
}