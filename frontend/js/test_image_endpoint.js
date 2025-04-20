document.addEventListener('DOMContentLoaded', () => {
  const container = document.getElementById('project-1-image'); // your HTML should have this container

  loadImage(1);

});


async function loadImage(projectId) {

  const response = await fetch(`http://localhost:8080/api/projects/${projectId}/image`);

  if (!response.ok) {
    console.error("Failed to load image");
    return;
  }

  const blob = await response.blob();
  const imgUrl = URL.createObjectURL(blob);

  const img = document.createElement("img");
  img.src = imgUrl;
  img.alt = "Project Image";
  img.width = 300;

  document.body.appendChild(img);

}