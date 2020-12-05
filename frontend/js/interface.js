// application state
var state = {
  activeTab: 'encrypt',
  error: ''
};

// initialize
render();

// build a DOM element!
function el(tag, attrs, children){
  var el = document.createElement(tag);

  if (attrs) {
    Object.assign(el, attrs);
  }

  append(el, children);

  return el;
}

// recursively appendChild for (potentially) arrays of arrays
function append(el, children) {
  if (!children) {
    return;
  } else if (typeof children === 'string') {
    el.appendChild(document.createTextNode(children));
  } else if (Array.isArray(children)) {
    children.forEach(function(child){
      append(el, child);
    });
  } else if (children) {
    el.appendChild(children);
  }
}


// update and renders
function render() {
  var root = document.getElementById("app");
  root.innerHTML = '';
  append(root, [
    renderIntro(),
    renderWarnings(),
    renderTabs(),
    renderErrors(),
    renderContent()
  ]);
}

function update (obj){
  mergeDeep(state, obj);
  render();
}

function isObject(item) {
  return (item && typeof item === 'object' && !Array.isArray(item));
}

// recursive state updater, ignores already-set keys
function mergeDeep(target, source) {
  if (isObject(target) && isObject(source)) {
    for (var key in source) {
      if (isObject(source[key])) {
        if (!target[key]) {
          Object.assign(target, { [key]: {} });
        }
        mergeDeep(target[key], source[key]);
      } else {
        Object.assign(target, { [key]: source[key] });
      }
    }
  }
}

function renderIntro() {
  return [
    el('h1', {className: 'title'}, "Simple Shamir's Secret Sharing (s4)"),
    el('p', {className: 'subtitle'}, "Share your secret with a cryptographically secure method. All running locally in your browser. Your secrets never leave your machine!")
  ];
}

// TABS
function renderTabs() {
  return el('div', {className:'tabs is-boxed'}, 
    el('ul', {}, [
      renderTab('encrypt'),
      renderTab('decrypt'),
      renderTab('info')
    ]
  ));
}

function renderTab(key) {
  return el('li', {className: state.activeTab == key ? 'is-active' : ''}, 
    el('a', {
      onclick: function(){update({activeTab: key});}
    }, capitalize(key))
  );
}

function capitalize(str) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

// notification
function renderErrors() {
  if (!state.error) {
    return;
  }

  return el('div', {className: 'notification is-danger'}, [
    el('button', {className:'devare', onclick:function(){update({error: ''})}}),
    state.error
  ]);
}

function renderContent(){
  switch(state.activeTab) {
    case 'encrypt':
      return renderEncrypt();
    case 'decrypt':
      return renderDecrypt();
    case 'info':
      return renderInfo();
  }
}

function renderInfo() {
  return el('div', {className: 'columns is-horizontal', innerHTML: `
  <div class="column">
  <h3 class="subtitle">Videos</h3>
  <p>The following youtube videos explain you more about <a
          href="https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing" target="_blank">Shamir's
          Secret Sharing</a></p>
  <ul>
      <li><a href="https://www.youtube-nocookie.com/embed/TQ-DsEZBuQY" target="_blank">What is Key
              Sharding? Shamir’s Secret Sharing Explained</a> (<b>Easy</b>)</li>
      <li><a href="https://www.youtube-nocookie.com/embed/rWPZoz0aux4" target="_blank">Shamir's Secret
              Sharing - Solution and alternative to Lagrange</a> (<b>More in-depth</b>)</li>
  </ul>
  </div>
  <div class="column ">
  <h3 class="subtitle">About s4</h3>
  <p>With <b>S</b>imple <b>S</b>hamir's <b>S</b>ecret <b>S</b>haring (s4) I want to provide you an
      easy to use interface for this beautiful little piece of math.</p>
  <p>s4 is open source and provided as it is. You can find the code on <a
          href="https://github.com/simonfrey/s4/issues" target="_blank">Github</a> and this website
      works compvarely offline. Save it to your computer (Ctrl+S) in order to not lose access to the
      s4 in case this website will be down at any point in the future.</p>
  <p> Please note that s4 is provided as it is and I do not take responsibility for any bugs. s4 is a
      tiny layer around <a href="https://github.com/hashicorp/vault"
          target="_blank">hashicorp vault shamir</a> and golangs <a
          href="https://github.com/gtank/cryptopasta/blob/master/encrypt.go" target="_blank">AES
          encryption</a>.</p>
  <p>If you find any issues, please report them via <a href="https://github.com/simonfrey/s4/issues"
          target="_blank">Github issues</a> and if you want to tip me for the work on this project,<a
          href="https://simon-frey.com/tip" target="_blank"> feel free to do so !</a></p>
  </div>`});
}