var sidenav = document.getElementById("mySidenav");
var openBtn = document.getElementById("openBtn");
var closeBtn = document.getElementById("closeBtn");
openBtn.onclick = function(event) {
    openNav(event);
};
closeBtn.onclick = function(event) {
    closeNav(event);
};

function openNav() {
event.preventDefault(); 
  sidenav.classList.add("active");
}

function closeNav(event) {
    event.preventDefault(); 
    sidenav.classList.remove("active");
}

