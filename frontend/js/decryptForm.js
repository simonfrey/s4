const base64regex = /^([0-9a-zA-Z+/]{4})*(([0-9a-zA-Z+/]{2}==)|([0-9a-zA-Z+/]{3}=))?$/;

var decryptState = {
  shares: 2,
  inputs: ['', ''],
  output: ''
}

function updateDecrypt(obj) {
  mergeDeep(decryptState, obj)
  calcDecryptOutput();
  render();
}

function renderDecrypt() {
  return el('div', {class:'decryptForm columns'}, [
    el('div', {className:'column is-vertical'}, [
      el('div', {className:'columns is-horizontal'}, [
        el('div', {className:'column'}, [
          'Shares: ',
          el('input', {
            className: 'input', 
            type: 'number',
            min: 2,
            value: decryptState.shares,
            placeholder: 'Shares',
            oninput: function(e) {
              var val = Number(e.target.value);
              updateDecrypt({shares: val, inputs: Array(val).fill('')});
            }
          })
        ])
      ]),
      el('div', {className:'columns is-horizontal'}, renderInputs()),
      el('div', {className:'columns'}, [
        el('div', {className:'column'}, [
          el('textarea', {
            id: 'output',
            className:'textarea',
            placeholder: 'Your decrypted output'
          }, decryptState.output)
        ])
      ])
    ])
  ]);
}

function renderInputs(){
  return decryptState.inputs.map(function(item, i){
    return el('div', {className:'column'}, [
      el('textarea', {
        className:'textarea', 
        placeholder: 'Your encrypted input share',
        oninput: function(e){
          var pos = e.target.selectionEnd;
          decryptState.inputs[i] = e.target.value;
          updateDecrypt({inputs: decryptState.inputs});
        }
      }, item)
    ]);
  });
}

function calcDecryptOutput() {
  res = Recover_fours(decryptState.inputs);

  if (!base64regex.test(res)){
      decryptState.output = ""
      state.error = res
      return;
  }

  decryptState.output = atob(res)
  state.error = ""
}