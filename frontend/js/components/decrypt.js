class Encrypt extends Component {
  state = {
    error: '',
    requiredShares: 2,
    totalShares: 2,
    inputString: '',
    aes: true,
    outputs: []
  }

  calc() {
    this.state.outputs = Array(this.state.totalShares).fill("");
    if (this.state.inputString.length === 0 || this.state.totalShares <= 0 || this.state.requiredShares <= 0 || this.state.requiredShares>this.state.totalShares) {
      return  
    }

    var res = Distribute_fours((String(btoa(this.state.inputString))),Number(this.state.totalShares),Number(this.state.requiredShares),Boolean(this.state.aes));
    if (typeof res === 'string'){
      this.state.error = res;
      this.state.outputs = Array(this.state.totalShares).fill("")
    }

    this.state.outputs = res;
  }

  render() {
    return [
      this.renderError(),
      el('div', {class:'encryptForm'}, [
        el('div', {className:'columns is-horizontal is-vcentered'}, [
          el('div', {className:'column'}, [
            'Minimum required shares:',
            el('input', {
              className: 'input', 
              type: 'number',
              min: 2, 
              max: this.state.totalShares,
              value: this.state.requiredShares,
              placeholder: 'Minimum required shares',
              oninput: function(e) {
                var val = Number(e.target.value);
                if (val > this.state.totalShares) {
                  val = this.state.totalShares;
                }
                this.update({requiredShares: val});
              }.bind(this)
            })
          ]),
          el('div', {className:'column'}, [
            'Total Shares: ',
            el('input', {
              className: 'input', 
              type: 'number',
              min: 2, 
              value: this.state.totalShares,
              placeholder: 'Shares',
              oninput: function(e) {
                var val = Number(e.target.value);
                this.update({totalShares: val});
              }.bind(this)
            })
          ]),
          el('div', {className:'column'}, [
            el('label', {className: 'checkbox'}, [
              el('input', {
                type: 'checkbox', 
                className: 'checkbox', 
                checked: this.state.aes, 
                onchange: function(e){
                  this.update({aes: e.target.checked})
                }.bind(this)
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
                // We rebuild the textarea on each state change, so this hack
                // is to keep your cursor position correct on the new textarea
                var pos = e.target.selectionEnd;
                this.update({inputString: e.target.value});
                var newText = document.getElementById('inputString');
                newText.focus();
                newText.selectionStart = pos;
                newText.selectionEnd = pos;
              }.bind(this)
            }, this.state.inputString)
          ])
        ]),
        el('div', {className:'columns is-horizontal'}, this.outBoxes()),
      ])
    ];
  }

  renderError() {
    if (!this.state.error) {
      return;
    }

    return el('div', {className: 'notification is-danger'}, this.state.error);
  }

  outBoxes(){
    return this.state.outputs.map(function(item){
      return el('div', {className:'column'}, [
        el('textarea', {className:'textarea', readOnly: true}, item)
      ]);
    });
  }
}
