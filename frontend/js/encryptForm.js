var encryptState = {
  requiredShares: 2,
  totalShares: 2,
  inputString: '',
  aes: true,
  outputs: []
}

function updateEncrypt(obj) {
  mergeDeep(encryptState, obj)
  calcEncryptOutputs();
  render();
}

function renderEncrypt() {
  return el('div', {class:'encryptForm'}, [
    el('div', {className:'columns is-horizontal is-vcentered'}, [
      el('div', {className:'column'}, [
        'Minimum required shares:',
        el('input', {
          className: 'input', 
          type: 'number',
          min: 2, 
          max: encryptState.totalShares,
          value: encryptState.requiredShares,
          placeholder: 'Minimum required shares',
          oninput: function(e) {
            var val = Number(e.target.value);
            if (val > encryptState.totalShares) {
              val = encryptState.totalShares;
            }
            updateEncrypt({requiredShares: val});
          }
        })
      ]),
      el('div', {className:'column'}, [
        'Total Shares: ',
        el('input', {
          className: 'input', 
          type: 'number',
          min: 2, 
          value: encryptState.totalShares,
          placeholder: 'Shares',
          oninput: function(e) {
            var val = Number(e.target.value);
            updateEncrypt({totalShares: val});
          }
        })
      ]),
      el('div', {className:'column'}, [
        el('label', {className: 'checkbox'}, [
          el('input', {
            type: 'checkbox', 
            className: 'checkbox', 
            checked: encryptState.aes, 
            onchange: function(e){
              updateEncrypt({aes: e.target.checked})
            }
          }),
          ' Use AES(256 bit) for data and distribute key'          
        ])
      ])
    ]),
    el('div', {className:'columns'}, [
      el('div', {className:'column'}, [
        el('textarea', {
          id: 'inputString',
          className:'textarea',
          placeholder: 'Whatever message / binary data you want.',
          oninput: function(e) {
            var pos = e.target.selectionEnd;
            updateEncrypt({inputString: e.target.value});
            var newText = document.getElementById('inputString');
            newText.focus();
            newText.selectionStart = pos;
            newText.selectionEnd = pos;
          }
        }, encryptState.inputString)
      ])
    ]),
    el('div', {className:'columns is-horizontal'}, renderOutBoxes()),
  ]);
}

function renderOutBoxes(){
  return encryptState.outputs.map(function(item){
    return el('div', {className:'column'}, [
      el('textarea', {className:'textarea', readOnly: true}, item)
    ]);
  });
}

function calcEncryptOutputs() {
  encryptState.outputs = Array(encryptState.totalShares).fill("");
  if (encryptState.inputString.length === 0 || encryptState.totalShares <= 0 || encryptState.requiredShares <= 0 || encryptState.requiredShares>encryptState.totalShares) {
    return  
  }

  var res = Distribute_fours((String(btoa(encryptState.inputString))),Number(encryptState.totalShares),Number(encryptState.requiredShares),Boolean(encryptState.aes));
  if (typeof res === 'string'){
    state.error = res;
    encryptState.outputs = Array(encryptState.totalShares).fill("")
  }

  encryptState.outputs = res;
}
