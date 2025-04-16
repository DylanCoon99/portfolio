document.getElementById("login-form").addEventListener("submit", async function (e) {
  e.preventDefault();

  const username = document.getElementById("username").value.trim();
  const password = document.getElementById("password").value.trim();


  console.log(username);
  console.log(password);

  const res = await fetch("http://localhost:8080/api/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password })
  });

  if (res.ok) {
    const data = await res.json();
    localStorage.setItem("adminToken", data.token);
    window.location.href = "admin-dashboard.html"; // or your main admin page
  } else {
    const err = await res.json();
    document.getElementById("error-msg").innerText = err.message || "Login failed";
  }
});
