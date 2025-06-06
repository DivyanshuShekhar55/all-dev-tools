let entered_name;

// get the name value on submission of form from html
document.getElementById("form").addEventListener("submit", function (e) {
  e.preventDefault();
  entered_name = document.getElementById("name").value;

  sendToGoBackend();
});

const sendToGoBackend = () => {
  /* To-Do 
        send request to backend with the body containing the entered name 
    */

  console.log(entered_name);
};
