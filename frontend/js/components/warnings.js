class Warnings extends Component {
  constructor() {
    super();
    window.addEventListener('online', this.update);
    window.addEventListener('offline',  this.update);
    this.update();
  }

  state = {
    offline: false,
    file: false
  }

  calc() {
    this.state.offline = !navigator.onLine;
    this.state.file = window.location.protocol == 'file:';
  }

  render() {
    return [
      this.state.offline ? '' : el('p', {className: 'has-text-danger'}, 'Your browser is online. Please disconnect from your network to use s4 in a safer way'),
      this.state.file ? '' : el('p', {className: 'has-text-danger'}, 'You are running s4 from a webserver, this is insecure. Please savethe webpage (Ctrl+S) locally and open this file.')
    ];
  }
}

