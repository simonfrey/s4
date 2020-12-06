const base64regex = /^([0-9a-zA-Z+/]{4})*(([0-9a-zA-Z+/]{2}==)|([0-9a-zA-Z+/]{3}=))?$/;

class Decrypt extends Component {
  state = {
    error: '',
    shares: 2,
    inputs: ['', ''],
    output: ''
  }

  calc() {
    let res = Recover_fours(this.state.inputs);
  
    if (!base64regex.test(res)){
      this.state.output = '';
      this.state.error = res;
      return;
    }
  
    this.state.output = atob(res);
    this.state.error = '';
  }

  render() {
    return [
      this.renderError(),
      el('div', {class:'decryptForm columns'}, [
        el('div', {className:'column is-vertical'}, [
          el('div', {className:'columns is-horizontal'}, [
            el('div', {className:'column'}, [
              'Shares: ',
              el('input', {
                className: 'input', 
                type: 'number',
                min: 2,
                value: this.state.shares,
                placeholder: 'Shares',
                oninput: function(e) {
                  var val = Number(e.target.value);
                  this.update({
                    shares: val, 
                    inputs: Array(val).fill('')
                  });
                }.bind(this)
              })
            ])
          ]),
          el('div', {className:'columns is-horizontal'}, this.inputs()),
          el('div', {className:'columns'}, [
            el('div', {className:'column'}, [
              el('textarea', {
                id: 'output',
                className:'textarea',
                placeholder: 'Your decrypted output'
              }, this.state.output)
            ])
          ])
        ])
      ])
    ];
  }

  renderError() {
    if (!this.state.error) {
      return;
    }

    return el('div', {className: 'notification is-danger'}, this.state.error);
  }

  inputs(){
    return this.state.inputs.map(function(item, i){
      return el('div', {className:'column'}, [
        el('textarea', {
          className:'textarea', 
          placeholder: 'Your encrypted input share',
          oninput: function(e) {
            this.state.inputs[i] = e.target.value;
            this.update({inputs: this.state.inputs});
          }.bind(this)
        }, item)
      ]);
    }.bind(this));
  }
}
