document.addEventListener('DOMContentLoaded', () => {
  const path = window.location.pathname;

  if (path.includes('projects.html')) {
    loadProjects();
  }

  if (path.includes('contact.html')) {
    setupContactForm();
  }
});

// Load and display projects
function loadProjects() {
  fetch('/api/projects')
    .then(res => res.json())
    .then(projects => {
      const projectList = document.getElementById('project-list');
      projects.forEach(project => {
        const card = document.createElement('div');
        card.className = 'project-card';
        card.innerHTML = `
          <img src="${project.image_url}" alt="${project.name}" />
          <h3>${project.name}</h3>
          <p>${project.description}</p>
          <p><strong>${project.date_range}</strong></p>
          <a href="${project.github_url}" target="_blank">GitHub</a>
        `;
        projectList.appendChild(card);
      });
    })
    .catch(err => {
      console.error('Failed to load projects:', err);
    });
}

// Handle contact form (no backend yet)
function setupContactForm() {
  const form = document.querySelector('form');
  form.addEventListener('submit', e => {
    e.preventDefault();

    const name = form.querySelector('#name').value.trim();
    const email = form.querySelector('#email').value.trim();
    const message = form.querySelector('#message').value.trim();

    if (!name || !email || !message) {
      alert('Please fill in all fields.');
      return;
    }

    // TODO: Send to backend
    alert('Thanks for your message! (Not actually sent yet.)');

    form.reset();
  });
}
