function doEncrypt() {
  var thresholdEl = document.getElementById('threshold');
  var sharesEl = document.getElementById('shares');
  
  // validation
  if (sharesEl.valueAsNumber <= 1) {
    sharesEl.value = 2;
  }

  if (thresholdEl.valueAsNumber <= 1) {
    thresholdEl.value = 2;
  }

  // ensure threshold isn't higher than shares
  thresholdEl.setAttribute("max", sharesEl.valueAsNumber);
  if (sharesEl.valueAsNumber < thresholdEl.valueAsNumber) {
    thresholdEl.value = sharesEl.value;
  }

  // snag values from the dom
  var threshold = thresholdEl.valueAsNumber;
  var shares = sharesEl.valueAsNumber;
  var useAES = document.getElementById('useAES').checked;
  var input = document.getElementById('input').value;

  // handle no input
  if (input == "") {
    return fillOutputs(Array(shares).fill(""));
  }

  // retry if wasm didn't load yet.
  if (typeof Distribute_fours == "undefined") {
    setTimeout(doEncrypt, 3000);
    return setError('WASM not loaded, retrying...');
  }

  // do it!
  var res = Distribute_fours(
    String(btoa(input)),
    Number(shares),
    Number(threshold),
    Boolean(useAES)
  );

  // update dom w/ results
  if (typeof res === 'string'){
    setError(res)
    fillOutputs(Array(shares).fill(""));
  } else {
    setError('');
    fillOutputs(res);
  }
}

function fillOutputs(values) {
  var out = document.getElementById("outputs");
  out.innerText = '';

  values.forEach(function(o) {
    var div = document.createElement('div');
    div.className = "column";

    var ta = document.createElement("textarea");
    ta.className = 'textarea';
    ta.readOnly = 'readonly';
    ta.innerHTML = o;

    div.appendChild(ta);
    out.appendChild(div);
  });
}
