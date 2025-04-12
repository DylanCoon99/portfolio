document.addEventListener('DOMContentLoaded', () => {
  const container = document.getElementById('project-list'); // your HTML should have this container

  fetch('http://localhost:8080/api/projects')
    .then(response => response.json())
    .then(projects => {
      projects.forEach(project => {
        const card = document.createElement('div');
        card.className = 'project-card';

        card.innerHTML = `
          <img src="${project.image_url}" alt="${project.name}" />
          <h3>${project.name}</h3>
          <p class="date-range">${project.date_range}</p>
          <p class="description">${project.description}</p>
          <a href="${project.github_url}" target="_blank">GitHub</a>
        `;

        container.appendChild(card);
      });
    })
    .catch(error => {
      console.error('Error loading projects:', error);
      container.innerHTML = '<p>Failed to load projects.</p>';
    });
});
