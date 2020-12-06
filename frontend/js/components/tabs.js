class Tabs extends Component {
  state = {
    activeTab: 'encrypt'
  }

  render() {
    return [
      el('div', {className:'tabs is-boxed'}, 
        el('ul', {}, [
          this.tab('encrypt'),
          this.tab('decrypt'),
          this.tab('info'),
        ]),
      ),
      this.tabContent()
    ];
  }

  tab(key) {
    return el('li', {className: this.state.activeTab == key ? 'is-active' : ''}, 
      el('a', {
        onclick: function(){
          this.update({activeTab: key})
        }.bind(this)
      }, this.capitalize(key))
    );
  }

  capitalize(str) {
    return str.charAt(0).toUpperCase() + str.slice(1);
  }

  tabContent(){
    return this.children[this.state.activeTab];
  }
}