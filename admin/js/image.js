document.getElementById('upload-form').addEventListener('submit', async function (e) {
  e.preventDefault();

  const fileInput = document.getElementById('file');
  if (fileInput.files.length === 0) {
    alert("Please choose a file to upload.");
    return;
  }

  const formData = new FormData();
  formData.append("file", fileInput.files[0]);

  try {
    const res = await fetch("http://localhost:8080/api/admin/upload", {
      method: "POST",
      headers: { "Access-Control-Allow-Origin": "Origin" },
      body: formData
    });

    const data = await res.json();

    if (res.ok) {
      document.getElementById("upload-status").textContent = `Uploaded! Path: ${data.path}`;
      // Optionally store `data.path` for your DB insertion
    } else {
      document.getElementById("upload-status").textContent = `Upload failed: ${data.error}`;
    }
  } catch (err) {
    console.error(err);
    document.getElementById("upload-status").textContent = "An error occurred while uploading.";
  }
});
