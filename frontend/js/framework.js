class Component {
  constructor(children) {
    this.children = children;
    this.root = el('div', {className:'componentRoot'});
  }
  state = {}
  root = null

  // update merges in state change (not deep!), runs cals, and re-renders
  update(obj) {
    Object.assign(this.state, obj);
    this.calc();
    this.replaceContent();
  }

  // calc runs after update to calculate any dependent state
  calc() {}
  
  replaceContent() {
    this.root.innerHTML = '';
    append(this.root, this.render());
  }
}

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
  } 
  
  if (children.root && children.replaceContent && typeof children.replaceContent === 'function') {
    // it's a component, trigger content replacement within the child component
    children.replaceContent();
    append(el, children.root);
  } else if (typeof children === 'string') {
    // it's just text
    el.appendChild(document.createTextNode(children));
  } else if (Array.isArray(children)) {
    // it's an array!
    children.forEach(function(child){
      append(el, child);
    });
  } else {
    el.appendChild(children);
  }
}