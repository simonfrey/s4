// Root of the whole application!
class App extends Component {
  constructor(children) {
    super(children)
    // attach to the DOM
    this.root = document.getElementById("app");
    // trigger initial render
    this.update();
  }

  render() {
    return this.children;
  }
}    

// initialize
let instance = new App([
  new Intro(),
  new Warnings(),
  new Tabs({
    'encrypt': new Encrypt(),
    'decrypt': new Decrypt(),
    'info': new Info()
  })
]);