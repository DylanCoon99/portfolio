document.addEventListener("DOMContentLoaded", async () => {
  const token = localStorage.getItem("adminToken");
  if (!token) return window.location.href = "admin-login.html";

  const res = await fetch("http://localhost:8080/api/projects", {
    headers: {
      Authorization: `Bearer ${token}`
    }
  });
  const projects = await res.json();

  const container = document.getElementById("projects-container");
  projects.forEach(project => {
    const card = document.createElement("div");
    card.className = "project-card";
    card.innerHTML = `
      <h3>${project.name}</h3>
      <p>${project.description}</p>
      <img src="${project.image_url}" alt="Project Image" style="width:100px; height:auto;" />
      <form data-id="${project.project_id}" class="image-upload-form">
        <input type="file" name="image" accept="image/*" />
        <button type="submit">Upload Image</button>
      </form>
    `;
    container.appendChild(card);
  });

  document.querySelectorAll(".image-upload-form").forEach(form => {
    form.addEventListener("submit", async e => {
      e.preventDefault();

      const projectId = form.getAttribute("data-id");
      const fileInput = form.querySelector("input[name='image']");
      const file = fileInput.files[0];

      if (!file) return alert("Please select an image");

      const formData = new FormData();
      formData.append("file", file);


      const uploadRes = await fetch(`http://localhost:8080/api/admin/projects/${projectId}/upload-image`, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`
        },
        body: formData
      });

      const result = await uploadRes.json();
      if (uploadRes.ok) {
        alert("Image uploaded successfully!");
        window.location.reload();
      } else {
        alert(result.message || "Upload failed");
      }
    });
  });
});
