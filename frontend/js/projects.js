document.addEventListener('DOMContentLoaded', async () => {
  const container = document.getElementById('project-list');

  try {
    const response = await fetch('https://portfolio-04d9.onrender.com/api/projects');
    const projects = await response.json();

    for (const project of projects) {
      const card = document.createElement('div');
      card.className = 'project-card';

      const imgSrc = await loadImage(project.project_id);

      card.innerHTML = `
        <img src="${imgSrc}" alt="${project.name}" />
        <h3>${project.name}</h3>
        <p class="date-range">${project.date_range}</p>
        <p class="description">${project.description}</p>
        <a href="${project.github_url}" target="_blank">GitHub</a>
      `;

      container.appendChild(card);
    }

  } catch (error) {
    console.error('Error loading projects:', error);
    container.innerHTML = '<p>Failed to load projects.</p>';
  }
});



async function loadImage(projectId) {

  const response = await fetch(`https://portfolio-04d9.onrender.com/api/projects/${projectId}/image`);

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

  //document.body.appendChild(img);

  return img.src

}
