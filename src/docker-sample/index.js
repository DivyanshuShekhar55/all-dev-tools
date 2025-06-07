let entered_name;

// get the name value on submission of form from html
document.getElementById("form").addEventListener("submit", function (e) {
  e.preventDefault();
  entered_name = document.getElementById("name").value;

  sendToGoBackend();
});

const sendToGoBackend = async () => {
  const res = await fetch("http://localhost:8000/submit", {
    method: "POST",
    body: JSON.stringify({ name: entered_name }),
  });

  console.log(res.status);

  console.log(entered_name);
};
