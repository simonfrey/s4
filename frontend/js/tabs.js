// set initial tab
changeTab('encrypt');

function changeTab(activeTab) {
  // loop through tab handles
  document.querySelectorAll('.tabs li').forEach(function(li){
    if (li.children[0].dataset.tab === activeTab) {
      li.classList.add('is-active');
    } else {
      li.classList.remove('is-active');
    }
  });
  
  // loop through content tabs
  document.querySelectorAll('.tab').forEach(function(tab){
    tab.style.display = tab.id === activeTab ? 'block' : 'none';
  });

  // trigger initial render for decrypt tab
  if (activeTab == 'decrypt') {
    handleShareChange();
  }
}