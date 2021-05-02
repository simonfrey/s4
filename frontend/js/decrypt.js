// trigger initial render for decrypt tab
window.addEventListener('DOMContentLoaded', function() {
  handleShareChange();
});


function handleShareChange() {
  var total = document.getElementById("decryptShares").valueAsNumber;
  var existing = document.querySelectorAll("#inputs>div").length;

  // validate
  if (total < 2) {
    total = 2;
    document.getElementById("decryptShares").value = 2;
  }

  // add or remove textareas
  if (total < existing) {
    removeDecryptInputs(existing - total);
  } else if (total > existing) {
    addDecryptInputs(total - existing);
  }
}

function addDecryptInputs(count) {
  var inputsEl = document.getElementById("inputs");
  for (var i=0; i<count;i++) {
    var div = document.createElement('div');
    div.className = "column";
  
    var ta = document.createElement("textarea");
    ta.className = 'textarea';
    ta.addEventListener('input', doDecrypt);

    div.appendChild(ta);
    inputsEl.append(div);
  }
}

function removeDecryptInputs(count) {
  var inputsEl = document.getElementById("inputs");
  for (var i=0; i<count;i++) {
    inputsEl.removeChild(inputsEl.lastChild);
  }
}

function doDecrypt() {
  // collect inputs
  var inputs = [];
  document.querySelectorAll("#inputs>div").forEach(function(div){
    inputs.push(div.children[0].value);
  });

  // check if they are all empty for no-op
  if (inputs.filter(Boolean).length === 0) {
    return;
  }
  
  var res = Recover_fours(inputs)
  var outEl = document.getElementById("output");
  var base64regex = /^([0-9a-zA-Z+/]{4})*(([0-9a-zA-Z+/]{2}==)|([0-9a-zA-Z+/]{3}=))?$/;

  if (!base64regex.test(res)){
    outEl.innerText = '';
    setError(res);
  } else {
    outEl.innerText = atob(res);
    setError('');
  }
}
